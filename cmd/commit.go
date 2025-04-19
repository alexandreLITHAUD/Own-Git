/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/alexandreLITHAUD/Own-Git/internal/utils"

	"github.com/spf13/cobra"
)

var commitMessage string
var commitAuthor string
var commitDate string

// TODO
var commitBranch string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to the repository",
	Long: `Commit changes to the repository. This command will create a new commit with the changes made to the repository.
You can specify the commit message using the --message flag.
The default value is "Commit".
Example :
  own commit --message "My first commit" --author "John Doe" --date "2025-01-01T00:00:00Z"`,
	Run: func(cmd *cobra.Command, args []string) {

		if !utils.IsOwnFolder() {
			fmt.Println("error: .own-git folder does not exist")
			return
		}
		fmt.Println("commit called")

		// Get the current branch name
		commitBranch, err := utils.GetBranchName()
		if err != nil {
			fmt.Printf("Branch name not found: %v\n", err)
			return
		}

		fmt.Printf("Commit message: %s\nCommit author: %s\nCommit date: %s\nBranch name: %s\n", commitMessage, commitAuthor, commitDate, commitBranch)

	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "Commit", "Commit message")
	commitCmd.Flags().StringVarP(&commitAuthor, "author", "a", os.Getenv("USER"), "Author of the commit")
	commitCmd.Flags().StringVarP(&commitDate, "date", "d", time.Now().Format(time.RFC3339), "Date of the commit in RFC3339 format")
}
