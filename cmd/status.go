/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/types"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
	"github.com/spf13/cobra"
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
			fmt.Println("Error getting files:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("No files found in current directory")
			return
		}

		var groupedFiles map[types.FileStatus][]string = make(map[types.FileStatus][]string)
		for _, file := range files {
			fileStatusStruct, err := utils.GetFileStatus(file)
			if err != nil {
				fmt.Println("Error getting file status:", err)
				return
			}
			groupedFiles[fileStatusStruct.Status] = append(groupedFiles[fileStatusStruct.Status], file)
		}

		fmt.Printf("All files taken in account:\n")
		fmt.Printf("\t (use 'own-git add <file>' to include in what will be committed)\n")
		fmt.Printf("\t (use 'own-git restore <file>' to discard changes in working directory)\n")
		fmt.Printf("\t One all files are good to go, use 'own-git commit' to commit them\n\n")

		for status, files := range groupedFiles {
			strStatus, color := utils.GetFileStatusString(status)
			if len(files) == 0 {
				continue
			}
			fmt.Printf("[%s]\n\n", strStatus)
			for _, file := range files {
				if noColor {
					fmt.Printf("%s: %s\n", strStatus, file)
				} else {
					fmt.Printf("%s %s: %s %s\n", color, strStatus, file, types.NoColor)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolP("no-color", "n", false, "Disable color output")
}
