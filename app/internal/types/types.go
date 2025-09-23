package types

// EnvPreset represents an environment variable preset
type EnvPreset struct {
	Name        string `json:"name"`        // ボタン表示名
	Value       string `json:"value"`       // 環境変数値
	Description string `json:"description"` // 説明文
}

// Lesson represents a single tutorial lesson
type Lesson struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Code        string      `json:"code"`
	Filename    string      `json:"filename"`
	FilePath    string      `json:"file_path"` // Full path for version detection
	Stars       int         `json:"stars"`
	Version     string      `json:"version"`
	EnvPresets  []EnvPreset `json:"env_presets,omitempty"` // 環境変数プリセット
}

// Server represents the HTTP server with lesson data
type Server struct {
	Lessons map[string][]Lesson // バージョン -> レッスン
}
