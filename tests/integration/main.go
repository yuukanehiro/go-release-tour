package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		apiURL    = flag.String("url", "http://localhost:8080/api/run", "API URL for testing")
		outputDir = flag.String("output", "../results", "Output directory for test results")
		verbose   = flag.Bool("v", false, "Verbose output")
	)
	flag.Parse()

	// 結果ディレクトリの作成
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// テストランナーの初期化
	runner := NewTestRunner(*apiURL, *outputDir, *verbose)

	// テスト実行
	fmt.Println("=== Go Release Tour API経由統合テスト開始 ===")
	results, err := runner.RunAllTests()
	if err != nil {
		log.Fatalf("Test execution failed: %v", err)
	}

	// 結果レポート
	if err := runner.GenerateReport(results); err != nil {
		log.Fatalf("Failed to generate report: %v", err)
	}

	// 終了ステータス
	if results.HasFailures() {
		fmt.Printf("\n[ERROR] エラーのあるテスト一覧:\n")
		results.PrintFailures()
		os.Exit(1)
	} else {
		fmt.Printf("\n[SUCCESS] 全てのテストが成功しました！\n")
		os.Exit(0)
	}
}
