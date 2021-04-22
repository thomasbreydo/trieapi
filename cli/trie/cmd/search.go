package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

var Search = &cobra.Command{
	Use:               "search",
	Short:             "Search for a keyword in the trie",
	Long:              "Search for a keyword in the trie",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, code, err := api.WithWord("search", input)
		if err != nil {
			return err
		}
		if code < 200 || code >= 300 {
			return errors.New(fmt.Sprintf("status code %d", code))
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	Search.Flags().StringVarP(&input, "word", "w", "", "word to search for")
	_ = Search.MarkFlagRequired("word")
}
