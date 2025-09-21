package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TestRunner API経由統合テストのランナー
type TestRunner struct {
	APIURL    string
	OutputDir string
	Verbose   bool
	Client    *http.Client
}

// NewTestRunner テストランナーを作成
func NewTestRunner(apiURL, outputDir string, verbose bool) *TestRunner {
	return &TestRunner{
		APIURL:    apiURL,
		OutputDir: outputDir,
		Verbose:   verbose,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// RunAllTests 全テストを実行
func (r *TestRunner) RunAllTests() (*TestResults, error) {
	fmt.Printf("開始時刻: %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("API URL: %s\n\n", r.APIURL)

	results := NewTestResults()

	// 対象バージョンを動的に取得
	versions, err := r.getAvailableVersions()
	if err != nil {
		return nil, fmt.Errorf("failed to get available versions: %w", err)
	}

	// 各バージョンのテスト
	for _, version := range versions {
		fmt.Printf("=== Go %s のテスト ===\n", version)

		versionResults, err := r.runVersionTests(version)
		if err != nil {
			return nil, fmt.Errorf("failed to run tests for version %s: %w", version, err)
		}

		results.AddVersionResults(version, versionResults)
		fmt.Println()
	}

	// 基本APIテスト
	fmt.Println("=== 基本APIテスト ===")
	basicResults, err := r.runBasicAPITests(versions)
	if err != nil {
		return nil, fmt.Errorf("failed to run basic API tests: %w", err)
	}
	results.AddBasicResults(basicResults)

	// サマリー表示
	fmt.Printf("\n=== テスト結果サマリー ===\n")
	fmt.Printf("総テスト数: %d\n", results.TotalTests)
	fmt.Printf("成功: %d\n", results.SuccessTests)
	fmt.Printf("エラー: %d\n", results.ErrorTests)

	if results.TotalTests > 0 {
		successRate := float64(results.SuccessTests) * 100.0 / float64(results.TotalTests)
		fmt.Printf("成功率: %.1f%%\n", successRate)
	}

	fmt.Printf("\n終了時刻: %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))

	return results, nil
}

// runVersionTests 特定バージョンのテストを実行
func (r *TestRunner) runVersionTests(version string) ([]*TestResult, error) {
	versionDir := fmt.Sprintf("../../releases/v/%s", version)

	// .goファイルを検索
	goFiles, err := r.findGoFiles(versionDir)
	if err != nil {
		fmt.Printf("[WARN] ディレクトリが存在しないか読み取れません: %s\n", versionDir)
		return nil, nil
	}

	if len(goFiles) == 0 {
		fmt.Printf("[WARN] .goファイルが見つかりません: %s\n", versionDir)
		return nil, nil
	}

	fmt.Printf("テストファイル数: %8d\n", len(goFiles))

	var results []*TestResult
	for _, goFile := range goFiles {
		result, err := r.runSingleTest(version, goFile)
		if err != nil {
			return nil, fmt.Errorf("failed to run test %s: %w", goFile, err)
		}
		results = append(results, result)
	}

	return results, nil
}

// runBasicAPITests 基本APIテストを実行
func (r *TestRunner) runBasicAPITests(versions []string) ([]*TestResult, error) {
	helloWorldCode := `package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}`

	var results []*TestResult
	for _, version := range versions {
		result := r.testAPIWithCode(version, helloWorldCode, "Basic Hello World")
		results = append(results, result)
	}

	return results, nil
}

// runSingleTest 単一ファイルのテストを実行
func (r *TestRunner) runSingleTest(version, filePath string) (*TestResult, error) {
	filename := filepath.Base(filePath)

	// ファイル内容を読み取り
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("  [SKIP] ファイルが読み取れません: %s\n", filename)
		return &TestResult{
			TestName: filename,
			Version:  version,
			Status:   "SKIP",
			Error:    fmt.Sprintf("Failed to read file: %v", err),
		}, nil
	}

	return r.testAPIWithCode(version, string(content), filename), nil
}

// testAPIWithCode APIでコードをテスト
func (r *TestRunner) testAPIWithCode(version, code, testName string) *TestResult {
	fmt.Printf("  API テスト: %s (Go %s) ... ", testName, version)

	// APIリクエストのペイロード作成
	payload := map[string]string{
		"code":    code,
		"version": version,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("[FAIL]")
		return &TestResult{
			TestName: testName,
			Version:  version,
			Status:   "FAIL",
			Error:    fmt.Sprintf("Failed to marshal payload: %v", err),
		}
	}

	// API呼び出し
	resp, err := r.Client.Post(r.APIURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("[FAIL]")
		return &TestResult{
			TestName: testName,
			Version:  version,
			Status:   "FAIL",
			Error:    fmt.Sprintf("HTTP request failed: %v", err),
		}
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[FAIL]")
		return &TestResult{
			TestName: testName,
			Version:  version,
			Status:   "FAIL",
			Error:    fmt.Sprintf("Failed to read response: %v", err),
		}
	}

	// レスポンスの解析
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		fmt.Println("[FAIL]")
		return &TestResult{
			TestName: testName,
			Version:  version,
			Status:   "FAIL",
			Error:    fmt.Sprintf("Failed to parse response: %v", err),
			RawResponse: string(responseBody),
		}
	}

	// エラーチェック
	if errorField, exists := apiResponse["error"]; exists && errorField != nil && errorField != "" {
		fmt.Println("[FAIL]")
		return &TestResult{
			TestName: testName,
			Version:  version,
			Status:   "FAIL",
			Error:    fmt.Sprintf("%v", errorField),
			RawResponse: string(responseBody),
		}
	}

	// 成功
	fmt.Println("[PASS]")

	result := &TestResult{
		TestName: testName,
		Version:  version,
		Status:   "PASS",
		RawResponse: string(responseBody),
	}

	// 追加情報を表示
	if r.Verbose {
		if usedVersion, ok := apiResponse["used_version"]; ok {
			fmt.Printf("    実行バージョン: %v\n", usedVersion)
		}
		if output, ok := apiResponse["output"]; ok && output != nil {
			outputStr := fmt.Sprintf("%v", output)
			if len(outputStr) > 100 {
				outputStr = outputStr[:100] + "..."
			}
			fmt.Printf("    出力: %s\n", outputStr)
		}
	}

	return result
}

// getAvailableVersions 利用可能なGoバージョンを取得
func (r *TestRunner) getAvailableVersions() ([]string, error) {
	releasesDir := "../../releases/v"

	entries, err := os.ReadDir(releasesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read releases directory: %w", err)
	}

	var versions []string
	for _, entry := range entries {
		if entry.IsDir() {
			// ディレクトリ名がバージョン形式かチェック
			if r.isValidVersionFormat(entry.Name()) {
				versions = append(versions, entry.Name())
			}
		}
	}

	// バージョン順にソート（降順：新しいバージョンから）
	sort.Slice(versions, func(i, j int) bool {
		return r.compareVersions(versions[i], versions[j]) > 0
	})

	return versions, nil
}

// isValidVersionFormat バージョン形式が有効かチェック
func (r *TestRunner) isValidVersionFormat(version string) bool {
	// 正規表現で "1.XX" 形式をチェック
	matched, _ := regexp.MatchString(`^1\.\d+$`, version)
	return matched
}

// compareVersions バージョンを比較（v1 > v2 なら正の数を返す）
func (r *TestRunner) compareVersions(v1, v2 string) int {
	// "1.25" -> [1, 25] のように分割
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		n1, err1 := strconv.Atoi(parts1[i])
		n2, err2 := strconv.Atoi(parts2[i])

		if err1 != nil || err2 != nil {
			return strings.Compare(v1, v2)
		}

		if n1 != n2 {
			return n1 - n2
		}
	}

	return len(parts1) - len(parts2)
}

// findGoFiles ディレクトリ内の.goファイルを検索
func (r *TestRunner) findGoFiles(dir string) ([]string, error) {
	var goFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// ファイル名順にソート
	sort.Strings(goFiles)
	return goFiles, nil
}