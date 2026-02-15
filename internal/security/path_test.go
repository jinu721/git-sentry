package security

import (
	"testing"
)

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		input    string
		hasError bool
	}{
		{"normal/path", false},
		{"./relative", false},
		{"../traversal", true},
		{"", true},
		{"path/../traversal", true},
		{"clean/./path", false},
	}

	for _, test := range tests {
		result, err := SanitizePath(test.input)
		
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %s", test.input)
		}
		
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input %s: %v", test.input, err)
		}
		
		if !test.hasError && result == "" {
			t.Errorf("Expected non-empty result for valid input %s", test.input)
		}
	}
}

func TestValidateFilePath(t *testing.T) {
	validPaths := []string{
		"normal/file.txt",
		"./relative.go",
		"clean/path/file.md",
	}
	
	for _, path := range validPaths {
		if err := ValidateFilePath(path); err != nil {
			t.Errorf("Expected valid path %s to pass validation", path)
		}
	}
	
	invalidPaths := []string{
		"../traversal",
		"",
		"path/../../bad",
	}
	
	for _, path := range invalidPaths {
		if err := ValidateFilePath(path); err == nil {
			t.Errorf("Expected invalid path %s to fail validation", path)
		}
	}
}