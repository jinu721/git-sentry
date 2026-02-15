package security

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	SecureFileMode = 0644
	SecureDirMode  = 0755
)

func SecureWriteFile(path string, data []byte) error {
	cleanPath, err := SanitizePath(path)
	if err != nil {
		return fmt.Errorf("invalid file path: %w", err)
	}
	
	dir := filepath.Dir(cleanPath)
	if err := os.MkdirAll(dir, SecureDirMode); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	return os.WriteFile(cleanPath, data, SecureFileMode)
}

func SecureReadFile(path string) ([]byte, error) {
	cleanPath, err := SanitizePath(path)
	if err != nil {
		return nil, fmt.Errorf("invalid file path: %w", err)
	}
	
	return os.ReadFile(cleanPath)
}

func SecureCreateDir(path string) error {
	cleanPath, err := SanitizePath(path)
	if err != nil {
		return fmt.Errorf("invalid directory path: %w", err)
	}
	
	return os.MkdirAll(cleanPath, SecureDirMode)
}

func SecureFileExists(path string) (bool, error) {
	cleanPath, err := SanitizePath(path)
	if err != nil {
		return false, fmt.Errorf("invalid file path: %w", err)
	}
	
	_, err = os.Stat(cleanPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	
	return err == nil, err
}

func SecureRemoveFile(path string) error {
	cleanPath, err := SanitizePath(path)
	if err != nil {
		return fmt.Errorf("invalid file path: %w", err)
	}
	
	return os.Remove(cleanPath)
}