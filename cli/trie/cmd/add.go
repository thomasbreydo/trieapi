package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

var Add = &cobra.Command{
	Use:               "add",
	Short:             "Add a keyword to the trie",
	Long:              "Add a keyword to the trie",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		mod, code, err := api.Add(inp)
		if err != nil {
			return err
		}
		if code < 200 || code >= 300 {
			return errors.New(fmt.Sprintf("status code %d", code))
		}
		if mod { // modified because not found
			fmt.Printf("Keyword (%s) not found", inp)
		} else {
			fmt.Printf("Keyword (%s) found", inp)
		}
		return nil
	},
}

func init() {
	Add.Flags().StringVarP(&inp, "word", "w", "", "word to add")
	_ = Add.MarkFlagRequired("word")
}
