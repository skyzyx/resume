package cmd

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/charmbracelet/log"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/debug"
)

type TemplateVars struct {
	Resume    string
	ResumeRaw string
	ResumeDir string

	IsAll   bool
	IsCloud bool
	IsSDE   bool

	IsGoPDF   bool
	IsWorkday bool

	Counter int
}

var (
	// Roles
	jobRoles = []string{
		"all",
		"cloud",
		// "em",
		"sde",
		// "tpm",
	}

	// Flags
	fRoles []string

	//go:embed templates/*.gohtml
	embeds embed.FS

	// Global vars
	pdfDescription string
	pdfKeywords    string

	// Instantiate a logger.
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Level:           log.DebugLevel,
	})

	funcMap = sprig.FuncMap()

	// generateCmd represents the generate command
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Performs the generation step, producing multiple formats.",
		Long: clihelpers.LongHelpText(`
		Performs the generation step, producing multiple formats.

		Reads Markdown content from templates, then passes the content to Pandoc
		for transformation to other formats.`),
		Run: func(cmd *cobra.Command, args []string) {
			reJobRoles := regexp.MustCompile(`^(` + strings.Join(jobRoles, "|") + `)$`)

			// Validate the roles.
			for _, role := range fRoles {
				if !reJobRoles.MatchString(role) {
					logger.Fatalf("invalid role: %s", role)
				}
			}

			pp := debug.GetSpew()
			funcMap["debug"] = pp.Sdump
			funcMap["workdayJobDescription"] = func(jobTitle, company, location, from, to string) string {
				return strings.TrimSpace(dedent.Dedent(`
				Job Title: ` + jobTitle + ` \
				Company: ` + company + ` \
				Location: ` + location + ` \
				From: ` + from + ` \
				To: ` + to + `
				`))
			}

			pattern := filepath.Join("templates", "*.gohtml")

			// Load the template content from the embedded filesystem.
			templates, err := template.New("").Funcs(funcMap).ParseFS(embeds, pattern)
			if err != nil {
				logger.Fatal(fmt.Errorf("error parsing templates: %w", err))
			}

			// Get the absolute path to the `resumes` directory.
			resumeDir, err := filepath.Abs(
				filepath.Join(".", "resumes"),
			)
			if err != nil {
				logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", "./resumes", err))
			}

			// Make sure the directory exists.
			err = os.MkdirAll(resumeDir, 0o0777)
			if err != nil {
				logger.Fatal(fmt.Errorf("error creating `resumes` directory: %w", err))
			}

			// Start a local web server.
			server := &http.Server{
				Addr: ":11235",
			}

			http.Handle("/", http.FileServer(http.Dir("./render")))
			logger.Debug("Starting HTTP server on :11235")

			go func() {
				if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("HTTP server error: %v", err)
				}
			}()

			// For each role, write the Markdown content to a file.
			for _, role := range fRoles {
				audience := ""

				generateMarkdown(templates, role, resumeDir, audience)
				generateWord(role, resumeDir, audience, "docx", false)
				generateWord(role, resumeDir, audience, "docx", true)
				// generateWord(role, resumeDir, audience, "odt")
				generatePDF(role, resumeDir, audience)
			}

			// Terminate the web server.
			server.Close()
			logger.Debug("Stopped HTTP server")
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArrayVarP(&fRoles, "roles", "r", jobRoles, "Versions (roles) of the résumé to generate.")
}

func generateMarkdown(templates *template.Template, role, resumeDir, audience string) {
	vars := TemplateVars{
		Resume:    "https://github.com/skyzyx/resume/blob/master/resumes/" + fileNames[role],
		ResumeDir: "https://github.com/skyzyx/resume/blob/master/resumes/",
		ResumeRaw: "https://github.com/skyzyx/resume/raw/master/resumes/" + fileNames[role],
	}

	// Pass an identifier to the template, so that we can switch on it.
	switch role {
	case "all":
		vars.IsAll = true
	case "cloud":
		vars.IsCloud = true
	case "sde":
		vars.IsSDE = true
	}

	// -------------------------------------------------------------------------
	// Generate the full Markdown resume.

	absPathMd := filepath.Join(resumeDir, fileNames[role]+audience+".md")

	// Create the Markdown file.
	w, err := os.Create(absPathMd)
	if err != nil {
		logger.Fatal(fmt.Errorf("error creating `%s`: %w", filepath.Base(absPathMd), err))
	}

	defer w.Close()

	// We will be collecting the output in a buffer, so that we can
	// clean it before writing it to the file.
	bufIn := new(bytes.Buffer)

	// Execute the templates.
	err = templates.ExecuteTemplate(bufIn, "frame.gohtml", vars)
	if err != nil {
		logger.Fatal(fmt.Errorf("error executing templates: %w", err))
	}

	// Cleanup the contents of the buffer.
	reTooManyLinebreaks := regexp.MustCompile(`\n{3,}`)
	strOut := reTooManyLinebreaks.ReplaceAllString(bufIn.String(), "\n\n")

	// Reduce line breaks between bullets.
	reReduceLinebreaks := regexp.MustCompile(`\*\s([^\n]+)\n{2}\*`)
	strOut = reReduceLinebreaks.ReplaceAllString(strOut, "* ${1}\n*")
	strOut = reReduceLinebreaks.ReplaceAllString(strOut, "* ${1}\n*")

	bufOut := strings.NewReader(strOut)

	// Write the buffer to the file.
	_, err = bufOut.WriteTo(w)
	if err != nil {
		logger.Fatal(fmt.Errorf("error writing buffer to file: %w", err))
	}

	// Log the success.
	logger.Infof("wrote %s", absPathMd)

	// -------------------------------------------------------------------------
	// Generate the Workday version

	absPathMdWd := filepath.Join(resumeDir, fileNames[role]+audience+"-workday.md")

	// Create the Markdown file.
	wd, err := os.Create(absPathMdWd)
	if err != nil {
		logger.Fatal(fmt.Errorf("error creating `%s`: %w", filepath.Base(absPathMdWd), err))
	}

	defer wd.Close()

	// We will be collecting the output in a buffer, so that we can
	// clean it before writing it to the file.
	workdayIn := new(bytes.Buffer)

	vars.IsWorkday = true

	// Execute the templates.
	err = templates.ExecuteTemplate(workdayIn, "frame-workday.gohtml", vars)
	if err != nil {
		logger.Fatal(fmt.Errorf("error executing templates: %w", err))
	}

	// Cleanup the contents of the buffer.
	reTooManyLinebreaks = regexp.MustCompile(`\n{3,}`)
	strOut = reTooManyLinebreaks.ReplaceAllString(workdayIn.String(), "\n\n")

	// Adjustments
	strOut = strings.ReplaceAll(strOut, "á", "a")
	strOut = strings.ReplaceAll(strOut, "_", "")
	strOut = strings.ReplaceAll(strOut, "`", "")
	strOut = strings.ReplaceAll(strOut, "’", "'")
	strOut = strings.ReplaceAll(strOut, "“", "")
	strOut = strings.ReplaceAll(strOut, "”", "")
	strOut = strings.ReplaceAll(strOut, `"`, "")

	// Reduce line breaks between bullets.
	reReduceLinebreaks = regexp.MustCompile(`\*\s([^\n]+)\n{2}\*`)
	strOut = reReduceLinebreaks.ReplaceAllString(strOut, "* ${1}\n*")
	strOut = reReduceLinebreaks.ReplaceAllString(strOut, "* ${1}\n*")

	// Strip links
	strOut = regexp.MustCompile(`\[([^\]]+)\]\((?:[^\)]+)\)`).ReplaceAllString(strOut, "${1}")
	strOut = regexp.MustCompile(`\[([^\]]+)\]\[(?:[A-Za-z0-9\.\s]+)\]`).ReplaceAllString(strOut, "${1}")
	strOut = regexp.MustCompile(`\[([^\]]+)\]`).ReplaceAllString(strOut, "${1}")

	workdayOut := strings.NewReader(strOut)

	// Write the buffer to the file.
	_, err = workdayOut.WriteTo(wd)
	if err != nil {
		logger.Fatal(fmt.Errorf("error writing buffer to file: %w", err))
	}

	// Log the success.
	logger.Infof("wrote %s", absPathMdWd)

	vars.IsWorkday = false

	// -------------------------------------------------------------------------
	// Blurb ONLY (PDF metadata)

	// We will be collecting the output in a buffer, so that we can
	// clean it before writing it to the file.
	blurbIn := new(bytes.Buffer)

	// Execute the templates.
	err = templates.ExecuteTemplate(blurbIn, "blurb-only.gohtml", vars)
	if err != nil {
		logger.Fatal(fmt.Errorf("error executing templates: %w", err))
	}

	// Cleanup the contents of the buffer.
	blurbOut := strings.ReplaceAll(blurbIn.String(), "*", "")
	blurbOut = strings.ReplaceAll(blurbOut, "\n", "")

	// Save to global variables for PDF metadata.
	pdfDescription = blurbOut

	// -------------------------------------------------------------------------
	// Skills ONLY (PDF metadata)

	// We will be collecting the output in a buffer, so that we can
	// clean it before writing it to the file.
	skillsIn := new(bytes.Buffer)
	skillsOut := new(bytes.Buffer)

	vars.IsGoPDF = true

	// Execute the templates.
	err = templates.ExecuteTemplate(skillsIn, "skills-only.gohtml", vars)
	if err != nil {
		logger.Fatal(fmt.Errorf("error executing templates: %w", err))
	}

	// Cleanup the contents of the buffer.
	reMarkdownLinks := regexp.MustCompile(`\[([^\]]+)\](\[([^\]]+)\])?`)
	skillsTmp := strings.NewReader(
		reMarkdownLinks.ReplaceAllString(skillsIn.String(), "${1}"),
	)

	_, err = skillsTmp.WriteTo(skillsOut)
	if err != nil {
		logger.Fatal(fmt.Errorf("error writing buffer to file: %w", err))
	}

	// Save to global variables for PDF metadata.
	pdfKeywords = strings.TrimSpace(skillsOut.String())
	pdfKeywords = strings.ReplaceAll(pdfKeywords, "Amazon Web Services", "Amazon+Web+Services")
	pdfKeywords = strings.ReplaceAll(pdfKeywords, "Amazon Linux", "Amazon+Linux")
	pdfKeywords = strings.ReplaceAll(pdfKeywords, "Amazon ", "")
	pdfKeywords = strings.ReplaceAll(pdfKeywords, "Amazon+Web+Services", "Amazon Web Services")
	pdfKeywords = strings.ReplaceAll(pdfKeywords, "Amazon+Linux", "Amazon Linux")

	pdfKeywords = strings.ReplaceAll(pdfKeywords, "AWS ", "")
	pdfKeywords = strings.ReplaceAll(pdfKeywords, " (modern)", "")
}

func generateWord(role, resumeDir, audience, format string, isWorkday bool) {
	workdayFilename := ""
	if isWorkday {
		workdayFilename = "-workday"
	}

	docx := exec.Command(
		"pandoc",
		"-r", "gfm",
		"-w", format,
		"--reference-doc="+filepath.Join("render", "reference."+format),
		"--output="+filepath.Join(resumeDir, fileNames[role]+audience+workdayFilename+"."+format),
		filepath.Join(resumeDir, fileNames[role]+audience+workdayFilename+".md"),
	)
	err := docx.Run()
	if err != nil {
		logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+audience+workdayFilename+"."+format, err))
	}

	fpAbs, err := filepath.Abs(
		filepath.Join(resumeDir, fileNames[role]+audience+workdayFilename+"."+format),
	)
	if err != nil {
		logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", "./resumes", err))
	}

	// Log the success.
	logger.Infof("wrote %s", fpAbs)
}

func generatePDF(role, resumeDir, audience string) {
	// Convert the Markdown to HTML.
	html := exec.Command(
		"pandoc",
		"-r", "gfm",
		"-w", "html5+ascii_identifiers+auto_identifiers+gfm_auto_identifiers+smart+task_lists",
		"--columns=20000",
		"--eol=lf",
		"--template=cmd/templates/pdf.tmpl.html",
		"--output="+filepath.Join("render", fileNames[role]+audience+".html"),
		filepath.Join(resumeDir, fileNames[role]+audience+".md"),
	)
	err := html.Run()
	if err != nil {
		logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+audience+".html", err))
	}

	// Convert the HTML to PDF.
	pdf := exec.Command(
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"--headless",
		"--virtual-time-budget=5000",
		"--no-pdf-header-footer",
		`--print-to-pdf=`+filepath.Join(resumeDir, fileNames[role]+audience+".pdf"),
		"http://0.0.0.0:11235/"+fileNames[role]+audience+".html",
	)
	err = pdf.Run()
	if err != nil {
		logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+audience+".pdf", err))
	}

	fpAbs, err := filepath.Abs(
		filepath.Join(resumeDir, fileNames[role]+audience+".pdf"),
	)
	if err != nil {
		logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", "./resumes", err))
	}

	// Log the success.
	logger.Infof("wrote %s", fpAbs)

	// -------------------------------------------------------------------------

	// Strip PDF metadata
	exiftoolStrip := exec.Command(
		"exiftool",
		"-all:all=",
		filepath.Join(resumeDir, fileNames[role]+audience+".pdf"),
		"-overwrite_original",
	)
	err = exiftoolStrip.Run()
	if err != nil {
		logger.Fatal(fmt.Errorf("error stripping metadata from `%s`: %w", fileNames[role]+audience+".pdf", err))
	}

	logger.Debugf("stripped metadata from %s", filepath.Join(resumeDir, fileNames[role]+audience+".pdf"))

	time.Sleep(100 * time.Millisecond)

	var pdfRoles string
	switch role {
	case "all":
		pdfRoles = "Software Engineer, DevEx, Cloud, SRE, Platform, Infrastructure"
	case "cloud":
		pdfRoles = "Cloud, SRE, Site Reliability, DevOps, DevSecOps, DevTools, Infrastructure, Platform Engineering"
	case "sde":
		pdfRoles = "Software Engineer, DevTools, Internal Tools, Developer Productivity, Developer Experience, DevEx"
	}

	// Write correct PDF metadata
	exiftoolUpdate := exec.Command(
		"exiftool",
		`-title="Ryan Parman (`+pdfRoles+`)"`,
		`-author="Ryan Parman"`,
		`-subject="`+pdfDescription+`"`,
		`-sep ", "`,
		`-keywords="`+pdfKeywords+`"`,
		filepath.Join(resumeDir, fileNames[role]+audience+".pdf"),
	)
	err = exiftoolUpdate.Run()
	if err != nil {
		logger.Fatal(fmt.Errorf("error updating metadata from `%s`: %w", fileNames[role]+audience+".pdf", err))
	}

	logger.Debugf("updated metadata for %s", filepath.Join(resumeDir, fileNames[role]+audience+".pdf"))

	err = os.Remove(filepath.Join(resumeDir, fileNames[role]+audience+".pdf_original"))
	if err != nil {
		logger.Fatal(fmt.Errorf("error removing `%s`: %w", fileNames[role]+audience+".pdf_original", err))
	}
}
