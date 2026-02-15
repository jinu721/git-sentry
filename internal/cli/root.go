package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitsentry",
	Short: "GitSentry - Your local-first Git assistant",
	Long: `GitSentry is a background Git assistant that monitors your project folder
in real time and helps enforce clean Git habits without being intrusive.

It helps you know when to commit, write standard commit messages,
and safely push code for backup - all while keeping you in control.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GitSentry - Your Git mentor")
		fmt.Println("Use 'gitsentry --help' to see available commands")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(rulesCmd)
	rootCmd.AddCommand(statsCmd)
}