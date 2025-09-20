// Go 1.18 新機能: Generics (Type Parameters)
// 原文: "Type parameters for functions and types"
//
// 説明: Go 1.18では、ついにジェネリクス（型パラメーター）が追加され、
// 型安全で再利用可能なコードが書けるようになりました。
//
// 参考リンク:
// - Go 1.18 Release Notes: https://go.dev/doc/go1.18#generics
// - Go Language Specification: https://go.dev/ref/spec#Type_parameters

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Generics Demo ===")

	// ジェネリック関数の使用
	fmt.Println("--- ジェネリック関数 ---")
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))

	// ジェネリック型の使用
	fmt.Println("\n--- ジェネリック型 ---")

	// 整数のスタック
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("スタック: ")
	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// 文字列のスタック
	stringStack := NewStack[string]()
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Printf("文字列スタック: ")
	for !stringStack.IsEmpty() {
		val, _ := stringStack.Pop()
		fmt.Printf("%s ", val)
	}
	fmt.Println()

	fmt.Println("\n利点:")
	fmt.Println("  - 型安全性")
	fmt.Println("  - コードの再利用")
	fmt.Println("  - パフォーマンス向上")
	fmt.Println("  - interface{}の削減")
}

// ジェネリック関数：最大値を返す（順序可能な型制約が必要）
func Max[T interface{ ~int | ~float64 | ~string }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ジェネリック型：スタック
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// % go run 01_generics.go
// === Generics Demo ===
// --- ジェネリック関数 ---
// Max(10, 20): 20
// Max(3.14, 2.71): 3.14
// Max("apple", "banana"): banana
//
// --- ジェネリック型 ---
// スタック: 3 2 1
// 文字列スタック: world hello