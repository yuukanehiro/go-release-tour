// Go 1.22 新機能: for-range over integers
// 原文: "A for-range loop can now range over integers"
//
// 説明: Go 1.22では、for-rangeループで直接整数をイテレートできるようになり、
// 従来のC風のループがより簡潔に書けるようになりました。
//
// 参考リンク:
// - Go 1.22 Release Notes: https://go.dev/doc/go1.22#language
// - Go Language Specification: https://go.dev/ref/spec#For_range

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== For-Range over Integers Demo ===")

	// Go 1.22の新機能: 整数に対するrange
	fmt.Println("1から5までの数値:")
	for i := range 5 {
		fmt.Printf("  %d\n", i)
	}

	// 従来の方法と比較
	fmt.Println("\n従来の方法:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", i)
	}

	// 実用例: 配列の初期化
	fmt.Println("\n--- 配列の初期化 ---")
	squares := make([]int, 10)
	for i := range 10 {
		squares[i] = i * i
	}
	fmt.Printf("平方数: %v\n", squares)

	// 星を描画
	fmt.Println("\n--- パターン描画 ---")
	for i := range 5 {
		for range i + 1 {
			fmt.Print("★")
		}
		fmt.Println()
	}

	fmt.Println("\n利点:")
	fmt.Println("  - より簡潔な構文")
	fmt.Println("  - 初期化・条件・インクリメントが不要")
	fmt.Println("  - 0からn-1の範囲が明確")
}

// % go run 01_for_range_integers.go
// === For-Range over Integers Demo ===
// 1から5までの数値:
//   0
//   1
//   2
//   3
//   4
//
// 従来の方法:
//   0
//   1
//   2
//   3
//   4