// Based on examples in Cobra's documentation (see https://github.com/spf13/cobra)
package main

import (
	"os"

	"github.com/thomasbreydo/trieapi/cli/cmd"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:               "trie",
	Short:             "A command-line interface for the trie system.",
	Long:              "A command-line interface for the trie system.",
	DisableAutoGenTag: true,
}

func init() {

	command.AddCommand(cmd.Add)
	//command.AddCommand()
	//command.AddCommand()
	//command.AddCommand()
	//command.AddCommand()
}

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
