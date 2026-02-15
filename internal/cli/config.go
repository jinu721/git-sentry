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
		
		PrintHeader("GitSentry Configuration")
		
		fmt.Println(FormatKeyValue("Max files changed", fmt.Sprintf("%d", config.Rules.MaxFilesChanged)))
		fmt.Println(FormatKeyValue("Max lines changed", fmt.Sprintf("%d", config.Rules.MaxLinesChanged)))
		fmt.Println(FormatKeyValue("Max minutes since commit", fmt.Sprintf("%d", config.Rules.MaxMinutesSinceCommit)))
		fmt.Println(FormatKeyValue("Max unpushed commits", fmt.Sprintf("%d", config.Rules.MaxUnpushedCommits)))
		fmt.Println(FormatKeyValue("Auto-suggest commits", FormatBool(config.AutoSuggestCommits)))
		fmt.Println(FormatKeyValue("Auto-suggest pushes", FormatBool(config.AutoSuggestPushes)))
		
		return nil
	},
}