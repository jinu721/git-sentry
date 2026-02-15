package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitsentry",
	Short: "GitSentry - Your intelligent Git workflow assistant",
	Long: `GitSentry is an intelligent Git workflow assistant that monitors your repository
in real-time and provides smart suggestions for commits, pushes, and best practices.

Features:
  • Real-time file monitoring with smart commit suggestions
  • Configurable rules for different team workflows  
  • Background daemon mode for continuous monitoring
  • Interactive configuration management
  • Comprehensive diagnostics and health checks
  • Statistics export for productivity analysis

Examples:
  gitsentry init --template=team     Initialize with team configuration
  gitsentry start --daemon           Start monitoring in background
  gitsentry rules --interactive      Configure rules interactively
  gitsentry stats --export=json      Export statistics to JSON
  gitsentry doctor                   Run health diagnostics`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintHeader("GitSentry - Your Git Workflow Assistant")
		PrintInfo("Use 'gitsentry --help' to see all available commands")
		PrintInfo("Use 'gitsentry <command> --help' for detailed command help")
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
	rootCmd.AddCommand(doctorCmd)
}