package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GenerateReport テスト結果レポートを生成
func (r *TestRunner) GenerateReport(results *TestResults) error {
	// テキストレポート生成
	if err := r.generateTextReport(results); err != nil {
		return fmt.Errorf("failed to generate text report: %w", err)
	}

	// エラーレポート生成
	if err := r.generateErrorReport(results); err != nil {
		return fmt.Errorf("failed to generate error report: %w", err)
	}

	fmt.Printf("\n=== 実行完了 ===\n")
	fmt.Printf("詳細結果: %s\n", filepath.Join(r.OutputDir, "integration_test_results.txt"))
	fmt.Printf("エラー詳細: %s\n", filepath.Join(r.OutputDir, "integration_test_errors.txt"))

	return nil
}

// generateTextReport テキスト形式のレポートを生成
func (r *TestRunner) generateTextReport(results *TestResults) error {
	filePath := filepath.Join(r.OutputDir, "integration_test_results.txt")

	var sb strings.Builder

	// ヘッダー
	sb.WriteString("=== Go Release Tour API経由統合テスト開始 ===\n")
	sb.WriteString(fmt.Sprintf("開始時刻: %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006")))
	sb.WriteString(fmt.Sprintf("API URL: %s\n\n", r.APIURL))

	// バージョンごとの結果
	versionGroups := results.GroupByVersion()
	for _, group := range versionGroups {
		sb.WriteString(fmt.Sprintf("=== Go %s のテスト ===\n", group.Version))
		sb.WriteString(fmt.Sprintf("テストファイル数: %8d\n", len(group.Results)))

		for _, result := range group.Results {
			status := "[PASS]"
			if result.IsFailure() {
				status = "[FAIL]"
			} else if result.Status == "SKIP" {
				status = "[SKIP]"
			}

			sb.WriteString(fmt.Sprintf("  API テスト: %s (Go %s) ... %s\n",
				result.TestName, result.Version, status))

			// 詳細情報（verboseモードまたは失敗時）
			if r.Verbose || result.IsFailure() {
				if result.Error != "" {
					sb.WriteString(fmt.Sprintf("    エラー: %s\n", result.Error))
				}
				if result.RawResponse != "" && r.Verbose {
					// レスポンスの最初の100文字のみ
					response := result.RawResponse
					if len(response) > 100 {
						response = response[:100] + "..."
					}
					sb.WriteString(fmt.Sprintf("    レスポンス: %s\n", response))
				}
			}
		}
		sb.WriteString("\n")
	}

	// 基本APIテスト
	basicResults := r.getBasicAPIResults(results)
	if len(basicResults) > 0 {
		sb.WriteString("=== 基本APIテスト ===\n")
		for _, result := range basicResults {
			status := "[PASS]"
			if result.IsFailure() {
				status = "[FAIL]"
			}
			sb.WriteString(fmt.Sprintf("  API テスト: %s (Go %s) ... %s\n",
				result.TestName, result.Version, status))
		}
		sb.WriteString("\n")
	}

	// サマリー
	sb.WriteString(results.Summary())
	sb.WriteString(fmt.Sprintf("\n終了時刻: %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006")))

	// ファイルに書き込み
	return os.WriteFile(filePath, []byte(sb.String()), 0600)
}

// generateErrorReport エラーレポートを生成
func (r *TestRunner) generateErrorReport(results *TestResults) error {
	filePath := filepath.Join(r.OutputDir, "integration_test_errors.txt")

	failures := results.GetFailures()
	if len(failures) == 0 {
		// エラーがない場合は空ファイルを作成
		return os.WriteFile(filePath, []byte(""), 0600)
	}

	var sb strings.Builder

	for _, failure := range failures {
		sb.WriteString("[FAIL]\n")
		sb.WriteString(fmt.Sprintf("テスト: %s (Go %s)\n", failure.TestName, failure.Version))

		if failure.RawResponse != "" {
			sb.WriteString(fmt.Sprintf("レスポンス: %s\n", failure.RawResponse))
		} else if failure.Error != "" {
			sb.WriteString(fmt.Sprintf("エラー: %s\n", failure.Error))
		}

		sb.WriteString("---\n")
	}

	return os.WriteFile(filePath, []byte(sb.String()), 0600)
}

// getBasicAPIResults 基本APIテストの結果を取得
func (r *TestRunner) getBasicAPIResults(results *TestResults) []*TestResult {
	var basicResults []*TestResult
	for _, result := range results.Results {
		if result.TestName == "Basic Hello World" {
			basicResults = append(basicResults, result)
		}
	}
	return basicResults
}
