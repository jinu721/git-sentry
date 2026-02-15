package security

import (
	"testing"
)

func TestValidateGitCommand(t *testing.T) {
	validCommands := [][]string{
		{"status", "--porcelain"},
		{"log", "--oneline"},
		{"branch", "--show-current"},
		{"remote"},
		{"rev-list", "--count", "@{u}..HEAD"},
	}
	
	for _, cmd := range validCommands {
		if err := ValidateGitCommand(cmd); err != nil {
			t.Errorf("Expected valid command %v to pass validation: %v", cmd, err)
		}
	}
	
	invalidCommands := [][]string{
		{"push"},
		{"pull"},
		{"commit"},
		{"add"},
		{"reset", "--hard"},
		{"checkout"},
		{},
	}
	
	for _, cmd := range invalidCommands {
		if err := ValidateGitCommand(cmd); err == nil {
			t.Errorf("Expected invalid command %v to fail validation", cmd)
		}
	}
}

func TestSanitizeGitArgs(t *testing.T) {
	validArgs := []string{"status", "--porcelain"}
	
	result, err := SanitizeGitArgs(validArgs)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if len(result) != len(validArgs) {
		t.Errorf("Expected %d args, got %d", len(validArgs), len(result))
	}
	
	invalidArgs := []string{"push", "origin"}
	_, err = SanitizeGitArgs(invalidArgs)
	if err == nil {
		t.Error("Expected error for invalid args")
	}
}