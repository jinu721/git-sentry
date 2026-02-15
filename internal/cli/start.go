package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
)

var (
	daemonMode bool
)

var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "Start GitSentry file monitoring",
	Long: `Start GitSentry background monitoring of file changes.
Watches for file modifications and suggests commits based on configured rules.

Modes:
  • Interactive - Runs in foreground with live feedback (default)
  • Daemon     - Runs in background as system service

Examples:
  gitsentry start                    Start interactive monitoring
  gitsentry start --daemon           Start background daemon mode`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		if daemonMode {
			return sentry.StartDaemon()
		}
		
		if err := sentry.Start(); err != nil {
			return fmt.Errorf("failed to start GitSentry: %w", err)
		}
		
		PrintInfo("GitSentry is now monitoring your repository")
		PrintInfo("Press Ctrl+C to stop")
		
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		
		PrintInfo("Stopping GitSentry...")
		sentry.Stop()
		PrintSuccess("GitSentry stopped")
		
		return nil
	},
}

func init() {
	startCmd.Flags().BoolVar(&daemonMode, "daemon", false, "Run in background daemon mode")
}