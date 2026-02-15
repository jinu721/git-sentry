package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	
	if config.Rules.MaxFilesChanged <= 0 {
		t.Error("MaxFilesChanged should be positive")
	}
	
	if config.Rules.MaxLinesChanged <= 0 {
		t.Error("MaxLinesChanged should be positive")
	}
	
	if config.Rules.MaxMinutesSinceCommit <= 0 {
		t.Error("MaxMinutesSinceCommit should be positive")
	}
	
	if config.Rules.MaxUnpushedCommits <= 0 {
		t.Error("MaxUnpushedCommits should be positive")
	}
}

func TestGetConfigByTemplate(t *testing.T) {
	templates := []string{"team", "strict", "relaxed", "invalid"}
	
	for _, template := range templates {
		config := GetConfigByTemplate(template)
		if config == nil {
			t.Errorf("GetConfigByTemplate should never return nil for template: %s", template)
		}
		
		if config.Rules.MaxFilesChanged <= 0 {
			t.Errorf("Invalid MaxFilesChanged for template %s", template)
		}
	}
}

func TestConfigSaveLoad(t *testing.T) {
	tempDir := "test_config"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	originalConfig := TeamConfig()
	
	err := originalConfig.Save(tempDir)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}
	
	loadedConfig, err := Load(tempDir)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	
	if loadedConfig.Rules.MaxFilesChanged != originalConfig.Rules.MaxFilesChanged {
		t.Error("MaxFilesChanged mismatch after save/load")
	}
	
	if loadedConfig.AutoSuggestCommits != originalConfig.AutoSuggestCommits {
		t.Error("AutoSuggestCommits mismatch after save/load")
	}
}

func TestLoadWithTemplate(t *testing.T) {
	tempDir := "test_template"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	config, err := LoadWithTemplate(tempDir, "strict")
	if err != nil {
		t.Fatalf("Failed to load with template: %v", err)
	}
	
	strictConfig := StrictConfig()
	if config.Rules.MaxFilesChanged != strictConfig.Rules.MaxFilesChanged {
		t.Error("Template config not applied correctly")
	}
	
	configPath := filepath.Join(tempDir, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("Config file should be created")
	}
}