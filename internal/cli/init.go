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
	Use:   "init [flags]",
	Short: "Initialize GitSentry monitoring in current directory",
	Long: `Initialize GitSentry monitoring in the current directory.
Creates .gitsentry configuration folder and validates Git repository setup.

Available templates:
  • default  - Balanced settings for individual developers
  • team     - Stricter settings for team collaboration  
  • strict   - Very strict settings for critical projects
  • relaxed  - Relaxed settings for experimental work

Examples:
  gitsentry init                     Initialize with default settings
  gitsentry init --template=team     Initialize with team template
  gitsentry init --template=strict   Initialize with strict rules`,
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