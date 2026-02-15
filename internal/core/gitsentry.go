package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gitsentry/internal/config"
	"gitsentry/internal/daemon"
	"gitsentry/internal/git"
	"gitsentry/internal/monitor"
	"gitsentry/internal/security"
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
	gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
	if err := security.SecureCreateDir(gitsentryDir); err != nil {
		return fmt.Errorf("failed to create .gitsentry directory: %w", err)
	}
	
	logsDir := filepath.Join(gitsentryDir, "logs")
	if err := security.SecureCreateDir(logsDir); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}
	
	cfg, err := config.LoadWithTemplate(gitsentryDir, template)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	gs.config = cfg
	
	st, err := state.Load(gitsentryDir)
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}
	gs.state = st
	
	gitRepo, err := git.NewRepository(gs.repoPath)
	if err != nil {
		fmt.Println("Git repository not found")
		fmt.Println("Would you like to:")
		fmt.Println("1. Initialize git repository")
		fmt.Println("2. Initialize git + add GitHub remote")
		fmt.Println("3. Skip (GitSentry will work with limited functionality)")
		
		fmt.Println("Continuing with limited functionality...")
	} else {
		gs.gitRepo = gitRepo
	}
	
	if err := gs.addToGitignore(); err != nil {
		return fmt.Errorf("failed to update .gitignore: %w", err)
	}
	
	return nil
}

func (gs *GitSentry) StartDaemon() error {
	d := daemon.NewDaemon(gs.repoPath)
	
	if err := d.Daemonize(); err != nil {
		return fmt.Errorf("failed to start daemon: %w", err)
	}
	
	if err := gs.Start(); err != nil {
		d.RemovePID()
		return fmt.Errorf("failed to start monitoring: %w", err)
	}
	
	fmt.Println("GitSentry daemon started successfully")
	fmt.Printf("Monitoring: %s\n", gs.repoPath)
	
	select {}
}

func (gs *GitSentry) Start() error {
	if gs.isRunning {
		return fmt.Errorf("GitSentry is already running")
	}
	
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
	
	if gs.gitRepo == nil {
		gitRepo, err := git.NewRepository(gs.repoPath)
		if err == nil {
			gs.gitRepo = gitRepo
		}
	}
	
	monitor, err := monitor.NewFileMonitor(gs.repoPath, gs.onFileChange)
	if err != nil {
		return fmt.Errorf("failed to start file monitor: %w", err)
	}
	gs.monitor = monitor
	
	gs.isRunning = true
	
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
		filesChanged, linesAdded, linesRemoved, lastCommit, lastPush := gs.state.GetStats()
		status.FilesChanged = filesChanged
		status.LinesAdded = linesAdded
		status.LinesRemoved = linesRemoved
		
		if !lastCommit.IsZero() {
			status.LastCommit = lastCommit.Format("2006-01-02 15:04:05")
		} else {
			status.LastCommit = "Never"
		}
		
		if !lastPush.IsZero() {
			status.LastPush = lastPush.Format("2006-01-02 15:04:05")
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

func (gs *GitSentry) SaveConfig(config *config.Config) error {
	gitsentryDir := filepath.Join(gs.repoPath, ".gitsentry")
	
	if err := config.Save(gitsentryDir); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}
	
	gs.config = config
	return nil
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
	
	gs.state.IncrementFilesChanged()
	
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
	
	filesChanged, linesAdded, linesRemoved, lastCommit, _ := gs.state.GetStats()
	
	shouldSuggest := false
	
	if filesChanged >= gs.config.Rules.MaxFilesChanged {
		shouldSuggest = true
	}
	
	if linesAdded+linesRemoved >= gs.config.Rules.MaxLinesChanged {
		shouldSuggest = true
	}
	
	if time.Since(lastCommit).Minutes() >= float64(gs.config.Rules.MaxMinutesSinceCommit) {
		shouldSuggest = true
	}
	
	if shouldSuggest {
		fmt.Println("\nGitSentry suggests it's a good time to commit!")
		fmt.Printf("   Files changed: %d\n", filesChanged)
		fmt.Printf("   Lines changed: %d\n", linesAdded+linesRemoved)
		fmt.Printf("   Time since last commit: %.0f minutes\n", time.Since(lastCommit).Minutes())
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
		fmt.Println("\nGitSentry suggests pushing your commits for backup!")
		fmt.Printf("   Unpushed commits: %d\n", unpushed)
		fmt.Println("   Run 'git push' when ready")
	}
}

func (gs *GitSentry) addToGitignore() error {
	gitignorePath := filepath.Join(gs.repoPath, ".gitignore")
	
	if _, err := os.Stat(gitignorePath); err == nil {
		content, err := os.ReadFile(gitignorePath)
		if err != nil {
			return err
		}
		
		if string(content) != "" && !contains(string(content), ".gitsentry/") {
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