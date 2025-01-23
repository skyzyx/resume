package cmd

import (
	"os"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Generates multiple formats of my résumé from a series of templates.",
	Long: clihelpers.LongHelpText(`
    Generates multiple formats of my résumé from a series of templates.`),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
