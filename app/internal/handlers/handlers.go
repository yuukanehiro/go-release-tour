package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-release-tour/app/internal/types"
	"go-release-tour/app/internal/version"
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
			http.Error(w, "Version parameter is required", http.StatusBadRequest)
			return
		}
		if lessons, exists := s.Lessons[version]; exists {
			json.NewEncoder(w).Encode(lessons)
		} else {
			http.Error(w, "Version not found", http.StatusNotFound)
		}
	}
}

// CodeRunRequest represents a code execution request with version support
type CodeRunRequest struct {
	Code    string `json:"code"`
	Version string `json:"version"`    // 実行するGoバージョン（フロントエンドで決定済み）
}

// CodeRunResponse represents a code execution response with version info
type CodeRunResponse struct {
	Output          string `json:"output"`
	Error           string `json:"error,omitempty"`
	GoVersion       string `json:"go_version,omitempty"`       // 使用されたGoの完全バージョン
	UsedVersion     string `json:"used_version,omitempty"`     // 使用されたGoバージョン（例: 1.18）
	DetectedVersion string `json:"detected_version,omitempty"` // 検出されたバージョン
	ExecutionTime   string `json:"execution_time,omitempty"`   // 実行時間
	VersionPath     string `json:"version_path,omitempty"`     // 使用されたGoバイナリのパス
}

// HandleRun executes Go code with appropriate version and returns the result
func HandleRun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CodeRunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[DEBUG] HandleRun: Failed to decode JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("[DEBUG] HandleRun: Received request - Version=%q", req.Version)
	log.Printf("[DEBUG] HandleRun: Code length=%d characters", len(req.Code))

	if req.Version == "" {
		log.Printf("[DEBUG] HandleRun: No version specified")
		response := CodeRunResponse{
			Error: "バージョンが指定されていません",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// バージョン対応の実行器を作成
	executor := version.NewExecutor()

	// 実行リクエストを構築
	execReq := version.ExecutionRequest{
		Code:       req.Code,
		Version:    req.Version,
		AutoDetect: false, // フロントエンドで決定済みなので自動検出不要
		Timeout:    30 * time.Second,
	}

	log.Printf("[DEBUG] HandleRun: Using version: %s", req.Version)

	// コード検証
	log.Printf("[DEBUG] HandleRun: Validating code for version %s", req.Version)
	if err := executor.ValidateCode(req.Code, req.Version); err != nil {
		log.Printf("[DEBUG] HandleRun: Code validation failed: %v", err)
		response := CodeRunResponse{
			Error: fmt.Sprintf("コード検証エラー: %v", err),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	log.Printf("[DEBUG] HandleRun: Code validation passed")

	// コードを実行
	log.Printf("[DEBUG] HandleRun: Starting code execution")
	result, err := executor.Execute(execReq)
	log.Printf("[DEBUG] HandleRun: Execution completed - err=%v, result.Error=%q", err, result.Error)
	log.Printf("[DEBUG] HandleRun: Execution result - GoVersion=%q, UsedVersion=%q", result.GoVersion, result.UsedVersion)

	// レスポンスを構築
	response := CodeRunResponse{
		Output:          result.Output,
		GoVersion:       result.GoVersion,
		UsedVersion:     result.UsedVersion,
		DetectedVersion: req.Version, // フロントエンドで決定されたバージョンをそのまま返す
		ExecutionTime:   result.ExecutionTime.String(),
		VersionPath:     result.VersionPath,
	}

	if err != nil || result.Error != "" {
		errorMsg := ""
		if err != nil {
			errorMsg = err.Error()
		}
		if result.Error != "" {
			if errorMsg != "" {
				errorMsg += "; " + result.Error
			} else {
				errorMsg = result.Error
			}
		}
		response.Error = errorMsg
		log.Printf("[DEBUG] HandleRun: Response will include error: %s", errorMsg)
	} else {
		log.Printf("[DEBUG] HandleRun: Execution successful, no errors")
	}

	log.Printf("[DEBUG] HandleRun: Sending response")
	json.NewEncoder(w).Encode(response)
}

// HandleVersionInfo returns detailed version information
func HandleVersionInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	manager := version.GetManager()
	versionInfo := manager.Status()

	json.NewEncoder(w).Encode(versionInfo)
}