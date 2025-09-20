// Go 1.20 新機能: Comparable types in type parameters
// 原文: "Enhanced comparable constraint for generic types"
//
// 説明: Go 1.20では、ジェネリクスのcomparable制約が改善され、
// より柔軟な型パラメーターが使用できるようになりました。
//
// 参考リンク:
// - Go 1.20 Release Notes: https://go.dev/doc/go1.20#language
// - Go Language Specification: https://go.dev/ref/spec#Type_constraints

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Comparable Types Demo ===")

	// 基本的な使用例
	numbers := []int{1, 2, 3, 2, 1}
	unique := removeDuplicates(numbers)
	fmt.Printf("重複除去: %v → %v\n", numbers, unique)

	words := []string{"apple", "banana", "apple", "cherry"}
	uniqueWords := removeDuplicates(words)
	fmt.Printf("重複除去: %v → %v\n", words, uniqueWords)

	fmt.Println("\n利点:")
	fmt.Println("  - より柔軟なジェネリクス")
	fmt.Println("  - 型安全な比較操作")
	fmt.Println("  - コードの再利用性向上")
}

// comparable制約を使用した重複除去関数
func removeDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// % go run 01_comparable_types.go
// === Comparable Types Demo ===
// 重複除去: [1 2 3 2 1] → [1 2 3]
// 重複除去: [apple banana apple cherry] → [apple banana cherry]