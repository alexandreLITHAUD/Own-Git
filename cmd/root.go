/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alexandreLITHAUD/Own-Git/internal/utils"

	"github.com/spf13/cobra"
)

var Verbose bool
var Config string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "own",
	Short: "own-git a simple git copycat using go functionalities",
	Long: `Own Git is a simple copycat of git functionalities using go.
It is a simple command line tool that allows you to create, delete, and manage your own git repositories.

It was made by Alexandre Lithaud in order to learn go and further understand how git works.
If you know how the git cli works, you will understand how this tool works.
It is not meant to be a replacement for git, but rather a learning tool.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()

		if Config != "" {
			fmt.Printf("Config file: %s\n", Config)
			fmt.Print("Config file is not yet implemented\n")
			if !utils.IsConfFileValid(Config) {
				fmt.Printf("Config file is not valid\n")
				os.Exit(1)
			}

			// TODO
			argumentConfig, err := utils.ParseConfigFile(Config)

			if err != nil {
				fmt.Printf("Error parsing config file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("argumentConfig =%v\n", argumentConfig)
		}

		if Verbose {
			fmt.Printf("Verbose mode is enabled\n")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Config, "config", "c", "", "Path to the config file")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose mode")
}
