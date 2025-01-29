package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/charmbracelet/log"
	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/spf13/cobra"
)

var (
	fLoadFile string
    fWritePDF string

	// serveCmd represents the serve command
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Just run the web server with ./render as the root.",
		Long: clihelpers.LongHelpText(`
        Just run the web server with ./render as the root.

        This is useful for tweaking CSS styles and seeing the changes in real-time.
        If you pass the --load-file and --write-pdf flags, it will also convert the
        HTML to a PDF, then terminate the web server.`),
		Run: func(cmd *cobra.Command, args []string) {
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

            if fLoadFile != "" && fWritePDF != "" {
    			logger.Info("Converting HTML to PDF...")

                // Convert the HTML to PDF.
                pdf := exec.Command(
                    "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
                    "--headless",
                    "--virtual-time-budget=5000",
                    "--no-pdf-header-footer",
                    "--print-to-pdf="+fWritePDF,
                    "http://0.0.0.0:11235/"+fLoadFile,
                )
                err := pdf.Run()
                if err != nil {
                    logger.Fatal(fmt.Errorf("error creating PDF: %w", err))
                }

                server.Close()
    			logger.Debug("Stopped HTTP server")
            }
        },
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&fLoadFile, "load-file", "l", "", "The file to load from the render directory.")
	serveCmd.Flags().StringVarP(&fWritePDF, "write-pdf", "w", "", "The file to write as a PDF.")
}
