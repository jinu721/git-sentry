package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var (
	initTemplate string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize GitSentry in the current directory",
	Long: `Initialize GitSentry monitoring in the current directory.
This will create the .gitsentry configuration folder and check if Git is properly set up.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		if err := sentry.InitializeWithTemplate(initTemplate); err != nil {
			return fmt.Errorf("failed to initialize GitSentry: %w", err)
		}
		
		PrintSuccess("GitSentry initialized successfully!")
		if initTemplate != "" {
			PrintInfo(fmt.Sprintf("Applied template: %s", initTemplate))
		}
		PrintInfo("Use 'gitsentry start' to begin monitoring")
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initTemplate, "template", "", "Configuration template (team, strict, relaxed)")
}