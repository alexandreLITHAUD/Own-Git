//go:build generateDocs

package main

import (
	"log"
	"path/filepath"

	"github.com/alexandreLITHAUD/Own-Git/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, filepath.Join("docs", "content"))
	if err != nil {
		log.Fatalf("failed to generate docs: %v", err)
	}
}
