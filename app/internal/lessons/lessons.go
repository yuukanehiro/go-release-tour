package lessons

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-release-tour/app/internal/config"
	"go-release-tour/app/internal/types"
)

// LoadLessons loads all lessons using configuration
func LoadLessons(s *types.Server) {
	// 設定マネージャーを初期化
	configManager := config.NewConfigManager("")
	if err := configManager.LoadConfig(); err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	// GO_VERSION環境変数に基づいてレッスンを読み込み
	goVersion := os.Getenv("GO_VERSION")
	if goVersion != "" {
		loadVersionLessons(s, goVersion, configManager)
	} else {
		// 設定ファイルから全バージョンを読み込み
		for _, version := range configManager.GetAvailableVersions() {
			loadVersionLessons(s, version, configManager)
		}
	}
}

// loadVersionLessons loads lessons for a specific Go version using config
func loadVersionLessons(s *types.Server, version string, configManager *config.ConfigManager) {
	lessonsDir := fmt.Sprintf("releases/v/%s", version)
	files, err := filepath.Glob(filepath.Join(lessonsDir, "*.go"))
	if err != nil {
		log.Printf("Error loading lessons for version %s: %v", version, err)
		return
	}

	// 設定ファイルからレッスンデータを取得
	lessonData, err := configManager.GetAllLessonsForVersion(version)
	if err != nil {
		log.Printf("Error getting lesson data for version %s: %v", version, err)
		return
	}

	var lessons []types.Lesson
	for i, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}

		filename := filepath.Base(file)
		data, exists := lessonData[filename]
		if !exists {
			log.Printf("Warning: Lesson metadata not found for %s in version %s", filename, version)
			continue
		}

		// コメントから説明を抽出
		lines := strings.Split(string(content), "\n")
		var description string
		for _, line := range lines {
			if strings.HasPrefix(line, "// 説明:") {
				description = strings.TrimPrefix(line, "// 説明: ")
				break
			}
		}

		lesson := types.Lesson{
			ID:          i + 1,
			Title:       data.Title,
			Description: description,
			Code:        string(content),
			Filename:    filename,
			FilePath:    file,               // ファイルパスを追加
			Stars:       data.Stars,
			Version:     version,
		}
		lessons = append(lessons, lesson)
	}
	s.Lessons[version] = lessons
}

// Note: Lesson metadata is now loaded from config/versions.json
// This provides a flexible way to add new versions without code changes