package cmd

import (
	"fmt"
	"net/url"
	"os"
	"slices"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"

	clihelpers "github.com/northwood-labs/cli-helpers"
)

type SearchTOML struct {
	SearchEngine     string   `toml:"search_engine"`
	SearchDomains    []string `toml:"search_domains"`
	JobRoles         []string `toml:"job_roles"`
	ExcludedKeywords []string `toml:"excluded_keywords"`
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
			i := 0

			for sDomains := range slices.Chunk(cfg.SearchDomains, 4) {
				domains := []string{}
				i++

				for _, domain := range sDomains {
					domains = append(domains, fmt.Sprintf("site:%s", domain))
				}

				searchURL := cfg.SearchEngine + url.QueryEscape(
					fmt.Sprintf(
						`%s "%s" AND "remote" -%s`,
						strings.Join(domains, " | "),
						role,
						strings.Join(cfg.ExcludedKeywords, ` -`),
					),
				)

				fmt.Printf("* [%s %d](%s)\n", role, i, searchURL)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
