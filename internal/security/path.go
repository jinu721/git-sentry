package security

import (
	"fmt"
	"path/filepath"
	"strings"
)

func SanitizePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("empty path not allowed")
	}
	
	cleaned := filepath.Clean(path)
	
	if strings.Contains(cleaned, "..") {
		return "", fmt.Errorf("directory traversal not allowed")
	}
	
	if filepath.IsAbs(cleaned) && !isAllowedAbsolutePath(cleaned) {
		return "", fmt.Errorf("absolute path not allowed: %s", cleaned)
	}
	
	return cleaned, nil
}

func isAllowedAbsolutePath(path string) bool {
	allowedPrefixes := []string{
		"/tmp/",
		"/var/tmp/",
	}
	
	for _, prefix := range allowedPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	
	return false
}

func ValidateFilePath(path string) error {
	_, err := SanitizePath(path)
	return err
}