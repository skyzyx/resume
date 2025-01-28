package cmd

import (
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Just run the web server with ./render as the root.",
	Long: clihelpers.LongHelpText(`
    Just run the web server with ./render as the root.

    This is useful for tweaking CSS styles and seeing the changes in real-time.`),
	Run: func(cmd *cobra.Command, args []string) {
		// Start a local web server.
		server := &http.Server{
			Addr: ":11235",
		}

		http.Handle("/", http.FileServer(http.Dir("./render")))
		logger.Debug("Starting HTTP server on :11235")

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
