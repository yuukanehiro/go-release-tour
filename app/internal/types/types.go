package types

// Lesson represents a single tutorial lesson
type Lesson struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Filename    string `json:"filename"`
	Stars       int    `json:"stars"`
	Version     string `json:"version"`
}

// Server represents the HTTP server with lesson data
type Server struct {
	Lessons map[string][]Lesson // バージョン -> レッスン
}