package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/spf13/cobra"
)

type TemplateVars struct {
	IsAll   bool
	IsSDE   bool
	IsSRE   bool
	IsCloud bool
	IsTPM   bool
}

var (
	jobRoles = []string{"all", "sde", "sre", "cloud", "tpm"}

	fRoles []string

	//go:embed templates/*.gohtml
	embeds embed.FS

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})

	funcMap = template.FuncMap{
		"join": strings.Join,
	}

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

			pattern := filepath.Join("templates", "*.gohtml")

			templates, err := template.New("").Funcs(funcMap).ParseFS(embeds, pattern)
			if err != nil {
				logger.Fatal(fmt.Errorf("error parsing templates: %w", err))
			}

			resumeDir, err := filepath.Abs(
				filepath.Join(".", "resumes"),
			)
			if err != nil {
				logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", "../resumes", err))
			}

			err = os.MkdirAll(resumeDir, 0o0777)
			if err != nil {
				logger.Fatal(fmt.Errorf("error creating `resumes` directory: %w", err))
			}

			if _, err := os.Stat(resumeDir); !os.IsNotExist(err) {
				logger.Infof("created %s", resumeDir)
			} else {
				logger.Fatal(fmt.Errorf("error creating `resumes` directory: %w", err))
			}

			for _, role := range fRoles {
				var vars TemplateVars

				switch role {
				case "all":
					vars.IsAll = true
				case "sde":
					vars.IsSDE = true
				case "sre":
					vars.IsSRE = true
				case "cloud":
					vars.IsCloud = true
				case "tpm":
					vars.IsTPM = true
				}

				absPathMd, err := filepath.Abs(
					filepath.Join(resumeDir, fileNames[role]+".md"),
				)
				if err != nil {
					logger.Fatal(fmt.Errorf("error determining absolute path for `%s`: %w", fileNames[role]+".md", err))
				}

				w, err := os.Create(absPathMd)
				if err != nil {
					logger.Fatal(fmt.Errorf("error creating `%s`: %w", fileNames[role]+".md", err))
				}

				defer w.Close()

				bufIn := new(bytes.Buffer)

				err = templates.ExecuteTemplate(bufIn, "frame.gohtml", vars)
				if err != nil {
					logger.Fatal(fmt.Errorf("error executing templates: %w", err))
				}

				// Cleanup
				reTooManyLinebreaks := regexp.MustCompile(`\n{3,}`)
				bufOut := strings.NewReader(
					reTooManyLinebreaks.ReplaceAllString(bufIn.String(), "\n\n"),
				)

				_, err = bufOut.WriteTo(w)
				if err != nil {
					logger.Fatal(fmt.Errorf("error writing buffer to file: %w", err))
				}

				logger.Infof("wrote %s", absPathMd)

			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArrayVarP(&fRoles, "roles", "r", jobRoles, "Versions (roles) of the résumé to generate.")
}
