// Go 1.21 新機能: slices package
// 原文: "New slices package for common slice operations"
//
// 説明: Go 1.21では、スライス操作のためのslicesパッケージが標準ライブラリに追加されました。
//
// 参考リンク:
// - Go 1.21 Release Notes: https://go.dev/doc/go1.21#slices
// - slices Package: https://pkg.go.dev/slices

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("=== slices Package Demo ===")

	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("元のスライス: %v\n", numbers)

	// ソート
	slices.Sort(numbers)
	fmt.Printf("ソート後: %v\n", numbers)

	// 検索
	fmt.Printf("4が含まれている: %v\n", slices.Contains(numbers, 4))
	fmt.Printf("3のインデックス: %d\n", slices.Index(numbers, 3))

	// 最大・最小
	fmt.Printf("最大値: %d\n", slices.Max(numbers))
	fmt.Printf("最小値: %d\n", slices.Min(numbers))

	// 比較
	other := []int{1, 1, 2, 3, 4, 5, 6, 9}
	fmt.Printf("他のスライスと等しい: %v\n", slices.Equal(numbers, other))

	fmt.Println("\n利点:")
	fmt.Println("  - 一般的なスライス操作の標準化")
	fmt.Println("  - 型安全な実装")
	fmt.Println("  - 最適化されたパフォーマンス")
}

// % go run 02_slices_package.go
// === slices Package Demo ===
// 元のスライス: [3 1 4 1 5 9 2 6]
// ソート後: [1 1 2 3 4 5 6 9]
// 4が含まれている: true
// 3のインデックス: 3