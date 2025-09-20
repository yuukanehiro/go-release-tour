// Go 1.24 新機能: Generic Type Aliases
// 原文: "Go 1.24 now fully supports generic type aliases: a type alias may be parameterized like a defined type"
//
// 説明: Go 1.24では、型エイリアスがジェネリクス対応となり、パラメータ化された型エイリアスが作成できるようになりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

// 注意: これらはGo 1.24の新機能です。現在のGoバージョンでは動作しません。
// 以下は概念を示すための疑似コードです。

// Go 1.24での Generic type alias の例:
// type List[T any] = []T
// type Dict[K comparable, V any] = map[K]V
// type Transformer[T, U any] = func(T) U
// type Numeric[T ~int | ~float64] = T

// 現在のGoバージョンで動作する通常の型定義:
type IntList = []int
type StringList = []string
type StringIntMap = map[string]int
type IntToStringFunc = func(int) string
type StringToIntFunc = func(string) int

// 通常のジェネリック型（Go 1.18+）
type GenericList[T any] []T
type GenericDict[K comparable, V any] map[K]V
type GenericTransformer[T, U any] func(T) U

// 結果構造体
type Result[T any] struct {
	Value T
	Error error
}

func main() {
	fmt.Println("=== Generic Type Aliases Demo ===")

	// 現在のGoバージョンでの型エイリアス使用例
	numbers := IntList{1, 2, 3, 4, 5}
	strings := StringList{"hello", "world", "go", "1.24"}
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Strings: %v\n", strings)

	// 通常のマップ
	userAge := StringIntMap{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}
	fmt.Printf("User ages: %v\n", userAge)

	// 関数型エイリアスの使用
	var intToString IntToStringFunc = func(i int) string {
		return fmt.Sprintf("Number: %d", i)
	}

	var stringToInt StringToIntFunc = func(s string) int {
		return len(s)
	}

	fmt.Printf("Transform 42: %s\n", intToString(42))
	fmt.Printf("Transform 'hello': %d\n", stringToInt("hello"))

	// ジェネリック型の使用例
	genericNumbers := GenericList[int]{100, 200, 300}
	genericStrings := GenericList[string]{"generic", "type", "example"}
	fmt.Printf("Generic numbers: %v\n", genericNumbers)
	fmt.Printf("Generic strings: %v\n", genericStrings)

	// Using Result type alias
	successResult := Result[string]{
		Value: "Operation completed",
		Error: nil,
	}

	errorResult := Result[int]{
		Value: 0,
		Error: fmt.Errorf("something went wrong"),
	}

	fmt.Printf("Success result: %+v\n", successResult)
	fmt.Printf("Error result: %+v\n", errorResult)

	// 型エイリアスの互換性を実証
	fmt.Println("\n--- Type Alias Compatibility ---")

	// 型エイリアスは基底型と同一
	var regularSlice []int = []int{1, 2, 3}
	var aliasSlice IntList = IntList{4, 5, 6}

	// 相互に代入可能
	regularSlice = aliasSlice
	aliasSlice = regularSlice

	fmt.Printf("Regular slice: %v\n", regularSlice)
	fmt.Printf("Alias slice: %v\n", aliasSlice)

	// 両方で動作する関数
	printSlice := func(s []int) {
		fmt.Printf("Slice content: %v\n", s)
	}

	printSlice(regularSlice)
	printSlice(aliasSlice)

	// Go 1.24での Generic Type Alias の期待される動作
	fmt.Println("\n--- Go 1.24 Generic Type Alias 期待動作 ---")
	fmt.Println("Go 1.24では以下のような書き方が可能になります:")
	fmt.Println("type List[T any] = []T")
	fmt.Println("type Dict[K comparable, V any] = map[K]V")
	fmt.Println("type Transform[T, U any] = func(T) U")
}

// 現在のGoバージョンでのジェネリック関数例
func processData[T any](data GenericList[T], transform GenericTransformer[T, string]) GenericList[string] {
	result := make(GenericList[string], len(data))
	for i, item := range data {
		result[i] = transform(item)
	}
	return result
}

// Go 1.24でのGeneric Type Aliasを使った関数（期待される形）
// func processDataWithAlias[T any](data List[T], transform Transformer[T, string]) List[string] {
//     result := make(List[string], len(data))
//     for i, item := range data {
//         result[i] = transform(item)
//     }
//     return result
// }

// % go run 01_generic_type_aliases.go
// === Generic Type Aliases Demo ===
// Numbers: [1 2 3 4 5]
// Strings: [hello world go 1.24]
// User ages: map[Alice:30 Bob:25 Carol:35]
// Transform 42: Number: 42
// Transform 'hello': 5
// Int numeric: 100
// Float numeric: 3.14
// Success result: {Value:Operation completed Error:<nil>}
// Error result: {Value:0 Error:something went wrong}
//
// --- Type Alias Compatibility ---
// Regular slice: [4 5 6]
// Alias slice: [4 5 6]
// Slice content: [4 5 6]
// Slice content: [4 5 6]