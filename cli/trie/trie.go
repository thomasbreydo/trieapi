// Based on examples in Cobra's documentation (see https://github.com/spf13/cobra)
package main

import (
	"os"

	"github.com/thomasbreydo/trieapi/cli/trie/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
