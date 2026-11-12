package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ScriptConfig struct {
	Command      string   `json:"command"`
	Args         []string `json:"args"`
	Description  string   `json:"description"`
	RequiresArgs bool     `json:"requires_args"`
}

type Config struct {
	Scripts map[string]*ScriptConfig `json:"scripts"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(home, ".config", "repo", "config.json"), nil
}

func loadConfig() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Scripts: make(map[string]*ScriptConfig)}, nil
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", path, err)
	}

	if config.Scripts == nil {
		config.Scripts = make(map[string]*ScriptConfig)
	}

	return &config, nil
}

func findProjectConfig() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for {
		projPath := filepath.Join(dir, ".repo.json")
		if info, err := os.Stat(projPath); err == nil && !info.IsDir() {
			return projPath, nil
		}

		gitPath := filepath.Join(dir, ".git")
		if info, err := os.Stat(gitPath); err == nil && info.IsDir() {
			return "", nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", nil
		}
		dir = parent
	}
}

func loadProjectConfig() (*Config, error) {
	path, err := findProjectConfig()
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read project config %s: %w", path, err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse project config %s: %w", path, err)
	}

	if config.Scripts == nil {
		config.Scripts = make(map[string]*ScriptConfig)
	}

	return &config, nil
}

func loadMergedConfig() (*Config, error) {
	global, err := loadConfig()
	if err != nil {
		return nil, err
	}

	proj, err := loadProjectConfig()
	if err != nil {
		return nil, err
	}

	if proj == nil {
		return global, nil
	}

	merged := &Config{Scripts: make(map[string]*ScriptConfig)}
	for name, script := range global.Scripts {
		merged.Scripts[name] = script
	}
	for name, script := range proj.Scripts {
		merged.Scripts[name] = script
	}

	return merged, nil
}
