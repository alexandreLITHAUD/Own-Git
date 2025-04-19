/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the status of the files in the current directory",
	Example: `own-git status`,
	Aliases: []string{"st"},
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		noColor, err := cmd.Flags().GetBool("no-color")
		if err != nil {
			noColor = false
		}

		files, err := paths.GetAllFiles(".")
		if err != nil {
			cmd.Println("Error getting files:", err)
			return
		}

		var fileStatuses map[string]uint8 = make(map[string]uint8)
		for _, file := range files {
			intStatuses, err := utils.GetFileStatus(file)
			if err != nil {
				cmd.Println("Error getting file status:", err)
				continue
			}
			fileStatuses[file] = intStatuses
		}

		// Collect file paths
		filePaths := make([]string, 0, len(fileStatuses))
		for file := range fileStatuses {
			filePaths = append(filePaths, file)
		}

		// Sort file paths by their associated status (uint8)
		sort.Slice(filePaths, func(i, j int) bool {
			return fileStatuses[filePaths[i]] < fileStatuses[filePaths[j]]
		})

		fmt.Printf("Files getting gited:\n\n")
		for _, file := range filePaths {
			status := fileStatuses[file]
			strStatus, color := utils.GetFileStatusString(status)
			if noColor {
				fmt.Printf("%s: %s\n", file, strStatus)
			} else {
				fmt.Printf("%s %s: %s %s\n", color, strStatus, file, utils.NoColor)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolP("no-color", "n", false, "Disable color output")
}
