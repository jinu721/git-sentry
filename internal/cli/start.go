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
	Use:   "start",
	Short: "Start GitSentry monitoring",
	Long: `Start GitSentry background monitoring of file changes.
GitSentry will watch for changes and suggest commits based on your configured rules.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		if daemonMode {
			return sentry.StartDaemon()
		}
		
		if err := sentry.Start(); err != nil {
			return fmt.Errorf("failed to start GitSentry: %w", err)
		}
		
		fmt.Println("⭐ GitSentry is now monitoring your repository")
		fmt.Println("Press Ctrl+C to stop")
		
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		
		fmt.Println("\n🛑 Stopping GitSentry...")
		sentry.Stop()
		fmt.Println("GitSentry stopped")
		
		return nil
	},
}

func init() {
	startCmd.Flags().BoolVar(&daemonMode, "daemon", false, "Run in background daemon mode")
}