package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

var Complete = &cobra.Command{
	Use:               "complete",
	Short:             "Generate completions for a prefix",
	Long:              "Generate completions for a prefix",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, code, err := api.WithWord("complete", input)
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
	Complete.Flags().StringVarP(&input, "prefix", "p", "", "prefix to complete")
	_ = Complete.MarkFlagRequired("prefix")
}
