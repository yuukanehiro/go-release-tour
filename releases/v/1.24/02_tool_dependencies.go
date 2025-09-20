// Go 1.24 新機能: Tool Dependencies in go.mod
// 原文: "Go modules can now track executable dependencies using tool directives in go.mod"
//
// 説明: go.modファイルで実行可能ツールの依存関係を追跡できるようになり、tools.goファイルが不要になりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("=== Tool Dependencies Demo ===")

	// go.modファイルの例を表示
	fmt.Println("go.modファイルでのツール依存関係の例:")
	fmt.Println("")

	goModExample := `module example.com/myproject

go 1.24

require (
    github.com/golang/protobuf v1.5.3
)

// Go 1.24の新機能: tool ディレクティブ
tool (
    golang.org/x/tools/cmd/goimports
    github.com/golangci/golangci-lint/cmd/golangci-lint
    github.com/swaggo/swag/cmd/swag
    google.golang.org/protobuf/cmd/protoc-gen-go
)

// 以前は tools.go ファイルが必要だった:
// //go:build tools
// package tools
// import (
//     _ "golang.org/x/tools/cmd/goimports"
//     _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
// )`

	fmt.Println(goModExample)

	fmt.Println("\n--- 新しいコマンドの使用例 ---")

	// go get -tool の例
	fmt.Println("1. ツールの追加:")
	fmt.Println("   go get -tool golang.org/x/tools/cmd/goimports")
	fmt.Println("")

	// go run -tool の例
	fmt.Println("2. ツールの実行:")
	fmt.Println("   go run -tool goimports -w .")
	fmt.Println("   go run -tool golangci-lint run")
	fmt.Println("")

	// 実際のコマンド実行例（安全な例のみ）
	fmt.Println("3. 現在のGoバージョン確認:")
	if output, err := exec.Command("go", "version").Output(); err == nil {
		fmt.Printf("   %s\n", strings.TrimSpace(string(output)))
	} else {
		fmt.Printf("   Error: %v\n", err)
	}

	fmt.Println("\n4. go envの一部情報:")
	if output, err := exec.Command("go", "env", "GOOS", "GOARCH").Output(); err == nil {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		for i, line := range lines {
			if i == 0 {
				fmt.Printf("   GOOS: %s\n", line)
			} else if i == 1 {
				fmt.Printf("   GOARCH: %s\n", line)
			}
		}
	}

	fmt.Println("\n--- ツール依存関係の利点 ---")
	fmt.Println("✅ tools.goファイルが不要")
	fmt.Println("✅ 明示的なツール管理")
	fmt.Println("✅ バージョン固定が可能")
	fmt.Println("✅ チーム開発での一貫性")
	fmt.Println("✅ CIでの再現可能なビルド")

	// 現在のディレクトリのgo.modファイルをチェック
	fmt.Println("\n--- 現在のディレクトリ情報 ---")
	pwd, _ := os.Getwd()
	fmt.Printf("現在のディレクトリ: %s\n", pwd)

	goModPath := filepath.Join(pwd, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		fmt.Println("✅ go.modファイルが存在します")
		if content, err := os.ReadFile(goModPath); err == nil {
			lines := strings.Split(string(content), "\n")
			fmt.Println("go.modの内容:")
			for i, line := range lines {
				if i < 5 { // 最初の5行のみ表示
					fmt.Printf("   %s\n", line)
				}
			}
			if len(lines) > 5 {
				fmt.Println("   ...")
			}
		}
	} else {
		fmt.Println("❌ go.modファイルが見つかりません")
		fmt.Println("   'go mod init <module-name>' で初期化してください")
	}

	fmt.Println("\n--- 移行例 ---")
	fmt.Println("従来の方法 (tools.go):")
	fmt.Println(`//go:build tools
package tools

import (
    _ "golang.org/x/tools/cmd/goimports"
    _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)`)

	fmt.Println("\nGo 1.24の新しい方法 (go.mod):")
	fmt.Println(`tool (
    golang.org/x/tools/cmd/goimports
    github.com/golangci/golangci-lint/cmd/golangci-lint
)`)
}

// % go run 02_tool_dependencies.go
// === Tool Dependencies Demo ===
// go.modファイルでのツール依存関係の例:
//
// module example.com/myproject
//
// go 1.24
//
// require (
//     github.com/golang/protobuf v1.5.3
// )
//
// // Go 1.24の新機能: tool ディレクティブ
// tool (
//     golang.org/x/tools/cmd/goimports
//     github.com/golangci/golangci-lint/cmd/golangci-lint
//     github.com/swaggo/swag/cmd/swag
//     google.golang.org/protobuf/cmd/protoc-gen-go
// )
//
// --- 新しいコマンドの使用例 ---
// 1. ツールの追加:
//    go get -tool golang.org/x/tools/cmd/goimports
//
// 2. ツールの実行:
//    go run -tool goimports -w .
//    go run -tool golangci-lint run