package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/charmbracelet/log"
	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/debug"
	"github.com/spf13/cobra"
)

type TemplateVars struct {
	Resume    string
	ResumeRaw string
	ResumeDir string

	IsAll   bool
	IsATS   bool
	IsCloud bool
	IsSDE   bool
	IsTPM   bool
}

var (
	jobRoles = []string{
		"all",
		"cloud",
		"sde",
		"tpm",
	}

	fRoles []string

	//go:embed templates/*.gohtml
	embeds embed.FS

	// Instantiate a logger.
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
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
				logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", "../resumes", err))
			}

			// Make sure the directory exists.
			err = os.MkdirAll(resumeDir, 0o0777)
			if err != nil {
				logger.Fatal(fmt.Errorf("error creating `resumes` directory: %w", err))
			}

			// For each role, write the Markdown content to a file.
			for _, role := range fRoles {
				for _, isATS := range []bool{false, true} {
					generateMarkdown(templates, role, resumeDir, isATS)

					filename := fileNames[role] + ".docx"
					if isATS {
						filename = fileNames[role] + ".ats.docx"
					}

					// Convert the Markdown to DOCX.
					docx := exec.Command(
						"pandoc",
						"-r", "gfm",
						"-w", "docx",
						"--reference-doc="+filepath.Join("render", "reference.docx"),
						"--output="+filepath.Join("resumes", filename),
						filepath.Join("resumes", fileNames[role]+".md"),
					)
					err := docx.Run()
					if err != nil {
						logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+".docx", err))
					}

					filename = fileNames[role] + ".odt"
					if isATS {
						filename = fileNames[role] + ".ats.odt"
					}

					// Convert the Markdown to ODT.
					odt := exec.Command(
						"pandoc",
						"-r", "gfm",
						"-w", "odt",
						"--reference-doc="+filepath.Join("render", "reference.odt"),
						"--output="+filepath.Join("resumes", filename),
						filepath.Join("resumes", fileNames[role]+".md"),
					)
					err = odt.Run()
					if err != nil {
						logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+".odt", err))
					}
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArrayVarP(&fRoles, "roles", "r", jobRoles, "Versions (roles) of the résumé to generate.")
}

func generateMarkdown(templates *template.Template, role, resumeDir string, isATS bool) {
	vars := TemplateVars{
		Resume:    "https://github.com/skyzyx/resume/blob/master/resumes/" + fileNames[role],
		ResumeDir: "https://github.com/skyzyx/resume/blob/master/resumes/",
		ResumeRaw: "https://github.com/skyzyx/resume/raw/master/resumes/" + fileNames[role],
	}

	if isATS {
		vars.IsATS = true
	}

	// Pass an identifier to the template, so that we can switch on it.
	switch role {
	case "all":
		vars.IsAll = true
	case "cloud":
		vars.IsCloud = true
	case "sde":
		vars.IsSDE = true
	case "tpm":
		vars.IsTPM = true
	}

	absPathMd := filepath.Join(resumeDir, fileNames[role]+".md")
	if isATS {
		absPathMd = filepath.Join(resumeDir, fileNames[role]+".ats.md")
	}

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
	bufOut := strings.NewReader(
		reTooManyLinebreaks.ReplaceAllString(bufIn.String(), "\n\n"),
	)

	// Write the buffer to the file.
	_, err = bufOut.WriteTo(w)
	if err != nil {
		logger.Fatal(fmt.Errorf("error writing buffer to file: %w", err))
	}

	// Log the success.
	logger.Infof("wrote %s", absPathMd)
}
