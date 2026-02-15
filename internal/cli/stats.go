package cli

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
	"gitsentry/internal/security"
)

var (
	exportFormat string
	outputFile   string
)

type StatsExport struct {
	Timestamp       time.Time `json:"timestamp"`
	RepoPath        string    `json:"repo_path"`
	IsGitRepo       bool      `json:"is_git_repo"`
	IsMonitoring    bool      `json:"is_monitoring"`
	FilesChanged    int       `json:"files_changed"`
	LinesAdded      int       `json:"lines_added"`
	LinesRemoved    int       `json:"lines_removed"`
	LastCommit      string    `json:"last_commit"`
	LastPush        string    `json:"last_push"`
	UnpushedCommits int       `json:"unpushed_commits"`
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Display and export GitSentry statistics",
	Long:  `Show current GitSentry statistics and optionally export to JSON format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sentry := core.NewGitSentry(".")
		
		status, err := sentry.GetStatus()
		if err != nil {
			return fmt.Errorf("failed to get status: %w", err)
		}
		
		if exportFormat == "json" {
			return exportStatsJSON(status)
		}
		
		fmt.Println("GitSentry Statistics")
		fmt.Println("===================")
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

func exportStatsJSON(status *core.Status) error {
	export := StatsExport{
		Timestamp:       time.Now(),
		RepoPath:        status.RepoPath,
		IsGitRepo:       status.IsGitRepo,
		IsMonitoring:    status.IsMonitoring,
		FilesChanged:    status.FilesChanged,
		LinesAdded:      status.LinesAdded,
		LinesRemoved:    status.LinesRemoved,
		LastCommit:      status.LastCommit,
		LastPush:        status.LastPush,
		UnpushedCommits: status.UnpushedCommits,
	}
	
	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	
	if outputFile != "" {
		if err := security.SecureWriteFile(outputFile, data); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
		
		absPath, _ := filepath.Abs(outputFile)
		fmt.Printf("Statistics exported to: %s\n", absPath)
	} else {
		fmt.Println(string(data))
	}
	
	return nil
}

func init() {
	statsCmd.Flags().StringVar(&exportFormat, "export", "", "Export format (json)")
	statsCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path")
}