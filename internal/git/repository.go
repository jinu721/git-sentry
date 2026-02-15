package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitsentry/internal/security"
)

type Repository struct {
	path string
}

func NewRepository(path string) (*Repository, error) {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("not a git repository")
	}
	
	return &Repository{path: path}, nil
}

func (r *Repository) GetUnpushedCommitsCount() (int, error) {
	output, err := r.execGitCommand("rev-list", "--count", "@{u}..HEAD")
	if err != nil {
		return 0, nil
	}
	
	countStr := strings.TrimSpace(string(output))
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}
	
	return count, nil
}

func (r *Repository) GetStatus() ([]string, error) {
	output, err := r.execGitCommand("status", "--porcelain")
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
	}
	
	return lines, nil
}

func (r *Repository) GetChangedFiles() ([]string, error) {
	status, err := r.GetStatus()
	if err != nil {
		return nil, err
	}
	
	var files []string
	for _, line := range status {
		if len(line) > 3 {
			files = append(files, line[3:])
		}
	}
	
	return files, nil
}

func (r *Repository) IsClean() (bool, error) {
	status, err := r.GetStatus()
	if err != nil {
		return false, err
	}
	
	return len(status) == 0, nil
}

func (r *Repository) GetLastCommitTime() (string, error) {
	output, err := r.execGitCommand("log", "-1", "--format=%ci")
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}

func (r *Repository) HasRemote() (bool, error) {
	output, err := r.execGitCommand("remote")
	if err != nil {
		return false, err
	}
	
	return strings.TrimSpace(string(output)) != "", nil
}

func (r *Repository) GetBranch() (string, error) {
	output, err := r.execGitCommand("branch", "--show-current")
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}

func (r *Repository) execGitCommand(args ...string) ([]byte, error) {
	sanitizedArgs, err := security.SanitizeGitArgs(args)
	if err != nil {
		return nil, fmt.Errorf("invalid git command: %w", err)
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	cmd := exec.CommandContext(ctx, "git", sanitizedArgs...)
	cmd.Dir = r.path
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git command failed: %w", err)
	}
	
	return output, nil
}