package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

var Display = &cobra.Command{
	Use:               "display",
	Short:             "Display the trie",
	Long:              "Display the trie",
	Args:              cobra.NoArgs,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if json {
			ww, code, err := api.DisplayJSON()
			if err != nil {
				return err
			}
			if code < 200 || code >= 300 {
				return errors.New(fmt.Sprintf("status code %d", code))
			}
			fmt.Print(ww)
		} else {
			ww, code, err := api.Display()
			if err != nil {
				return err
			}
			if code < 200 || code >= 300 {
				return errors.New(fmt.Sprintf("status code %d", code))
			}
			fmt.Print(ww)
		}
		return nil
	},
}

func init() {
	Display.Flags().BoolVar(&json, "json", false, "use JSON output format")
}
