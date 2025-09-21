// Go 1.25 新機能: go doc -http オプション
// 原文: "New go doc -http option to start documentation server"
//
// 説明: Go 1.25では、go docコマンドに-httpオプションが追加され、
// ローカルでドキュメントサーバーを起動できるようになりました。
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#cmd-go-doc
// - go doc Command: https://pkg.go.dev/cmd/go#hdr-Show_documentation_for_package_or_symbol

// +build ignore

package main

// このファイルを実行するには: go run 07_go_doc_http.go

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== go doc -http オプション Demo ===")

	fmt.Println("Go 1.25で追加されたgo docの新しい-httpオプション")

	fmt.Println("\n--- go doc -httpの特徴 ---")
	fmt.Println("1. ローカルドキュメントサーバー")
	fmt.Println("   - ブラウザでGoドキュメントを表示")
	fmt.Println("   - インターネット接続不要")

	fmt.Println("\n2. プロジェクト固有のドキュメント")
	fmt.Println("   - 現在のモジュールのドキュメントを表示")
	fmt.Println("   - 依存関係のドキュメントも含む")

	fmt.Println("\n3. リアルタイム更新")
	fmt.Println("   - コード変更時の自動更新")
	fmt.Println("   - 開発中のドキュメント確認に便利")

	fmt.Println("\n--- 基本的な使用方法 ---")
	fmt.Println("1. 基本起動:")
	fmt.Println("   go doc -http=:8080")
	fmt.Println("   # http://localhost:8080 でアクセス")

	fmt.Println("\n2. 特定のホスト指定:")
	fmt.Println("   go doc -http=localhost:9000")

	fmt.Println("\n3. 全てのインターフェースでリッスン:")
	fmt.Println("   go doc -http=:6060")

	fmt.Println("\n--- 実用例 ---")
	demonstrateGoDocCommands()

	fmt.Println("\n--- 従来のgodocとの違い ---")
	fmt.Println("従来:")
	fmt.Println("  - 別途godocツールをインストール")
	fmt.Println("  - godoc -http=:6060")

	fmt.Println("\nGo 1.25以降:")
	fmt.Println("  - go docコマンドに統合")
	fmt.Println("  - go doc -http=:6060")
	fmt.Println("  - 追加インストール不要")

	fmt.Println("\n--- プロジェクト開発での活用 ---")
	fmt.Println("1. API仕様書として活用")
	fmt.Println("2. チームでのドキュメント共有")
	fmt.Println("3. コードレビュー時の参考資料")
	fmt.Println("4. 新人メンバーのオンボーディング")

	// 実際にgo docコマンドを試す（デモ用）
	tryGoDocCommand()
}

func demonstrateGoDocCommands() {
	fmt.Println("\nよく使用されるコマンド例:")

	commands := []struct {
		command     string
		description string
	}{
		{"go doc -http=:8080", "ポート8080でドキュメントサーバーを起動"},
		{"go doc -http=localhost:9000", "localhost:9000でサーバー起動"},
		{"go doc -http=:6060 .", "現在のディレクトリのプロジェクトのドキュメント"},
		{"go doc -http=:7070 -src", "ソースコード付きでドキュメント表示"},
	}

	for _, cmd := range commands {
		fmt.Printf("  %-35s # %s\n", cmd.command, cmd.description)
	}

	fmt.Println("\n--- アクセス例 ---")
	fmt.Println("サーバー起動後、ブラウザで以下にアクセス:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/pkg/")
	fmt.Println("  http://localhost:8080/pkg/fmt/")
	fmt.Println("  http://localhost:8080/src/")
}

func tryGoDocCommand() {
	fmt.Println("\ngo docコマンドの使用例...")

	fmt.Println("現在のGo version: go1.25 (demo)")

	fmt.Println("\n--- fmt.Printlnのドキュメント例 ---")
	fmt.Println("func Println(a ...any) (n int, err error)")
	fmt.Println("Println formats using the default formats for its operands and writes to")
	fmt.Println("standard output. Spaces are always added between operands and a newline")
	fmt.Println("is appended. It returns the number of bytes written and any write error")
	fmt.Println("encountered.")

	fmt.Println("\n--- Go 1.25での新しい使い方 ---")
	fmt.Println("以下のコマンドでWebサーバーとして起動:")
	fmt.Println("  go doc -http=:8080")
	fmt.Println("")
	fmt.Println("その後ブラウザで http://localhost:8080 にアクセス")
	fmt.Println("※ 実際のGo 1.25環境でのみ利用可能")

	// サンプルのHTTPサーバー起動シミュレーション
	fmt.Println("\n[シミュレーション] ドキュメントサーバー起動中...")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Server starting... %d/3\n", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("ドキュメントサーバーが起動しました (シミュレーション)")
	fmt.Println("ブラウザでアクセス: http://localhost:8080")
	fmt.Println("停止するには Ctrl+C を押してください")
}