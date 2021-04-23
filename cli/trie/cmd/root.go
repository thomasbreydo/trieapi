package cmd

import "github.com/spf13/cobra"

var Root = &cobra.Command{
	Use:               "trie",
	Short:             "A command-line interface for the trie system.",
	Long:              "A command-line interface for the trie system.",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
}

func init() {
	Root.AddCommand(Add)
	Root.AddCommand(Clear)
	Root.AddCommand(Complete)
	Root.AddCommand(Delete)
	Root.AddCommand(Display)
	Root.AddCommand(Search)
	Root.CompletionOptions.DisableDefaultCmd = true
}
