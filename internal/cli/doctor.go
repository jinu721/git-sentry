package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitsentry/internal/core"
	"gitsentry/internal/daemon"
)

type DiagnosticResult struct {
	Name    string
	Status  string
	Message string
	Success bool
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run GitSentry diagnostics",
	Long:  `Diagnose GitSentry installation and configuration issues.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("GitSentry Health Check")
		fmt.Println("=====================")
		
		results := runDiagnostics()
		
		allGood := true
		for _, result := range results {
			status := "FAIL"
			if result.Success {
				status = "PASS"
			} else {
				allGood = false
			}
			
			fmt.Printf("[%s] %s: %s\n", status, result.Name, result.Message)
		}
		
		fmt.Println()
		if allGood {
			fmt.Println("All checks passed! GitSentry is ready to use.")
		} else {
			fmt.Println("Some issues found. Please address them for optimal performance.")
		}
		
		return nil
	},
}

func runDiagnostics() []DiagnosticResult {
	var results []DiagnosticResult
	
	results = append(results, checkGitAvailability())
	results = append(results, checkWorkingDirectory())
	results = append(results, checkGitRepository())
	results = append(results, checkConfigValidity())
	results = append(results, checkFileWatcher())
	results = append(results, checkPermissions())
	results = append(results, checkDaemonStatus())
	
	return results
}

func checkGitAvailability() DiagnosticResult {
	cmd := exec.Command("git", "--version")
	output, err := cmd.Output()
	
	if err != nil {
		return DiagnosticResult{
			Name:    "Git Installation",
			Status:  "FAIL",
			Message: "Git not found in PATH",
			Success: false,
		}
	}
	
	return DiagnosticResult{
		Name:    "Git Installation",
		Status:  "PASS",
		Message: string(output[:len(output)-1]),
		Success: true,
	}
}

func checkWorkingDirectory() DiagnosticResult {
	wd, err := os.Getwd()
	if err != nil {
		return DiagnosticResult{
			Name:    "Working Directory",
			Status:  "FAIL",
			Message: "Cannot determine working directory",
			Success: false,
		}
	}
	
	return DiagnosticResult{
		Name:    "Working Directory",
		Status:  "PASS",
		Message: wd,
		Success: true,
	}
}

func checkGitRepository() DiagnosticResult {
	gitDir := filepath.Join(".", ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return DiagnosticResult{
			Name:    "Git Repository",
			Status:  "WARN",
			Message: "Not a Git repository",
			Success: false,
		}
	}
	
	return DiagnosticResult{
		Name:    "Git Repository",
		Status:  "PASS",
		Message: "Valid Git repository detected",
		Success: true,
	}
}

func checkConfigValidity() DiagnosticResult {
	sentry := core.NewGitSentry(".")
	_, err := sentry.GetConfig()
	
	if err != nil {
		return DiagnosticResult{
			Name:    "Configuration",
			Status:  "FAIL",
			Message: fmt.Sprintf("Config error: %v", err),
			Success: false,
		}
	}
	
	return DiagnosticResult{
		Name:    "Configuration",
		Status:  "PASS",
		Message: "Configuration is valid",
		Success: true,
	}
}

func checkFileWatcher() DiagnosticResult {
	gitsentryDir := filepath.Join(".", ".gitsentry")
	if _, err := os.Stat(gitsentryDir); os.IsNotExist(err) {
		return DiagnosticResult{
			Name:    "File Watcher",
			Status:  "WARN",
			Message: "GitSentry not initialized (run 'gitsentry init')",
			Success: false,
		}
	}
	
	return DiagnosticResult{
		Name:    "File Watcher",
		Status:  "PASS",
		Message: "Ready for file monitoring",
		Success: true,
	}
}

func checkPermissions() DiagnosticResult {
	testFile := filepath.Join(".", ".gitsentry", "test_permissions")
	
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return DiagnosticResult{
			Name:    "Permissions",
			Status:  "FAIL",
			Message: "Cannot write to .gitsentry directory",
			Success: false,
		}
	}
	
	os.Remove(testFile)
	
	return DiagnosticResult{
		Name:    "Permissions",
		Status:  "PASS",
		Message: "Read/write permissions OK",
		Success: true,
	}
}

func checkDaemonStatus() DiagnosticResult {
	d := daemon.NewDaemon(".")
	
	if d.IsRunning() {
		return DiagnosticResult{
			Name:    "Daemon Status",
			Status:  "INFO",
			Message: "GitSentry daemon is running",
			Success: true,
		}
	}
	
	return DiagnosticResult{
		Name:    "Daemon Status",
		Status:  "INFO",
		Message: "GitSentry daemon is not running",
		Success: true,
	}
}