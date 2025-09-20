package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"go-release-tour/app/internal/types"
)

// HandleVersions returns available Go versions
func HandleVersions(s *types.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		versions := make([]string, 0, len(s.Lessons))
		for version := range s.Lessons {
			versions = append(versions, version)
		}
		json.NewEncoder(w).Encode(versions)
	}
}

// HandleLessons returns lessons for a specific version
func HandleLessons(s *types.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		version := r.URL.Query().Get("version")
		if version == "" {
			version = "1.25" // デフォルトバージョン
		}
		if lessons, exists := s.Lessons[version]; exists {
			json.NewEncoder(w).Encode(lessons)
		} else {
			http.Error(w, "Version not found", http.StatusNotFound)
		}
	}
}

// CodeRunRequest represents a code execution request
type CodeRunRequest struct {
	Code string `json:"code"`
}

// CodeRunResponse represents a code execution response
type CodeRunResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

// HandleRun executes Go code and returns the result
func HandleRun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CodeRunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 一時ファイルを作成
	tempDir := os.TempDir()
	filename := filepath.Join(tempDir, fmt.Sprintf("gocode_%d_%d.go", time.Now().Unix(), time.Now().Nanosecond()))

	// ファイルに書き込み
	if err := os.WriteFile(filename, []byte(req.Code), 0644); err != nil {
		response := CodeRunResponse{
			Error: fmt.Sprintf("Failed to write code to file: %v", err),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 実行後にファイルを削除
	defer os.Remove(filename)

	// Goコードを実行
	cmd := exec.Command("go", "run", filename)
	output, err := cmd.CombinedOutput()

	response := CodeRunResponse{
		Output: string(output),
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}