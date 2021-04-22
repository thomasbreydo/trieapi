package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

var Delete = &cobra.Command{
	Use:               "delete",
	Short:             "Delete a keyword from the trie",
	Long:              "Delete a keyword from the trie",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, code, err := api.WithWord("delete", inp)
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
	Delete.Flags().StringVarP(&inp, "word", "w", "", "word to delete")
	_ = Delete.MarkFlagRequired("word")
}
