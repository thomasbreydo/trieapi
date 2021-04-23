// Automatically generate a Markdown reference for the CLI
package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra/doc"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd"
)

func main() {
	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	err := doc.GenMarkdownTree(cmd.Root, dir)
	if err != nil {
		log.Fatal(err)
	}
}
