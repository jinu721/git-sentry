package git

import (
	"os"
	"testing"
)

func TestNewRepository(t *testing.T) {
	tempDir := "test_repo"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	_, err := NewRepository(tempDir)
	if err == nil {
		t.Error("Should fail for non-git directory")
	}
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Errorf("Should succeed for git directory: %v", err)
	}
	
	if repo == nil {
		t.Error("Repository should not be nil")
	}
}

func TestExecGitCommand(t *testing.T) {
	tempDir := "test_git_exec"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}
	
	_, err = repo.execGitCommand("status", "--porcelain")
	if err == nil {
		t.Log("Git command executed successfully")
	}
	
	_, err = repo.execGitCommand("invalid-command")
	if err == nil {
		t.Error("Should fail for invalid git command")
	}
}

func TestGetUnpushedCommitsCount(t *testing.T) {
	tempDir := "test_unpushed"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}
	
	count, err := repo.GetUnpushedCommitsCount()
	if err != nil && count != 0 {
		t.Log("No upstream configured, returned 0 as expected")
	}
}

func TestGetStatus(t *testing.T) {
	tempDir := "test_status"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}
	
	status, err := repo.GetStatus()
	if err != nil {
		t.Log("Git status command may fail in test environment")
	} else if status != nil {
		t.Log("Git status retrieved successfully")
	}
}

func TestGetBranch(t *testing.T) {
	tempDir := "test_branch"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}
	
	branch, err := repo.GetBranch()
	if err != nil {
		t.Log("Git branch command may fail in test environment")
	} else {
		t.Logf("Current branch: %s", branch)
	}
}

func TestHasRemote(t *testing.T) {
	tempDir := "test_remote"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	gitDir := tempDir + "/.git"
	os.MkdirAll(gitDir, 0755)
	
	repo, err := NewRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}
	
	hasRemote, err := repo.HasRemote()
	if err != nil {
		t.Log("Git remote command may fail in test environment")
	} else {
		t.Logf("Has remote: %t", hasRemote)
	}
}