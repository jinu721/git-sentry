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
		
		fmt.Println("📊 GitSentry Status")
		fmt.Println("==================")
		fmt.Printf("Repository: %s\n", status.RepoPath)
		fmt.Printf("Git initialized: %t\n", status.IsGitRepo)
		fmt.Printf("Monitoring: %t\n", status.IsMonitoring)
		fmt.Printf("Files changed: %d\n", status.FilesChanged)
		fmt.Printf("Lines added: %d\n", status.LinesAdded)
		fmt.Printf("Lines removed: %d\n", status.LinesRemoved)
		fmt.Printf("Last commit: %s\n", status.LastCommit)
		fmt.Printf("Last push: %s\n", status.LastPush)
		fmt.Printf("Unpushed commits: %d\n", status.UnpushedCommits)
		
		return nil
	},
}