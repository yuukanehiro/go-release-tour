//go:build ignore
// +build ignore

// Go 1.18 新機能: Type Inference
// 原文: "Type inference for generics"
//
// 説明: Type Inference - ジェネリクスにおける型推論の仕組み
//
// 参考リンク:
// - Go 1.18 Release Notes: https://go.dev/doc/go1.18#generics
// - Go Language Specification: https://go.dev/ref/spec#Type_inference
package main

import (
	"fmt"
)

// Identity は引数をそのまま返すジェネリック関数
func Identity[T any](x T) T {
	return x
}

// Map はスライスの各要素に関数を適用する
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter は条件に合う要素のみを残す
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce はスライスを単一の値に畳み込む
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func main() {
	// 型推論の例1: 明示的な型指定なし
	fmt.Println("=== Type Inference Examples ===")

	// コンパイラが型を推論
	x := Identity(42)      // T は int と推論
	y := Identity("hello") // T は string と推論
	z := Identity(3.14)    // T は float64 と推論

	fmt.Printf("Identity(42): %v (type: %T)\n", x, x)
	fmt.Printf("Identity(\"hello\"): %v (type: %T)\n", y, y)
	fmt.Printf("Identity(3.14): %v (type: %T)\n", z, z)

	// 型推論の例2: スライス操作
	fmt.Println("\n=== Slice Operations with Type Inference ===")

	numbers := []int{1, 2, 3, 4, 5}

	// Map: int から string への変換（型推論される）
	strings := Map(numbers, func(n int) string {
		return fmt.Sprintf("num_%d", n)
	})
	fmt.Printf("Numbers to strings: %v\n", strings)

	// Map: int から float64 への変換
	floats := Map(numbers, func(n int) float64 {
		return float64(n) * 1.5
	})
	fmt.Printf("Numbers to floats: %v\n", floats)

	// Filter: 偶数のみを残す
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// Reduce: 合計を計算
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("Sum: %d\n", sum)

	// Reduce: 文字列連結
	words := []string{"Go", "is", "awesome"}
	sentence := Reduce(words, "", func(acc, word string) string {
		if acc == "" {
			return word
		}
		return acc + " " + word
	})
	fmt.Printf("Sentence: %s\n", sentence)

	// 型推論の例3: 明示的な型指定が必要な場合
	fmt.Println("\n=== Explicit Type Specification ===")

	// 空のスライスから始める場合は型推論できない
	var emptySlice []int
	doubled := Map[int, int](emptySlice, func(n int) int {
		return n * 2
	})
	fmt.Printf("Doubled empty slice: %v\n", doubled)

	// 複雑な型の場合
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Carol", 28},
	}

	// Person から string への変換
	names := Map(people, func(p Person) string {
		return p.Name
	})
	fmt.Printf("Names: %v\n", names)

	// Person から int への変換
	ages := Map(people, func(p Person) int {
		return p.Age
	})
	fmt.Printf("Ages: %v\n", ages)
}
