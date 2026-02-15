package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show GitSentry status and repository information",
	Long:  `Display current GitSentry status, repository state, and monitoring statistics.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		status, err := sentry.GetStatus()
		if err != nil {
			return fmt.Errorf("failed to get status: %w", err)
		}
		
		PrintHeader("GitSentry Status")
		
		fmt.Println(FormatKeyValue("Repository", status.RepoPath))
		fmt.Println(FormatKeyValue("Git initialized", FormatBool(status.IsGitRepo)))
		fmt.Println(FormatKeyValue("Monitoring", FormatStatus(status.IsMonitoring, "Active", "Inactive")))
		fmt.Println(FormatKeyValue("Files changed", fmt.Sprintf("%d", status.FilesChanged)))
		fmt.Println(FormatKeyValue("Lines added", fmt.Sprintf("%d", status.LinesAdded)))
		fmt.Println(FormatKeyValue("Lines removed", fmt.Sprintf("%d", status.LinesRemoved)))
		fmt.Println(FormatKeyValue("Last commit", status.LastCommit))
		fmt.Println(FormatKeyValue("Last push", status.LastPush))
		fmt.Println(FormatKeyValue("Unpushed commits", fmt.Sprintf("%d", status.UnpushedCommits)))
		
		return nil
	},
}