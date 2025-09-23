package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// TestResult 単一テストの結果
type TestResult struct {
	TestName    string `json:"test_name"`
	Version     string `json:"version"`
	Status      string `json:"status"` // PASS, FAIL, SKIP
	Error       string `json:"error,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
}

// IsSuccess テストが成功したかどうか
func (tr *TestResult) IsSuccess() bool {
	return tr.Status == "PASS"
}

// IsFailure テストが失敗したかどうか
func (tr *TestResult) IsFailure() bool {
	return tr.Status == "FAIL"
}

// TestResults 全テスト結果
type TestResults struct {
	TotalTests   int
	SuccessTests int
	ErrorTests   int
	Results      []*TestResult
}

// NewTestResults テスト結果を初期化
func NewTestResults() *TestResults {
	return &TestResults{
		Results: make([]*TestResult, 0),
	}
}

// AddVersionResults バージョン別結果を追加
func (tr *TestResults) AddVersionResults(version string, results []*TestResult) {
	for _, result := range results {
		tr.addResult(result)
	}
}

// AddBasicResults 基本テスト結果を追加
func (tr *TestResults) AddBasicResults(results []*TestResult) {
	for _, result := range results {
		tr.addResult(result)
	}
}

// addResult 単一結果を追加
func (tr *TestResults) addResult(result *TestResult) {
	tr.Results = append(tr.Results, result)
	tr.TotalTests++

	switch result.Status {
	case "PASS":
		tr.SuccessTests++
	case "FAIL":
		tr.ErrorTests++
	}
}

// HasFailures 失敗したテストがあるかどうか
func (tr *TestResults) HasFailures() bool {
	return tr.ErrorTests > 0
}

// GetFailures 失敗したテストを取得
func (tr *TestResults) GetFailures() []*TestResult {
	var failures []*TestResult
	for _, result := range tr.Results {
		if result.IsFailure() {
			failures = append(failures, result)
		}
	}
	return failures
}

// PrintFailures 失敗したテストを表示
func (tr *TestResults) PrintFailures() {
	failures := tr.GetFailures()
	for _, failure := range failures {
		fmt.Printf("テスト: %s (Go %s)\n", failure.TestName, failure.Version)
	}
}

// VersionGroup バージョンごとのテスト結果グループ
type VersionGroup struct {
	Version string
	Results []*TestResult
}

// GroupByVersion バージョンごとにグループ化
func (tr *TestResults) GroupByVersion() []*VersionGroup {
	versionMap := make(map[string][]*TestResult)

	for _, result := range tr.Results {
		versionMap[result.Version] = append(versionMap[result.Version], result)
	}

	var groups []*VersionGroup

	// バージョンを取得して降順ソート
	var versions []string
	for version := range versionMap {
		versions = append(versions, version)
	}

	// バージョン順にソート（降順：新しいバージョンから）
	sort.Slice(versions, func(i, j int) bool {
		return compareVersions(versions[i], versions[j]) > 0
	})

	for _, version := range versions {
		if results, exists := versionMap[version]; exists {
			groups = append(groups, &VersionGroup{
				Version: version,
				Results: results,
			})
		}
	}

	return groups
}

// compareVersions バージョンを比較（v1 > v2 なら正の数を返す）
func compareVersions(v1, v2 string) int {
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

// Summary 結果サマリーを文字列で取得
func (tr *TestResults) Summary() string {
	var sb strings.Builder

	sb.WriteString("=== テスト結果サマリー ===\n")
	sb.WriteString(fmt.Sprintf("総テスト数: %d\n", tr.TotalTests))
	sb.WriteString(fmt.Sprintf("成功: %d\n", tr.SuccessTests))
	sb.WriteString(fmt.Sprintf("エラー: %d\n", tr.ErrorTests))

	if tr.TotalTests > 0 {
		successRate := float64(tr.SuccessTests) * 100.0 / float64(tr.TotalTests)
		sb.WriteString(fmt.Sprintf("成功率: %.1f%%\n", successRate))
	}

	return sb.String()
}
