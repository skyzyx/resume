package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"

	clihelpers "github.com/northwood-labs/cli-helpers"
)

type SearchTOML struct {
	SearchEngine  string   `toml:"search_engine"`
	SearchDomains []string `toml:"search_domains"`
	JobRoles      []string `toml:"job_roles"`
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Generate a list of search URLs for specific phrases.",
	Long: clihelpers.LongHelpText(`
    Generate a list of search URLs for specific phrases.

    Reads information from searches.toml to compose a set of URLs.`),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("searches.toml")
		if err != nil {
			logger.Errorf("Error reading searches.toml: %v", err)
		}

		var cfg SearchTOML

		err = toml.Unmarshal(data, &cfg)
		if err != nil {
			logger.Errorf("Error unmarshalling searches.toml: %v", err)
		}

		for _, role := range cfg.JobRoles {
			for _, domain := range cfg.SearchDomains {
				fmt.Println(cfg.SearchEngine + url.QueryEscape(fmt.Sprintf("%s site:%s", role, domain)))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
