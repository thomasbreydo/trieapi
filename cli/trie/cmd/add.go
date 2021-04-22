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
		res, code, err := api.WithWord("add", inp)
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
	Add.Flags().StringVarP(&inp, "word", "w", "", "word to add")
	_ = Add.MarkFlagRequired("word")
}
