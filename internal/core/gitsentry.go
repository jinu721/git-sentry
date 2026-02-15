package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gitsentry/internal/config"
	"gitsentry/internal/git"
	"gitsentry/internal/monitor"
	"gitsentry/internal/state"
)

type GitSentry struct {
	repoPath    string
	config      *config.Config
	state       *state.State
	gitRepo     *git.Repository
	monitor     *monitor.FileMonitor
	isRunning   bool
}

type Status struct {
	RepoPath        string
	IsGitRepo       bool
	IsMonitoring    bool
	FilesChanged    int
	LinesAdded      int
	LinesRemoved    int
	LastCommit      string
	LastPush        string
	UnpushedCommits int
}

func NewGitSentry(repoPath string) *GitSentry {
	return &GitSentry{
		repoPath: repoPath,
	}
}

func (gs *GitSentry) Initialize() error {
	return gs.InitializeWithTemplate("")
}

func (gs *GitSentry) InitializeWithTemplate(template string) error {
	// Create .gitsentry directory
	gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
	if err := os.MkdirAll(gitsentryDir, 0755); err != nil {
		return fmt.Errorf("failed to create .gitsentry directory: %w", err)
	}
	
	// Create logs directory
	logsDir := filepath.Join(gitsentryDir, "logs")
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}
	
	// Initialize configuration
	cfg, err := config.LoadWithTemplate(gitsentryDir, template)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	gs.config = cfg
	
	// Initialize state
	st, err := state.Load(gitsentryDir)
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}
	gs.state = st
	
	// Check if Git repository exists
	gitRepo, err := git.NewRepository(gs.repoPath)
	if err != nil {
		fmt.Println("⚠️  Git repository not found")
		fmt.Println("Would you like to:")
		fmt.Println("1. Initialize git repository")
		fmt.Println("2. Initialize git + add GitHub remote")
		fmt.Println("3. Skip (GitSentry will work with limited functionality)")
		
		// For now, just warn - in a real implementation, we'd prompt for user input
		fmt.Println("Continuing with limited functionality...")
	} else {
		gs.gitRepo = gitRepo
	}
	
	// Add .gitsentry to .gitignore
	if err := gs.addToGitignore(); err != nil {
		return fmt.Errorf("failed to update .gitignore: %w", err)
	}
	
	return nil
}

func (gs *GitSentry) Start() error {
	if gs.isRunning {
		return fmt.Errorf("GitSentry is already running")
	}
	
	// Load configuration and state if not already loaded
	if gs.config == nil {
		gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
		cfg, err := config.Load(gitsentryDir)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		gs.config = cfg
	}
	
	if gs.state == nil {
		gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
		st, err := state.Load(gitsentryDir)
		if err != nil {
			return fmt.Errorf("failed to load state: %w", err)
		}
		gs.state = st
	}
	
	// Initialize Git repository if not already done
	if gs.gitRepo == nil {
		gitRepo, err := git.NewRepository(gs.repoPath)
		if err == nil {
			gs.gitRepo = gitRepo
		}
	}
	
	// Start file monitoring
	monitor, err := monitor.NewFileMonitor(gs.repoPath, gs.onFileChange)
	if err != nil {
		return fmt.Errorf("failed to start file monitor: %w", err)
	}
	gs.monitor = monitor
	
	gs.isRunning = true
	
	// Start monitoring in background
	go gs.monitorLoop()
	
	return nil
}

func (gs *GitSentry) Stop() error {
	if !gs.isRunning {
		return nil
	}
	
	gs.isRunning = false
	
	if gs.monitor != nil {
		gs.monitor.Stop()
	}
	
	return nil
}

func (gs *GitSentry) GetStatus() (*Status, error) {
	status := &Status{
		RepoPath:     gs.repoPath,
		IsGitRepo:    gs.gitRepo != nil,
		IsMonitoring: gs.isRunning,
	}
	
	if gs.state != nil {
		status.FilesChanged = gs.state.FilesChanged
		status.LinesAdded = gs.state.LinesAdded
		status.LinesRemoved = gs.state.LinesRemoved
		
		if !gs.state.LastCommit.IsZero() {
			status.LastCommit = gs.state.LastCommit.Format("2006-01-02 15:04:05")
		} else {
			status.LastCommit = "Never"
		}
		
		if !gs.state.LastPush.IsZero() {
			status.LastPush = gs.state.LastPush.Format("2006-01-02 15:04:05")
		} else {
			status.LastPush = "Never"
		}
	}
	
	if gs.gitRepo != nil {
		unpushed, err := gs.gitRepo.GetUnpushedCommitsCount()
		if err == nil {
			status.UnpushedCommits = unpushed
		}
	}
	
	return status, nil
}

func (gs *GitSentry) GetConfig() (*config.Config, error) {
	if gs.config == nil {
		gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
		cfg, err := config.Load(gitsentryDir)
		if err != nil {
			return nil, err
		}
		gs.config = cfg
	}
	
	return gs.config, nil
}

func (gs *GitSentry) onFileChange(path string) {
	if gs.state == nil {
		return
	}
	
	gs.state.FilesChanged++
	gs.state.LastActivity = time.Now()
	
	// Save state
	gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
	gs.state.Save(gitsentryDir)
}

func (gs *GitSentry) monitorLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for gs.isRunning {
		select {
		case <-ticker.C:
			gs.checkCommitSuggestion()
			gs.checkPushSuggestion()
		}
	}
}

func (gs *GitSentry) checkCommitSuggestion() {
	if gs.config == nil || gs.state == nil || gs.gitRepo == nil {
		return
	}
	
	if !gs.config.AutoSuggestCommits {
		return
	}
	
	// Check rules
	shouldSuggest := false
	
	if gs.state.FilesChanged >= gs.config.Rules.MaxFilesChanged {
		shouldSuggest = true
	}
	
	if gs.state.LinesAdded+gs.state.LinesRemoved >= gs.config.Rules.MaxLinesChanged {
		shouldSuggest = true
	}
	
	if time.Since(gs.state.LastCommit).Minutes() >= float64(gs.config.Rules.MaxMinutesSinceCommit) {
		shouldSuggest = true
	}
	
	if shouldSuggest {
		fmt.Println("\n💡 GitSentry suggests it's a good time to commit!")
		fmt.Printf("   Files changed: %d\n", gs.state.FilesChanged)
		fmt.Printf("   Lines changed: %d\n", gs.state.LinesAdded+gs.state.LinesRemoved)
		fmt.Printf("   Time since last commit: %.0f minutes\n", time.Since(gs.state.LastCommit).Minutes())
		fmt.Println("   Run 'git add . && git commit' when ready")
	}
}

func (gs *GitSentry) checkPushSuggestion() {
	if gs.config == nil || gs.gitRepo == nil {
		return
	}
	
	if !gs.config.AutoSuggestPushes {
		return
	}
	
	unpushed, err := gs.gitRepo.GetUnpushedCommitsCount()
	if err != nil {
		return
	}
	
	if unpushed >= gs.config.Rules.MaxUnpushedCommits {
		fmt.Println("\n📤 GitSentry suggests pushing your commits for backup!")
		fmt.Printf("   Unpushed commits: %d\n", unpushed)
		fmt.Println("   Run 'git push' when ready")
	}
}

func (gs *GitSentry) addToGitignore() error {
	gitignorePath := filepath.Join(gs.repoPath, ".gitignore")
	
	// Check if .gitignore exists and if .gitsentry is already in it
	if _, err := os.Stat(gitignorePath); err == nil {
		content, err := os.ReadFile(gitignorePath)
		if err != nil {
			return err
		}
		
		// Check if .gitsentry is already in .gitignore
		if string(content) != "" && !contains(string(content), ".gitsentry/") {
			// Append .gitsentry to .gitignore
			f, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			defer f.Close()
			
			if _, err := f.WriteString("\n# GitSentry\n.gitsentry/\n"); err != nil {
				return err
			}
		}
	} else {
		// Create .gitignore with .gitsentry
		content := "# GitSentry\n.gitsentry/\n"
		if err := os.WriteFile(gitignorePath, []byte(content), 0644); err != nil {
			return err
		}
	}
	
	return nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}