// Based on examples in Cobra's documentation (see https://github.com/spf13/cobra)
package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd"
)

var command = &cobra.Command{
	Use:               "trie",
	Short:             "A command-line interface for the trie system.",
	Long:              "A command-line interface for the trie system.",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
}

func init() {
	command.AddCommand(cmd.Add)
	command.AddCommand(cmd.Clear)
	command.AddCommand(cmd.Complete)
	command.AddCommand(cmd.Delete)
	command.AddCommand(cmd.Display)
	command.AddCommand(cmd.Search)
}

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
