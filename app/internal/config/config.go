// Package config - Configuration management for Go Release Tour
//
// This package handles loading and managing version configurations
// from external config files, making it easy to add new Go versions
// without code changes.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// LessonInfo represents metadata about a single lesson
type LessonInfo struct {
	Title string `json:"title"`
	Stars int    `json:"stars"`
}

// VersionConfig represents configuration for a specific Go version
type VersionConfig struct {
	FullVersion string                `json:"full_version"`
	Path        string                `json:"path"`
	Lessons     map[string]LessonInfo `json:"lessons"`
}

// Config represents the complete configuration
type Config struct {
	Versions map[string]*VersionConfig `json:"versions"`
}

// ConfigManager handles configuration loading and management
type ConfigManager struct {
	config     *Config
	configPath string
}

// NewConfigManager creates a new configuration manager
func NewConfigManager(configPath string) *ConfigManager {
	if configPath == "" {
		configPath = "config/versions.json"
	}

	return &ConfigManager{
		configPath: configPath,
	}
}

// LoadConfig loads configuration from the specified file
func (cm *ConfigManager) LoadConfig() error {
	// 設定ファイルの絶対パスを取得
	absPath, err := filepath.Abs(cm.configPath)
	if err != nil {
		return fmt.Errorf("設定ファイルパス解決エラー: %w", err)
	}

	// ファイルの存在確認
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("設定ファイルが見つかりません: %s", absPath)
	}

	// ファイル読み込み
	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("設定ファイル読み込みエラー: %w", err)
	}

	// JSON解析
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("設定ファイル解析エラー: %w", err)
	}

	cm.config = &config

	return nil
}

// GetVersionConfig returns configuration for a specific version
func (cm *ConfigManager) GetVersionConfig(version string) (*VersionConfig, error) {
	if cm.config == nil {
		return nil, fmt.Errorf("設定が読み込まれていません")
	}

	config, exists := cm.config.Versions[version]
	if !exists {
		return nil, fmt.Errorf("バージョン %s の設定が見つかりません", version)
	}

	return config, nil
}

// GetAvailableVersions returns all available Go versions sorted in descending order
func (cm *ConfigManager) GetAvailableVersions() []string {
	if cm.config == nil {
		return nil
	}

	versions := make([]string, 0, len(cm.config.Versions))
	for version := range cm.config.Versions {
		versions = append(versions, version)
	}

	// バージョンを降順でソート（最新が先頭）
	sort.Slice(versions, func(i, j int) bool {
		return versions[i] > versions[j]
	})

	return versions
}

// GetDefaultVersion is removed - versions must be explicitly specified
// This ensures version execution guarantees

// GetLessonInfo returns lesson metadata for a specific version and filename
func (cm *ConfigManager) GetLessonInfo(version, filename string) (*LessonInfo, error) {
	versionConfig, err := cm.GetVersionConfig(version)
	if err != nil {
		return nil, err
	}

	lessonInfo, exists := versionConfig.Lessons[filename]
	if !exists {
		return nil, fmt.Errorf("バージョン %s にレッスン %s が見つかりません", version, filename)
	}

	return &lessonInfo, nil
}

// GetAllLessonsForVersion returns all lesson info for a specific version
func (cm *ConfigManager) GetAllLessonsForVersion(version string) (map[string]LessonInfo, error) {
	versionConfig, err := cm.GetVersionConfig(version)
	if err != nil {
		return nil, err
	}

	return versionConfig.Lessons, nil
}

// ValidateVersionPaths checks if the configured Go binaries exist
func (cm *ConfigManager) ValidateVersionPaths() map[string]error {
	if cm.config == nil {
		return map[string]error{"config": fmt.Errorf("設定が読み込まれていません")}
	}

	errors := make(map[string]error)

	for version, versionConfig := range cm.config.Versions {
		if _, err := os.Stat(versionConfig.Path); err != nil {
			errors[version] = fmt.Errorf("Goバイナリが見つかりません: %s", versionConfig.Path)
		}
	}

	return errors
}

// ReloadConfig reloads the configuration from file
func (cm *ConfigManager) ReloadConfig() error {
	return cm.LoadConfig()
}

// GetConfigSummary returns a summary of the loaded configuration
func (cm *ConfigManager) GetConfigSummary() map[string]interface{} {
	if cm.config == nil {
		return map[string]interface{}{
			"status": "未読み込み",
		}
	}

	summary := map[string]interface{}{
		"total_versions":     len(cm.config.Versions),
		"available_versions": cm.GetAvailableVersions(),
		"config_path":        cm.configPath,
	}

	// バージョン別詳細
	versionDetails := make(map[string]interface{})
	for version, versionConfig := range cm.config.Versions {
		versionDetails[version] = map[string]interface{}{
			"full_version": versionConfig.FullVersion,
			"path":         versionConfig.Path,
			"lesson_count": len(versionConfig.Lessons),
		}
	}
	summary["versions"] = versionDetails

	return summary
}
