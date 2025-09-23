// Package version provides Go version management for multi-version code execution
//
// This package manages multiple Go versions installed in the system and provides
// functionality to execute code with the appropriate Go version based on lesson requirements.
//
// Architecture:
// - Version detection from lesson metadata
// - Path resolution for different Go versions
// - Execution environment validation
// - Fallback handling for unsupported versions
package version

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// VersionConfig represents a Go version configuration
type VersionConfig struct {
	Version     string `json:"version"`      // e.g., "1.18"
	Path        string `json:"path"`         // e.g., "/opt/go1.18/bin/go"
	FullVersion string `json:"full_version"` // e.g., "1.18.10"
	Available   bool   `json:"available"`    // Whether this version is actually available
}

// Manager handles Go version management
type Manager struct {
	versions map[string]*VersionConfig
	mutex    sync.RWMutex
}

// Global manager instance
var globalManager *Manager
var once sync.Once

// GetManager returns the singleton version manager
func GetManager() *Manager {
	once.Do(func() {
		globalManager = NewManager()
		globalManager.Initialize()
	})
	return globalManager
}

// NewManager creates a new version manager
func NewManager() *Manager {
	return &Manager{
		versions: make(map[string]*VersionConfig),
	}
}

// Initialize sets up the version manager with predefined Go versions
func (m *Manager) Initialize() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 対応するGoバージョンの定義
	supportedVersions := map[string]string{
		"1.18": "/opt/go1.18/bin/go",
		"1.19": "/opt/go1.19/bin/go",
		"1.20": "/opt/go1.20/bin/go",
		"1.21": "/opt/go1.21/bin/go",
		"1.22": "/opt/go1.22/bin/go",
		"1.23": "/opt/go1.23/bin/go",
		"1.24": "/opt/go1.24/bin/go",
		"1.25": "/opt/go1.25/bin/go",
	}

	// 各バージョンを初期化し、利用可能性をチェック
	for version, path := range supportedVersions {
		config := &VersionConfig{
			Version:   version,
			Path:      path,
			Available: m.checkVersionAvailability(path),
		}

		// 実際のバージョン情報を取得
		if config.Available {
			if fullVersion, err := m.getFullVersion(path); err == nil {
				config.FullVersion = fullVersion
			}
		}

		m.versions[version] = config
	}
}

// checkVersionAvailability checks if a Go version is available at the given path
func (m *Manager) checkVersionAvailability(goPath string) bool {
	if _, err := os.Stat(goPath); os.IsNotExist(err) {
		return false
	}

	// 実際にバージョン確認を実行
	cmd := exec.Command(goPath, "version")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

// getFullVersion retrieves the full version string from a Go binary
func (m *Manager) getFullVersion(goPath string) (string, error) {
	cmd := exec.Command(goPath, "version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// "go version go1.18.10 linux/amd64" から "1.18.10" を抽出
	versionPattern := regexp.MustCompile(`go(\d+\.\d+\.\d+)`)
	matches := versionPattern.FindStringSubmatch(string(output))
	if len(matches) >= 2 {
		return matches[1], nil
	}

	return "", fmt.Errorf("バージョン情報を解析できませんでした: %s", string(output))
}

// GetVersionConfig returns the configuration for a specific Go version
func (m *Manager) GetVersionConfig(version string) (*VersionConfig, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	config, exists := m.versions[version]
	if !exists {
		return nil, fmt.Errorf("サポートされていないGoバージョン: %s", version)
	}

	if !config.Available {
		return nil, fmt.Errorf("Goバージョン %s はインストールされていません", version)
	}

	return config, nil
}

// GetAvailableVersions returns all available Go versions
func (m *Manager) GetAvailableVersions() []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var available []string
	for version, config := range m.versions {
		if config.Available {
			available = append(available, version)
		}
	}

	return available
}

// GetAllVersionConfigs returns all version configurations
func (m *Manager) GetAllVersionConfigs() map[string]*VersionConfig {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	// コピーを作成して返す
	result := make(map[string]*VersionConfig)
	for version, config := range m.versions {
		result[version] = &VersionConfig{
			Version:     config.Version,
			Path:        config.Path,
			FullVersion: config.FullVersion,
			Available:   config.Available,
		}
	}

	return result
}

// ExtractVersionFromCode extracts the required Go version from lesson code or path
// This is kept for compatibility but now primarily uses path-based detection
func (m *Manager) ExtractVersionFromCode(code string) (string, error) {
	// 1. Try to extract from file path patterns in comments
	// Example: releases/v/1.18/01_generics.go
	pathPattern := regexp.MustCompile(`releases/v/(\d+\.\d+)/`)
	matches := pathPattern.FindStringSubmatch(code)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	// 2. Extract from "Go X.Y 新機能" or "Go X.Y" comments
	// Example: // Go 1.24 新機能: Generic Type Aliases
	// Example: // Go 1.18 generics
	goVersionPattern := regexp.MustCompile(`//.*Go\s+(\d+\.\d+)[\s新機能:]`)
	matches = goVersionPattern.FindStringSubmatch(code)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	// 3. Simpler pattern for "Go X.Y" in comments
	// Example: // Go 1.24 新機能
	simpleGoPattern := regexp.MustCompile(`//.*Go\s+(\d+\.\d+)`)
	matches = simpleGoPattern.FindStringSubmatch(code)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	// 4. Legacy: Extract from GO_VERSION metadata (if exists)
	// Example: // GO_VERSION: 1.18
	versionPattern := regexp.MustCompile(`//\s*GO_VERSION:\s*(\d+\.\d+)`)
	matches = versionPattern.FindStringSubmatch(code)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	return "", fmt.Errorf("コードからGoバージョンを特定できませんでした")
}

// ExtractVersionFromPath extracts version from lesson file path using path detector
func (m *Manager) ExtractVersionFromPath(filePath string) (string, error) {
	return ExtractVersionFromPath(filePath)
}

// ValidateVersionSupport validates if a version supports specific features
func (m *Manager) ValidateVersionSupport(version string, features []string) error {
	config, err := m.GetVersionConfig(version)
	if err != nil {
		return err
	}

	// 基本的なバージョン互換性チェック
	versionFloat := m.parseVersion(config.Version)

	for _, feature := range features {
		if !m.isFeatureSupported(versionFloat, feature) {
			return fmt.Errorf("機能 '%s' はGo %s でサポートされていません", feature, version)
		}
	}

	return nil
}

// parseVersion converts version string to float for comparison
func (m *Manager) parseVersion(version string) float64 {
	// "1.18" -> 1.18
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		major := parts[0]
		minor := parts[1]
		versionStr := major + "." + minor
		var versionFloat float64
		if f, err := fmt.Sscanf(versionStr, "%f", &versionFloat); err == nil && f == 1 {
			return versionFloat
		}
	}
	return 0.0
}

// isFeatureSupported checks if a feature is supported in the given version
func (m *Manager) isFeatureSupported(version float64, feature string) bool {
	// 機能別の最小バージョン要件
	featureRequirements := map[string]float64{
		"generics":             1.18,
		"workspace":            1.18,
		"type-parameters":      1.18,
		"atomic-types":         1.19,
		"memory-arenas":        1.19,
		"comparable-types":     1.20,
		"slice-to-array":       1.20,
		"errors-join":          1.20,
		"builtin-functions":    1.21,
		"slices-package":       1.21,
		"maps-package":         1.21,
		"for-range-int":        1.22,
		"enhanced-routing":     1.22,
		"loop-variables":       1.22,
		"structured-logging":   1.23,
		"iterators":            1.23,
		"generic-aliases":      1.24,
		"swiss-tables":         1.24,
		"weak-pointers":        1.24,
		"container-gomaxprocs": 1.25,
		"synctest":             1.25,
		"json-v2":              1.25,
	}

	requiredVersion, exists := featureRequirements[feature]
	if !exists {
		return true // 不明な機能はサポートされているとみなす
	}

	return version >= requiredVersion
}

// GetDefaultVersion is removed - versions must be explicitly specified
// This ensures version execution guarantees and prevents accidental version usage
func (m *Manager) GetDefaultVersion() string {
	// デフォルトバージョンは廃止 - 明示的な指定を強制
	panic("GetDefaultVersion is deprecated - explicit version specification required for execution guarantees")
}

// Status returns the current status of the version manager
func (m *Manager) Status() map[string]interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	availableCount := 0
	for _, config := range m.versions {
		if config.Available {
			availableCount++
		}
	}

	return map[string]interface{}{
		"total_versions":            len(m.versions),
		"available_versions":        availableCount,
		"multi_version_support":     true,
		"explicit_version_required": true,
		"versions":                  m.GetAllVersionConfigs(),
	}
}
