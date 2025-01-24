package cmd

import (
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
    IsEverything bool
    IsSDE       bool
    IsSRE       bool
    IsCloud     bool
    IsTPM       bool
}

var (
	jobRoles = []string{"everything", "sde", "sre", "cloud", "tpm"}

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

            for _, role := range fRoles {
                var vars TemplateVars

                switch role {
                case "everything":
                    vars.IsEverything = true
                case "sde":
                    vars.IsSDE = true
                case "sre":
                    vars.IsSRE = true
                case "cloud":
                    vars.IsCloud = true
                case "tpm":
                    vars.IsTPM = true
                }

                err = templates.ExecuteTemplate(os.Stdout, "frame.gohtml", vars)
                if err != nil {
                    logger.Fatal(fmt.Errorf("error executing templates: %w", err))
                }
            }
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArrayVarP(&fRoles, "roles", "r", jobRoles, "Versions (roles) of the résumé to generate.")
}
