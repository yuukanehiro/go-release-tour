// Go 1.25 新機能: go.mod ignore ディレクティブ
// 原文: "The go.mod file format now supports an ignore directive"
//
// 説明: Go 1.25では、go.modファイルで特定のモジュールを無視するignoreディレクティブが追加されました。
// これにより、セキュリティ上の問題があるモジュールや使用したくないモジュールを明示的に除外できます。
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#go-mod
// - Go Modules Reference: https://go.dev/ref/mod

//go:build ignore
// +build ignore

package main

// このファイルを実行するには: go run 04_go_mod_ignore.go

import (
	"fmt"
)

func main() {
	fmt.Println("=== go.mod ignore ディレクティブ Demo ===")

	fmt.Println("Go 1.25で追加されたgo.modのignoreディレクティブ")

	// サンプルのgo.modファイル内容を表示
	displaySampleGoMod()

	fmt.Println("\n--- ignoreディレクティブの特徴 ---")
	fmt.Println("1. セキュリティリスクのあるモジュール除外")
	fmt.Println("   - 脆弱性が発見されたパッケージを無視")
	fmt.Println("   - 信頼できないソースのパッケージを除外")

	fmt.Println("\n2. 開発ポリシーの強制")
	fmt.Println("   - 特定のライセンスのパッケージを禁止")
	fmt.Println("   - 非推奨パッケージの使用を防止")

	fmt.Println("\n3. 依存関係の制御")
	fmt.Println("   - 間接依存関係での問題パッケージを除外")
	fmt.Println("   - チーム開発での統一されたパッケージ管理")

	fmt.Println("\n--- 実用的な使用ケース ---")
	fmt.Println("1. セキュリティ対策:")
	fmt.Println("   ignore example.com/vulnerable-package")
	fmt.Println("")
	fmt.Println("2. ライセンス管理:")
	fmt.Println("   ignore example.com/gpl-licensed-package")
	fmt.Println("")
	fmt.Println("3. 非推奨パッケージの除外:")
	fmt.Println("   ignore example.com/deprecated-package")
	fmt.Println("")
	fmt.Println("4. 開発環境の統一:")
	fmt.Println("   ignore example.com/development-only-package")

	fmt.Println("\n--- goコマンドでの動作 ---")
	fmt.Println("- go get: ignoreされたパッケージは取得されない")
	fmt.Println("- go mod tidy: ignoreされたパッケージは除外される")
	fmt.Println("- go build: ignoreされたパッケージがあるとエラー")

}

func displaySampleGoMod() {
	fmt.Println("\n--- サンプルgo.modファイルの例 ---")
	fmt.Println("=====================================")

	goModContent := `module example.com/my-project

go 1.25

// 通常の依存関係
require (
    github.com/gorilla/mux v1.8.0
    github.com/stretchr/testify v1.8.0
)

// Go 1.25の新機能: ignoreディレクティブ
// セキュリティ上の問題があるパッケージを無視
ignore (
    example.com/vulnerable-package
    github.com/insecure/old-crypto
)

// 単一パッケージの無視も可能
ignore example.com/deprecated-library

// ライセンス上の問題があるパッケージを無視
ignore (
    example.com/gpl-package
    example.com/commercial-only
)`

	fmt.Println(goModContent)
	fmt.Println("=====================================")
}

// % go run 04_go_mod_ignore.go
// === go.mod ignore ディレクティブ Demo ===
// Go 1.25で追加されたgo.modのignoreディレクティブ

// サンプルgo.modファイルを作成...
// サンプルgo.modファイルを作成しました

// --- ignoreディレクティブの特徴 ---
// 1. セキュリティリスクのあるモジュール除外
//    - 脆弱性が発見されたパッケージを無視
//    - 信頼できないソースのパッケージを除外

// 2. 開発ポリシーの強制
//    - 特定のライセンスのパッケージを禁止
//    - 非推奨パッケージの使用を防止

// 3. 依存関係の制御
//    - 間接依存関係での問題パッケージを除外
//    - チーム開発での統一されたパッケージ管理

// --- 使用例 ---
// 作成されたgo.modファイルの内容:
// =====================================
// module example.com/my-project

// go 1.25

// // 通常の依存関係
// require (
//     github.com/gorilla/mux v1.8.0
//     github.com/stretchr/testify v1.8.0
// )

// // Go 1.25の新機能: ignoreディレクティブ
// // セキュリティ上の問題があるパッケージを無視
// ignore (
//     example.com/vulnerable-package
//     github.com/insecure/old-crypto
// )

// // 単一パッケージの無視も可能
// ignore example.com/deprecated-library

// // ライセンス上の問題があるパッケージを無視
// ignore (
//     example.com/gpl-package
//     example.com/commercial-only
// )

// =====================================

// --- 実用的な使用ケース ---
// 1. セキュリティ対策:
//    ignore example.com/vulnerable-package

// 2. ライセンス管理:
//    ignore example.com/gpl-licensed-package

// 3. 非推奨パッケージの除外:
//    ignore example.com/deprecated-package

// 4. 開発環境の統一:
//    ignore example.com/development-only-package

// --- goコマンドでの動作 ---
// - go get: ignoreされたパッケージは取得されない
// - go mod tidy: ignoreされたパッケージは除外される
// - go build: ignoreされたパッケージがあるとエラー

// クリーンアップ...
// サンプルファイルを削除しました
