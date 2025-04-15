/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// This variable will be set at build time through the -X linker flag
var Version string = "undefined"

// addCmd represents the add command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show the version of the application",
	Aliases: []string{"v"},
	Example: "own version",
	Long: `Show the version of the application.
	
Example :
  own version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("own-git version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
