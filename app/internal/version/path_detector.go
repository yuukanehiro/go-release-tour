// Package version - Path-based version detection
//
// This file implements version detection based on file paths,
// eliminating the need for metadata in lesson files.
package version

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// PathDetector handles version detection from file paths
type PathDetector struct {
	// Regex patterns for different path formats
	releasePathPattern *regexp.Regexp
	lessonPathPattern  *regexp.Regexp
}

// NewPathDetector creates a new path detector
func NewPathDetector() *PathDetector {
	return &PathDetector{
		// Pattern: releases/v/1.18/01_generics.go -> 1.18
		releasePathPattern: regexp.MustCompile(`releases/v/(\d+\.\d+)/`),
		// Pattern: /path/to/releases/v/1.18/01_generics.go -> 1.18
		lessonPathPattern: regexp.MustCompile(`(?:^|/)releases/v/(\d+\.\d+)/`),
	}
}

// ExtractVersionFromPath extracts Go version from file path
func (pd *PathDetector) ExtractVersionFromPath(filePath string) (string, error) {
	if filePath == "" {
		return "", fmt.Errorf("空のパスが指定されました")
	}

	// Normalize path separators
	normalizedPath := filepath.ToSlash(filePath)

	// Try release path pattern first
	matches := pd.releasePathPattern.FindStringSubmatch(normalizedPath)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	// Try lesson path pattern
	matches = pd.lessonPathPattern.FindStringSubmatch(normalizedPath)
	if len(matches) >= 2 {
		return matches[1], nil
	}

	return "", fmt.Errorf("パスからGoバージョンを特定できませんでした: %s", filePath)
}

// ExtractVersionFromLessonID extracts version from lesson identifier
// Format: "1.18/01_generics" -> "1.18"
func (pd *PathDetector) ExtractVersionFromLessonID(lessonID string) (string, error) {
	if lessonID == "" {
		return "", fmt.Errorf("空のレッスンIDが指定されました")
	}

	// Split by "/" and take the first part as version
	parts := strings.Split(lessonID, "/")
	if len(parts) >= 1 {
		version := parts[0]
		// Validate version format (e.g., "1.18")
		versionPattern := regexp.MustCompile(`^\d+\.\d+$`)
		if versionPattern.MatchString(version) {
			return version, nil
		}
	}

	return "", fmt.Errorf("レッスンIDからGoバージョンを特定できませんでした: %s", lessonID)
}

// IsValidVersionPath checks if a path contains a valid version directory
func (pd *PathDetector) IsValidVersionPath(filePath string) bool {
	_, err := pd.ExtractVersionFromPath(filePath)
	return err == nil
}

// GetVersionFromDirectory extracts version from directory name
// Input: "1.18" -> "1.18", "releases/v/1.18" -> "1.18"
func (pd *PathDetector) GetVersionFromDirectory(dirPath string) (string, error) {
	dirName := filepath.Base(dirPath)

	// Check if directory name is already a version
	versionPattern := regexp.MustCompile(`^(\d+\.\d+)$`)
	if matches := versionPattern.FindStringSubmatch(dirName); len(matches) >= 2 {
		return matches[1], nil
	}

	// Try to extract from full path
	return pd.ExtractVersionFromPath(dirPath)
}

// BuildLessonPath constructs a lesson file path from version and filename
func (pd *PathDetector) BuildLessonPath(version, filename string) string {
	return fmt.Sprintf("releases/v/%s/%s", version, filename)
}

// ValidateVersionFormat checks if a version string has valid format
func (pd *PathDetector) ValidateVersionFormat(version string) bool {
	versionPattern := regexp.MustCompile(`^\d+\.\d+$`)
	return versionPattern.MatchString(version)
}

// GetAllVersionsFromDirectory scans directory and returns all available versions
func (pd *PathDetector) GetAllVersionsFromDirectory(baseDir string) ([]string, error) {
	releasesDir := filepath.Join(baseDir, "releases", "v")

	entries, err := filepath.Glob(filepath.Join(releasesDir, "*"))
	if err != nil {
		return nil, fmt.Errorf("ディレクトリのスキャンに失敗しました: %v", err)
	}

	var versions []string
	for _, entry := range entries {
		if version, err := pd.GetVersionFromDirectory(entry); err == nil {
			versions = append(versions, version)
		}
	}

	return versions, nil
}

// LessonInfo represents lesson information extracted from path
type LessonInfo struct {
	Version    string
	Filename   string
	LessonName string
	FullPath   string
}

// ParseLessonPath parses a lesson file path and extracts information
func (pd *PathDetector) ParseLessonPath(filePath string) (*LessonInfo, error) {
	version, err := pd.ExtractVersionFromPath(filePath)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(filePath)
	lessonName := strings.TrimSuffix(filename, filepath.Ext(filename))

	return &LessonInfo{
		Version:    version,
		Filename:   filename,
		LessonName: lessonName,
		FullPath:   filePath,
	}, nil
}

// Global path detector instance
var globalPathDetector *PathDetector

// GetPathDetector returns the global path detector instance
func GetPathDetector() *PathDetector {
	if globalPathDetector == nil {
		globalPathDetector = NewPathDetector()
	}
	return globalPathDetector
}

// Convenience functions using global instance

// ExtractVersionFromPath extracts version from path using global detector
func ExtractVersionFromPath(filePath string) (string, error) {
	return GetPathDetector().ExtractVersionFromPath(filePath)
}

// ExtractVersionFromLessonID extracts version from lesson ID using global detector
func ExtractVersionFromLessonID(lessonID string) (string, error) {
	return GetPathDetector().ExtractVersionFromLessonID(lessonID)
}

// IsValidVersionPath checks if path contains valid version using global detector
func IsValidVersionPath(filePath string) bool {
	return GetPathDetector().IsValidVersionPath(filePath)
}

// BuildLessonPath builds lesson path using global detector
func BuildLessonPath(version, filename string) string {
	return GetPathDetector().BuildLessonPath(version, filename)
}
