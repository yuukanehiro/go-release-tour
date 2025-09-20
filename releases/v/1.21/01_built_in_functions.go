// Go 1.21 新機能: Built-in functions (min, max, clear)
// 原文: "New built-in functions min, max, and clear"
//
// 説明: Go 1.21では、min、max、clearの組み込み関数が追加され、
// より簡潔なコードが書けるようになりました。
//
// 参考リンク:
// - Go 1.21 Release Notes: https://go.dev/doc/go1.21#language
// - Go Language Specification: https://go.dev/ref/spec#Built-in_functions

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Built-in Functions Demo ===")

	// min関数
	fmt.Println("--- min関数 ---")
	fmt.Printf("min(3, 1, 4): %d\n", min(3, 1, 4))
	fmt.Printf("min(2.5, 1.2): %.1f\n", min(2.5, 1.2))

	// max関数
	fmt.Println("\n--- max関数 ---")
	fmt.Printf("max(3, 1, 4): %d\n", max(3, 1, 4))
	fmt.Printf("max(2.5, 1.2): %.1f\n", max(2.5, 1.2))

	// clear関数
	fmt.Println("\n--- clear関数 ---")

	// スライスのクリア
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("クリア前: %v\n", slice)
	clear(slice)
	fmt.Printf("クリア後: %v\n", slice)

	// マップのクリア
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("マップ クリア前: %v\n", m)
	clear(m)
	fmt.Printf("マップ クリア後: %v\n", m)

	fmt.Println("\n利点:")
	fmt.Println("  - 標準ライブラリ不要")
	fmt.Println("  - 型安全")
	fmt.Println("  - 高性能")
}

// % go run 01_built_in_functions.go
// === Built-in Functions Demo ===
// --- min関数 ---
// min(3, 1, 4): 1
// min(2.5, 1.2): 1.2
//
// --- max関数 ---
// max(3, 1, 4): 4
// max(2.5, 1.2): 2.5