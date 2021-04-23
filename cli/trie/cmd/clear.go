package cmd

import (
	"errors"
	"fmt"

	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"

	"github.com/spf13/cobra"
)

var Clear = &cobra.Command{
	Use:               "clear",
	Short:             "Clear the trie",
	Long:              "Clear the trie",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		mod, code, err := api.Clear()
		if err != nil {
			return err
		}
		if code < 200 || code >= 300 {
			return errors.New(fmt.Sprintf("status code %d", code))
		}
		if mod {
			cmd.Print("Trie cleared")
		} else {
			cmd.Print("Already empty")
		}
		return nil
	},
}
