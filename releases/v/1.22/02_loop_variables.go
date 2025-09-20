// Go 1.22 新機能: Enhanced for-loop variable semantics
// 原文: "Loop variables are now captured properly in closures within for loops"
//
// 説明: Go 1.22では、forループ内のクロージャーで変数が正しくキャプチャされるようになり、
// 長年の混乱の原因となっていた問題が解決されました。
//
// 参考リンク:
// - Go 1.22 Release Notes: https://go.dev/doc/go1.22#language
// - Go FAQ: https://go.dev/doc/faq#closures_and_goroutines

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Enhanced Loop Variables Demo ===")

	// Go 1.22以前の問題（現在は修正済み）
	demonstrateOldProblem()

	// Go 1.22の改善された動作
	demonstrateNewBehavior()

	// 実用的な例
	demonstratePracticalUse()
}

func demonstrateOldProblem() {
	fmt.Println("\n--- Go 1.22以前の問題（現在は修正済み） ---")

	// Go 1.22以前では、全ての関数が同じ変数を参照していた
	var functions []func()

	for i := 0; i < 3; i++ {
		// Go 1.22: 各イテレーションでiが新しい変数になる
		functions = append(functions, func() {
			fmt.Printf("  Go 1.22では正しく出力: %d\n", i)
		})
	}

	for _, f := range functions {
		f()
	}

	fmt.Println("  注意: Go 1.22以前では全て「3」が出力されていました")
}

func demonstrateNewBehavior() {
	fmt.Println("\n--- Go 1.22の改善された動作 ---")

	// スライス内の各要素を非同期で処理
	items := []string{"apple", "banana", "cherry"}
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Go 1.22: i と item が正しくキャプチャされる
			time.Sleep(time.Duration(i*10) * time.Millisecond)
			fmt.Printf("  処理完了: %d番目の%s\n", i, item)
		}()
	}

	wg.Wait()
	fmt.Println("  全ての非同期処理が完了")
}

func demonstratePracticalUse() {
	fmt.Println("\n--- 実用的な使用例 ---")

	// イベントハンドラーの登録
	fmt.Println("1. イベントハンドラーパターン:")
	buttons := []string{"保存", "削除", "編集"}
	var handlers []func()

	for i, label := range buttons {
		// Go 1.22: ラベルとインデックスが正しく保持される
		handler := func() {
			fmt.Printf("   ボタン[%d]「%s」がクリックされました\n", i, label)
		}
		handlers = append(handlers, handler)
	}

	// ハンドラーを実行
	for i, handler := range handlers {
		fmt.Printf("   ボタン%d実行:", i)
		handler()
	}

	// 設定値の処理
	fmt.Println("\n2. 設定値処理パターン:")
	configs := map[string]string{
		"database": "localhost:5432",
		"redis":    "localhost:6379",
		"api":      "localhost:8080",
	}

	var validators []func() bool

	for service, endpoint := range configs {
		// Go 1.22: service と endpoint が正しくキャプチャ
		validator := func() bool {
			fmt.Printf("   %sサービス（%s）の接続チェック\n", service, endpoint)
			return true // 実際の接続チェックをシミュレート
		}
		validators = append(validators, validator)
	}

	fmt.Println("   設定値の検証実行:")
	for _, validate := range validators {
		validate()
	}

	// ファイル処理のパターン
	fmt.Println("\n3. ファイル処理パターン:")
	files := []string{"config.yaml", "data.json", "log.txt"}
	var processors []func()

	for index, filename := range files {
		// Go 1.22: filename と index が正しく保持
		processor := func() {
			fmt.Printf("   [%d] %sを処理中...\n", index, filename)
			// 実際のファイル処理をシミュレート
			time.Sleep(1 * time.Millisecond)
			fmt.Printf("   [%d] %sの処理完了\n", index, filename)
		}
		processors = append(processors, processor)
	}

	for _, process := range processors {
		process()
	}

	fmt.Println("\n--- 修正前後の比較 ---")
	fmt.Println("Go 1.22以前: 開発者は明示的にループ変数をコピーする必要があった")
	fmt.Println("Go 1.22以降: 言語が自動的に正しい動作を保証")
	fmt.Println("")
	fmt.Println("修正前のワークアラウンド例:")
	fmt.Println("  for i := range items {")
	fmt.Println("    i := i  // 明示的なコピーが必要だった")
	fmt.Println("    go func() { use(i) }()")
	fmt.Println("  }")
	fmt.Println("")
	fmt.Println("Go 1.22以降:")
	fmt.Println("  for i := range items {")
	fmt.Println("    go func() { use(i) }()  // 自動的に正しく動作")
	fmt.Println("  }")
}

// % go run 02_loop_variables.go
// === Enhanced Loop Variables Demo ===
//
// --- Go 1.22以前の問題（現在は修正済み） ---
//   Go 1.22では正しく出力: 0
//   Go 1.22では正しく出力: 1
//   Go 1.22では正しく出力: 2
//   注意: Go 1.22以前では全て「3」が出力されていました
//
// --- Go 1.22の改善された動作 ---
//   処理完了: 0番目のapple
//   処理完了: 1番目のbanana
//   処理完了: 2番目のcherry
//   全ての非同期処理が完了