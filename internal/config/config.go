package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"gitsentry/internal/security"
)

type Config struct {
	Rules               Rules `yaml:"rules"`
	AutoSuggestCommits  bool  `yaml:"auto_suggest_commits"`
	AutoSuggestPushes   bool  `yaml:"auto_suggest_pushes"`
	CommitMessageFormat string `yaml:"commit_message_format"`
}

type Rules struct {
	MaxFilesChanged        int `yaml:"max_files_changed"`
	MaxLinesChanged        int `yaml:"max_lines_changed"`
	MaxMinutesSinceCommit  int `yaml:"max_minutes_since_commit"`
	MaxUnpushedCommits     int `yaml:"max_unpushed_commits"`
}

func DefaultConfig() *Config {
	return &Config{
		Rules: Rules{
			MaxFilesChanged:       5,
			MaxLinesChanged:       100,
			MaxMinutesSinceCommit: 30,
			MaxUnpushedCommits:    3,
		},
		AutoSuggestCommits:  true,
		AutoSuggestPushes:   true,
		CommitMessageFormat: "conventional", // conventional, simple
	}
}

func Load(gitsentryDir string) (*Config, error) {
	configPath := filepath.Join(gitsentryDir, "config.yaml")
	
	// If config doesn't exist, create default
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := DefaultConfig()
		if err := config.Save(gitsentryDir); err != nil {
			return nil, err
		}
		return config, nil
	}
	
	// Load existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	
	if err := security.ValidateConfigStruct(&config); err != nil {
		return nil, err
	}
	
	return &config, nil
}

func (c *Config) Save(gitsentryDir string) error {
	configPath := filepath.Join(gitsentryDir, "config.yaml")
	
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	
	return os.WriteFile(configPath, data, 0644)
}