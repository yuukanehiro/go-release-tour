// Package version - Code execution with specific Go versions
//
// This file implements the core execution logic for running Go code
// with the appropriate Go version based on lesson requirements.
package version

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ExecutionRequest represents a code execution request
type ExecutionRequest struct {
	Code          string            `json:"code"`
	Version       string            `json:"version,omitempty"`        // 明示的なバージョン指定
	AutoDetect    bool              `json:"auto_detect,omitempty"`    // コードからバージョン自動検出
	Timeout       time.Duration     `json:"timeout,omitempty"`        // 実行タイムアウト
	Environment   map[string]string `json:"environment,omitempty"`    // 環境変数
	EnvVars       string            `json:"env_vars,omitempty"`       // 環境変数文字列（例: "GOEXPERIMENT=jsonv2"）
	WorkingDir    string            `json:"working_dir,omitempty"`    // 作業ディレクトリ
	StrictVersion bool              `json:"strict_version,omitempty"` // 厳密なバージョンチェック
}

// ExecutionResult represents the result of code execution
type ExecutionResult struct {
	Output          string        `json:"output"`
	Error           string        `json:"error,omitempty"`
	ExitCode        int           `json:"exit_code"`
	ExecutionTime   time.Duration `json:"execution_time"`
	GoVersion       string        `json:"go_version"`
	UsedVersion     string        `json:"used_version"`               // 実際に使用されたバージョン
	DetectedVersion string        `json:"detected_version,omitempty"` // 検出されたバージョン
	VersionPath     string        `json:"version_path,omitempty"`     // 使用されたGoバイナリのパス
}

// Executor handles Go code execution with version management
type Executor struct {
	manager *Manager
}

// NewExecutor creates a new code executor
func NewExecutor() *Executor {
	return &Executor{
		manager: GetManager(),
	}
}

// Execute runs Go code with the appropriate version
func (e *Executor) Execute(req ExecutionRequest) (*ExecutionResult, error) {
	startTime := time.Now()

	// デフォルト値の設定
	if req.Timeout == 0 {
		req.Timeout = 30 * time.Second
	}

	result := &ExecutionResult{}

	// バージョンの決定
	targetVersion, err := e.determineVersion(req)
	if err != nil {
		result.Error = fmt.Sprintf("バージョン決定エラー: %v", err)
		return result, err
	}

	result.UsedVersion = targetVersion

	// 検出されたバージョンを記録
	if req.AutoDetect {
		if detectedVersion, err := e.manager.ExtractVersionFromCode(req.Code); err == nil {
			result.DetectedVersion = detectedVersion
		}
	}

	// バージョン設定の取得
	versionConfig, err := e.manager.GetVersionConfig(targetVersion)
	if err != nil {
		result.Error = fmt.Sprintf("バージョン設定エラー: %v", err)
		return result, err
	}

	result.VersionPath = versionConfig.Path
	result.GoVersion = versionConfig.FullVersion

	// 厳密なバージョンチェック
	if req.StrictVersion && req.Version != "" && req.Version != targetVersion {
		err := fmt.Errorf("厳密モード: 要求バージョン %s と決定バージョン %s が一致しません", req.Version, targetVersion)
		result.Error = err.Error()
		return result, err
	}

	// コードの実行
	output, exitCode, err := e.executeCode(req.Code, versionConfig, req)

	result.Output = output
	result.ExitCode = exitCode
	result.ExecutionTime = time.Since(startTime)

	if err != nil {
		result.Error = err.Error()
	}

	return result, nil
}

// determineVersion determines which Go version to use for execution
func (e *Executor) determineVersion(req ExecutionRequest) (string, error) {
	log.Printf("[DEBUG] determineVersion: Starting version determination")
	log.Printf("[DEBUG] determineVersion: Request params - Version=%q, WorkingDir=%q, AutoDetect=%t", req.Version, req.WorkingDir, req.AutoDetect)
	log.Printf("[DEBUG] determineVersion: Code length=%d characters", len(req.Code))

	// 1. 明示的なバージョン指定がある場合
	if req.Version != "" {
		log.Printf("[DEBUG] determineVersion: Using explicit version: %s", req.Version)
		return req.Version, nil
	}

	// 2. ワーキングディレクトリからパス判定（優先）
	if req.WorkingDir != "" {
		log.Printf("[DEBUG] determineVersion: Attempting path-based detection from WorkingDir: %s", req.WorkingDir)
		if version, err := ExtractVersionFromPath(req.WorkingDir); err == nil {
			log.Printf("[DEBUG] determineVersion: Successfully extracted version from WorkingDir: %s", version)
			return version, nil
		} else {
			log.Printf("[DEBUG] determineVersion: Failed to extract version from WorkingDir: %v", err)
		}
	} else {
		log.Printf("[DEBUG] determineVersion: No WorkingDir provided")
	}

	// 3. 自動検出が有効な場合（コードとパス両方）
	if req.AutoDetect {
		log.Printf("[DEBUG] determineVersion: AutoDetect enabled, attempting code-based detection")
		// 3a. コードからパス情報を抽出
		if version, err := e.manager.ExtractVersionFromCode(req.Code); err == nil {
			log.Printf("[DEBUG] determineVersion: Successfully extracted version from code: %s", version)
			return version, nil
		} else {
			log.Printf("[DEBUG] determineVersion: Failed to extract version from code: %v", err)
			log.Printf("[DEBUG] determineVersion: Code sample (first 200 chars): %q", truncateString(req.Code, 200))
		}
	} else {
		log.Printf("[DEBUG] determineVersion: AutoDetect disabled")
	}

	// 4. バージョンが特定できない場合はエラー
	log.Printf("[DEBUG] determineVersion: All detection methods failed")
	return "", fmt.Errorf("バージョンを特定できませんでした。明示的なバージョン指定またはレッスンパスが必要です")
}

// Helper function to truncate strings for debugging
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// executeCode executes the Go code with the specified version
func (e *Executor) executeCode(code string, config *VersionConfig, req ExecutionRequest) (string, int, error) {
	// 一時ファイルの作成（常にシステム一時ディレクトリを使用）
	tempDir := os.TempDir()
	filename := filepath.Join(tempDir, fmt.Sprintf("gocode_%d_%d.go", time.Now().Unix(), time.Now().Nanosecond()))

	// コードをファイルに書き込み
	if err := os.WriteFile(filename, []byte(code), 0644); err != nil {
		return "", 1, fmt.Errorf("コードファイル作成エラー: %v", err)
	}

	// 実行後にファイルを削除
	defer func() {
		if err := os.Remove(filename); err != nil {
			// ログに記録するが、エラーは無視
			fmt.Printf("Warning: failed to remove temp file %s: %v\n", filename, err)
		}
	}()

	// Go実行コマンドの作成
	cmd := exec.Command(config.Path, "run", filename)

	// 環境変数の設定
	cmd.Env = os.Environ()
	if req.Environment != nil {
		for key, value := range req.Environment {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
	}

	// EnvVars文字列の処理（例: "GOEXPERIMENT=jsonv2"）
	if req.EnvVars != "" {
		envPairs := strings.Split(req.EnvVars, ",")
		for _, pair := range envPairs {
			if pair = strings.TrimSpace(pair); pair != "" {
				cmd.Env = append(cmd.Env, pair)
				log.Printf("[DEBUG] Execute: Added environment variable: %s", pair)
			}
		}
	}

	// 作業ディレクトリの設定
	// WorkingDirはバージョン検出のみに使用し、実行時の作業ディレクトリは設定しない
	// これにより、コードは一時ディレクトリで実行される

	// タイムアウト付きでコマンド実行
	done := make(chan struct{})
	var output []byte
	var err error
	var exitCode int

	go func() {
		defer close(done)
		output, err = cmd.CombinedOutput()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode = exitError.ExitCode()
			} else {
				exitCode = 1
			}
		}
	}()

	// タイムアウト処理
	select {
	case <-done:
		// 正常終了
		return string(output), exitCode, err
	case <-time.After(req.Timeout):
		// タイムアウト
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return "", 124, fmt.Errorf("実行タイムアウト (%v)", req.Timeout)
	}
}

// ValidateCode performs basic validation on the Go code before execution
func (e *Executor) ValidateCode(code string, version string) error {
	// 基本的なGoコードの検証
	if code == "" {
		return fmt.Errorf("空のコードは実行できません")
	}

	// 危険なコードパターンの検出
	dangerousPatterns := []string{
		"os.RemoveAll",
		"os.Remove",
		"exec.Command",
		"syscall",
		"unsafe",
		"//go:linkname",
	}

	for _, pattern := range dangerousPatterns {
		if contains(code, pattern) {
			return fmt.Errorf("セキュリティ上の理由により、'%s' を含むコードは実行できません", pattern)
		}
	}

	// バージョン固有の検証
	if version != "" {
		if err := e.validateVersionSpecificFeatures(code, version); err != nil {
			return err
		}
	}

	return nil
}

// validateVersionSpecificFeatures validates version-specific Go features
func (e *Executor) validateVersionSpecificFeatures(code string, version string) error {
	// Go 1.18未満でのジェネリクス使用チェック
	versionFloat := e.manager.parseVersion(version)

	if versionFloat < 1.18 {
		genericPatterns := []string{
			"[T any]",
			"[T comparable]",
			"[T constraint",
			"type.*\\[.*\\]",
		}

		for _, pattern := range genericPatterns {
			if matched, _ := regexp.MatchString(pattern, code); matched {
				return fmt.Errorf("ジェネリクスはGo 1.18以降で利用可能です（現在: %s）", version)
			}
		}
	}

	return nil
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// ExecuteWithVersion is a convenience method for executing code with a specific version
func (e *Executor) ExecuteWithVersion(code, version string) (*ExecutionResult, error) {
	return e.Execute(ExecutionRequest{
		Code:          code,
		Version:       version,
		AutoDetect:    false,
		Timeout:       30 * time.Second,
		StrictVersion: true,
	})
}

// ExecuteWithAutoDetect is a convenience method for executing code with auto-detected version
func (e *Executor) ExecuteWithAutoDetect(code string) (*ExecutionResult, error) {
	return e.Execute(ExecutionRequest{
		Code:       code,
		AutoDetect: true,
		Timeout:    30 * time.Second,
	})
}

// GetSupportedVersions returns all supported Go versions
func (e *Executor) GetSupportedVersions() []string {
	return e.manager.GetAvailableVersions()
}

// GetVersionInfo returns detailed information about all versions
func (e *Executor) GetVersionInfo() map[string]*VersionConfig {
	return e.manager.GetAllVersionConfigs()
}
