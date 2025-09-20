// Go 1.18 新機能: Workspace Mode
// 原文: "Go 1.18 adds workspace mode to Go, which lets you work on multiple modules simultaneously"
//
// 説明: Go 1.18では、ワークスペースモードが追加され、
// 複数のモジュールを同時に開発できるようになりました。
//
// 参考リンク:
// - Go 1.18 Release Notes: https://go.dev/doc/go1.18#workspace
// - Go Modules Reference: https://go.dev/ref/mod#workspaces

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Workspace Mode Demo ===")

	fmt.Println("\n--- ワークスペースモードとは ---")
	fmt.Println("Go 1.18のワークスペースモードでは、")
	fmt.Println("複数のモジュールを同時に開発・テストできます。")
	fmt.Println("go.workファイルを使用してモジュール間の依存関係を管理します。")

	fmt.Println("\n--- go.workファイルの構造 ---")

	// go.workファイルの例を表示
	workFileExample := `go 1.18

use (
    ./api
    ./client
    ./shared
)

replace (
    example.com/api => ./api
    example.com/client => ./client
    example.com/shared => ./shared
)`

	fmt.Printf("go.workファイルの例:\n%s\n", workFileExample)

	fmt.Println("\n--- ワークスペースコマンド ---")

	commands := []struct {
		command     string
		description string
		example     string
	}{
		{
			command:     "go work init",
			description: "新しいワークスペースを初期化",
			example:     "go work init ./module1 ./module2",
		},
		{
			command:     "go work use",
			description: "モジュールをワークスペースに追加",
			example:     "go work use ./new-module",
		},
		{
			command:     "go work sync",
			description: "ワークスペースの依存関係を同期",
			example:     "go work sync",
		},
		{
			command:     "go work edit",
			description: "go.workファイルを編集",
			example:     "go work edit -replace old.com/mod=./local/mod",
		},
	}

	for _, cmd := range commands {
		fmt.Printf("• %s\n", cmd.command)
		fmt.Printf("  説明: %s\n", cmd.description)
		fmt.Printf("  例: %s\n\n", cmd.example)
	}

	fmt.Println("--- 実際のプロジェクト構造例 ---")

	projectStructure := `myproject/
├── go.work
├── api/
│   ├── go.mod
│   ├── main.go
│   └── handlers/
├── client/
│   ├── go.mod
│   ├── main.go
│   └── config/
└── shared/
    ├── go.mod
    ├── models/
    └── utils/`

	fmt.Printf("%s\n", projectStructure)

	fmt.Println("\n--- ワークスペースの利点 ---")
	benefits := []string{
		"複数モジュールの同時開発",
		"ローカル依存関係の簡単な管理",
		"replace ディレクティブが不要",
		"モジュール間の変更の即座反映",
		"統一されたビルド・テスト環境",
	}

	for i, benefit := range benefits {
		fmt.Printf("%d. %s\n", i+1, benefit)
	}

	fmt.Println("\n--- 従来の方法との比較 ---")

	fmt.Println("従来の方法（Go 1.17以前）:")
	fmt.Println("  - 各モジュールのgo.modにreplaceディレクティブを追加")
	fmt.Println("  - 依存関係の変更時に複数ファイルを更新")
	fmt.Println("  - モジュール間の開発が複雑")

	fmt.Println("\nワークスペースモード（Go 1.18+）:")
	fmt.Println("  - 単一のgo.workファイルで管理")
	fmt.Println("  - 自動的なローカル依存関係解決")
	fmt.Println("  - 簡潔な開発ワークフロー")

	fmt.Println("\n--- 使用例シナリオ ---")

	scenarios := []struct {
		title       string
		description string
		modules     []string
	}{
		{
			title:       "マイクロサービス開発",
			description: "API、フロントエンド、共通ライブラリを同時開発",
			modules:     []string{"user-service", "auth-service", "shared-lib"},
		},
		{
			title:       "ライブラリ開発",
			description: "ライブラリとその使用例を同時開発",
			modules:     []string{"mylib", "examples", "benchmarks"},
		},
		{
			title:       "プラグインシステム",
			description: "コアシステムと複数プラグインを同時開発",
			modules:     []string{"core", "plugin-auth", "plugin-storage"},
		},
	}

	for _, scenario := range scenarios {
		fmt.Printf("• %s\n", scenario.title)
		fmt.Printf("  %s\n", scenario.description)
		fmt.Printf("  モジュール: %v\n\n", scenario.modules)
	}

	fmt.Println("--- ワークスペース設定のベストプラクティス ---")

	bestPractices := []string{
		"go.workファイルはバージョン管理から除外（.gitignore）",
		"相対パスを使用してポータビリティを確保",
		"定期的にgo work syncで依存関係を同期",
		"チーム開発では各開発者が独自のgo.workを作成",
		"CI/CDではワークスペースモードを無効化",
	}

	for i, practice := range bestPractices {
		fmt.Printf("%d. %s\n", i+1, practice)
	}

	fmt.Println("\n--- 環境変数とフラグ ---")

	envVars := []struct {
		name        string
		description string
	}{
		{"GOWORK", "使用するgo.workファイルのパスを指定"},
		{"GOWORK=off", "ワークスペースモードを無効化"},
	}

	fmt.Println("重要な環境変数:")
	for _, env := range envVars {
		fmt.Printf("• %s: %s\n", env.name, env.description)
	}

	fmt.Println("\nビルドフラグ:")
	fmt.Println("• -workfile: 特定のgo.workファイルを指定")

	fmt.Println("\n--- 実践的なワークフロー ---")

	workflow := []string{
		"1. プロジェクトルートでgo work initを実行",
		"2. 各モジュールをgo work useで追加",
		"3. 通常通りgo build/go testを実行",
		"4. 依存関係変更時はgo work syncで同期",
		"5. リリース前にワークスペースモードを無効化してテスト",
	}

	for _, step := range workflow {
		fmt.Printf("%s\n", step)
	}

	fmt.Println("\n--- 注意点とトラブルシューティング ---")

	warnings := []string{
		"go.workファイルは開発時のみ使用",
		"本番環境ではワークスペースモードを無効化",
		"モジュールパスの競合に注意",
		"大きなワークスペースではビルド時間が増加する可能性",
	}

	for _, warning := range warnings {
		fmt.Printf("⚠️  %s\n", warning)
	}

	// デモ用のファイル作成例
	fmt.Println("\n--- デモ: 簡単なgo.workファイル ---")
	demoWorkspace()
}

func demoWorkspace() {
	// 一時的なディレクトリ構造をシミュレート
	fmt.Println("仮想的なワークスペース構造:")

	structure := map[string][]string{
		"./api": {
			"go.mod: module example.com/api",
			"main.go: APIサーバーのエントリーポイント",
			"handlers/: HTTPハンドラー",
		},
		"./client": {
			"go.mod: module example.com/client",
			"main.go: クライアントアプリケーション",
			"config/: 設定ファイル",
		},
		"./shared": {
			"go.mod: module example.com/shared",
			"models/: 共通データモデル",
			"utils/: 共通ユーティリティ",
		},
	}

	for path, files := range structure {
		fmt.Printf("\n%s/\n", path)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	fmt.Println("\n対応するgo.workファイル:")
	workContent := `go 1.18

use (
    ./api
    ./client
    ./shared
)

// 必要に応じてreplace文を追加
replace example.com/shared => ./shared`

	fmt.Printf("%s\n", workContent)

	fmt.Println("\nこの設定により:")
	fmt.Println("• apiモジュールがsharedモジュールを直接参照可能")
	fmt.Println("• clientモジュールがsharedモジュールを直接参照可能")
	fmt.Println("• 各モジュールの変更が即座に他のモジュールに反映")
	fmt.Println("• 統一されたビルド・テスト環境")
}

// 実際のファイル操作の例（コメントアウト）
func createExampleWorkspace() {
	// 実際のプロジェクトではこのような操作を行う
	/*
	// ワークスペースの初期化
	workDir := "./example-workspace"
	os.MkdirAll(workDir, 0755)
	os.Chdir(workDir)

	// go.workファイルの作成
	workContent := `go 1.18

use (
    ./api
    ./client
    ./shared
)
`
	os.WriteFile("go.work", []byte(workContent), 0644)

	// 各モジュールディレクトリの作成
	modules := []string{"api", "client", "shared"}
	for _, module := range modules {
		os.MkdirAll(module, 0755)

		// 各モジュールのgo.modファイル
		modContent := fmt.Sprintf("module example.com/%s\n\ngo 1.18\n", module)
		os.WriteFile(filepath.Join(module, "go.mod"), []byte(modContent), 0644)
	}
	*/
}

// % go run 02_workspace_mode.go
// === Workspace Mode Demo ===
//
// --- ワークスペースモードとは ---
// Go 1.18のワークスペースモードでは、
// 複数のモジュールを同時に開発・テストできます。
// go.workファイルを使用してモジュール間の依存関係を管理します。
//
// --- go.workファイルの構造 ---
// go.workファイルの例:
// go 1.18
//
// use (
//     ./api
//     ./client
//     ./shared
// )
//
// replace (
//     example.com/api => ./api
//     example.com/client => ./client
//     example.com/shared => ./shared
// )