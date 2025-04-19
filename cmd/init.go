/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alexandreLITHAUD/Own-Git/internal/utils"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize own-git folder",
	Example: `own init --initial-branch ILoveSourceControl`,
	Aliases: []string{"i"},
	Long: `Initialize own-git folder. And create multiples folders for objects and refs.
You can specify the name of the initial branch, using the --initial-branch flag.
The default value is "main".`,
	Run: func(cmd *cobra.Command, args []string) {
		branch, _ := cmd.Flags().GetString("initial-branch")
		if Verbose {
			fmt.Printf("Creating own-git folder with initial branch: %s\n", branch)
			if Config != "" {
				fmt.Printf("Config file: %s\n", Config)
			}
		}
		utils.CreateOwnFolder(branch, Config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("initial-branch", "b", "main", "Name of the initial branch")
}
