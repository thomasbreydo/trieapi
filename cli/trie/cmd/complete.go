package cmd

import (
	"errors"
	"fmt"

	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"

	"github.com/spf13/cobra"
)

var Complete = &cobra.Command{
	Use:               "complete",
	Short:             "Generate completions for a prefix",
	Long:              "Generate completions for a prefix",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var comps string
		var code int
		var err error
		if json {
			comps, code, err = api.CompleteJSON(inp)
		} else {
			comps, code, err = api.Complete(inp)
		}
		if err != nil {
			return err
		}
		if code < 200 || code >= 300 {
			return errors.New(fmt.Sprintf("status code %d", code))
		}
		cmd.Print(comps)
		return nil
	},
}

func init() {
	Complete.Flags().StringVarP(&inp, "prefix", "p", "", "prefix to complete")
	_ = Complete.MarkFlagRequired("prefix")
	Complete.Flags().BoolVar(&json, "json", false, "use JSON output format")
}
