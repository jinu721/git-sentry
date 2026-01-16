package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Repository struct {
	path string
}

func NewRepository(path string) (*Repository, error) {
	// Check if .git directory exists
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("not a git repository")
	}
	
	return &Repository{path: path}, nil
}

func (r *Repository) GetUnpushedCommitsCount() (int, error) {
	cmd := exec.Command("git", "rev-list", "--count", "@{u}..HEAD")
	cmd.Dir = r.path
	
	output, err := cmd.Output()
	if err != nil {
		// If there's no upstream, return 0
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
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = r.path
	
	output, err := cmd.Output()
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
	cmd := exec.Command("git", "log", "-1", "--format=%ci")
	cmd.Dir = r.path
	
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}

func (r *Repository) HasRemote() (bool, error) {
	cmd := exec.Command("git", "remote")
	cmd.Dir = r.path
	
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	
	return strings.TrimSpace(string(output)) != "", nil
}

func (r *Repository) GetBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = r.path
	
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}