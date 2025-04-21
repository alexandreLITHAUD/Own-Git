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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file to index before committing",
	Example: `own add --file myfile.txt
own add --all`,
	Aliases: []string{"a"},
	Long: `Add file to own-git index in order to be added to next commit.
The --file flag can be used to add a specific file to the index.
The --all flag will add all files in the current directory to the index.
The index will serve as a "staging area" for the next commit.`,
	Run: func(cmd *cobra.Command, args []string) {

		if !utils.IsOwnFolder() {
			fmt.Println("This is not an own-git repository. Please run 'own-git init' first. Or cd into an existing own-git repository.")
			return
		}

		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			filepath = ""
		}
		allBool, err := cmd.Flags().GetBool("all")
		if err != nil {
			allBool = false
		}

		if filepath == "" && !allBool {
			fmt.Println("Please specify a file using --file to add or use the --all flag to add all files.")
			return
		}

		if allBool {
			fmt.Println("Adding all files to the index...")
			files, err := paths.GetAllFiles(".")
			if err != nil {
				fmt.Println("Error getting files:", err)
				return
			}
			fileIndexEntryArr := make([]types.IndexEntry, 0)
			for _, file := range files {
				fileIndexEntry, err := utils.FilePathtoIndexEntry(file)
				if err != nil {
					fmt.Println("Error getting file index entry:", err)
					return
				}
				fileIndexEntryArr = append(fileIndexEntryArr, fileIndexEntry)
			}
			err = utils.WriteEntryToIndex(fileIndexEntryArr)
			if err != nil {
				fmt.Println("Error writing index entry:", err)
				return
			}
			fmt.Println("All files added to the index.")
			fmt.Println("You can now commit them using 'own-git commit' or see them using 'own-git status'.")
			return
		}

		if filepath != "" {
			fmt.Printf("Adding file %s to the index...\n", filepath)
			fileIndexEntry, err := utils.FilePathtoIndexEntry(filepath)
			if err != nil {
				fmt.Println("Error getting file index entry:", err)
				return
			}
			err = utils.WriteEntryToIndex([]types.IndexEntry{fileIndexEntry})
			if err != nil {
				fmt.Println("Error writing index entry:", err)
				return
			}
			fmt.Printf("File %s added to the index.\n", filepath)
			fmt.Println("You can now commit it using 'own-git commit' or see it using 'own-git status'.")
			return
		}

		fmt.Println("If you arrive here that mean you managed to find an edge cases that I didn't think about. (Good Job!)")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("file", "f", "", "File to add to the repository")
	addCmd.Flags().BoolP("all", "a", false, "Add all files to the repository")
}
