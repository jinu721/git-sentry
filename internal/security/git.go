package security

import (
	"fmt"
	"strings"
)

var allowedGitCommands = map[string]bool{
	"status":     true,
	"log":        true,
	"rev-list":   true,
	"branch":     true,
	"remote":     true,
	"diff":       true,
	"show":       true,
	"ls-files":   true,
	"rev-parse":  true,
}

var allowedGitFlags = map[string]bool{
	"--porcelain":     true,
	"--count":         true,
	"--show-current":  true,
	"--format":        true,
	"--oneline":       true,
	"--name-only":     true,
	"--cached":        true,
	"--short":         true,
}

func ValidateGitCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("empty git command")
	}
	
	command := args[0]
	if !allowedGitCommands[command] {
		return fmt.Errorf("git command not allowed: %s", command)
	}
	
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			if !isAllowedFlag(arg) {
				return fmt.Errorf("git flag not allowed: %s", arg)
			}
		}
	}
	
	return nil
}

func isAllowedFlag(flag string) bool {
	if allowedGitFlags[flag] {
		return true
	}
	
	for allowedFlag := range allowedGitFlags {
		if strings.HasPrefix(flag, allowedFlag+"=") {
			return true
		}
	}
	
	return false
}

func SanitizeGitArgs(args []string) ([]string, error) {
	if err := ValidateGitCommand(args); err != nil {
		return nil, err
	}
	
	sanitized := make([]string, len(args))
	copy(sanitized, args)
	
	return sanitized, nil
}