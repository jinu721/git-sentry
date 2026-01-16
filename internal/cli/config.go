package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage GitSentry configuration",
	Long:  `View and modify GitSentry configuration settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		config, err := sentry.GetConfig()
		if err != nil {
			return fmt.Errorf("failed to get config: %w", err)
		}
		
		fmt.Println("⚙️  GitSentry Configuration")
		fmt.Println("==========================")
		fmt.Printf("Max files changed: %d\n", config.Rules.MaxFilesChanged)
		fmt.Printf("Max lines changed: %d\n", config.Rules.MaxLinesChanged)
		fmt.Printf("Max minutes since commit: %d\n", config.Rules.MaxMinutesSinceCommit)
		fmt.Printf("Max unpushed commits: %d\n", config.Rules.MaxUnpushedCommits)
		fmt.Printf("Auto-suggest commits: %t\n", config.AutoSuggestCommits)
		fmt.Printf("Auto-suggest pushes: %t\n", config.AutoSuggestPushes)
		
		return nil
	},
}