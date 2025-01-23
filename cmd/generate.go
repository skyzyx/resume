package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/charmbracelet/log"
	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
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

	// funcMap = template.FuncMap{
	// 	"join": strings.Join,
	// }

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

			templates, err := template.ParseFS(embeds, pattern)
			if err != nil {
				logger.Fatal(fmt.Errorf("error parsing templates: %w", err))
			}

			err = templates.Execute(os.Stdout, nil)
			if err != nil {
				logger.Fatal(fmt.Errorf("error executing templates: %w", err))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArrayVarP(&fRoles, "roles", "r", jobRoles, "Versions (roles) of the résumé to generate.")
}

// -----------------------------------------------------------------------------

// 	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
// 		log.Fatal(err)
// 	}

// -----------------------------------------------------------------------------

// // T0.tmpl is a plain template file that just invokes T1.
// {"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
// // T1.tmpl defines a template, T1 that invokes T2.
// {"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
// // T2.tmpl defines a template T2.
// {"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},

// pattern is the glob pattern used to find all the template files.
// pattern := filepath.Join(dir, "*.tmpl")

// // Here starts the example proper.
// // T0.tmpl is the first name matched, so it becomes the starting template,
// // the value returned by ParseGlob.
// tmpl := template.Must(template.ParseGlob(pattern))

// err := tmpl.Execute(os.Stdout, nil)
// if err != nil {
// 	log.Fatalf("template execution: %s", err)
// }
