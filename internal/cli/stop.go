package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop GitSentry monitoring",
	Long:  `Stop GitSentry background monitoring if it's currently running.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		if err := sentry.Stop(); err != nil {
			return fmt.Errorf("failed to stop GitSentry: %w", err)
		}
		
		PrintSuccess("GitSentry stopped")
		return nil
	},
}