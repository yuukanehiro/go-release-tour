// 説明: Type Constraints - ジェネリクスで型を制約する方法を学ぶ

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

// Ordered は順序付け可能な型を定義する制約
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Numeric は数値型のみの制約
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Max は順序付け可能な型の最大値を返す
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Add は数値型の加算を行う
func Add[T Numeric](a, b T) T {
	return a + b
}

// Comparable は比較可能な型（Goの組み込み制約）
func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// 整数での使用
	fmt.Println("Max(10, 20):", Max(10, 20))
	fmt.Println("Max(-5, 3):", Max(-5, 3))

	// 浮動小数点数での使用
	fmt.Println("Max(3.14, 2.71):", Max(3.14, 2.71))

	// 文字列での使用（辞書順）
	fmt.Println("Max(\"apple\", \"banana\"):", Max("apple", "banana"))

	// 数値の加算
	fmt.Println("Add(10, 20):", Add(10, 20))
	fmt.Println("Add(3.5, 2.5):", Add(3.5, 2.5))

	// 比較操作
	fmt.Println("Equal(10, 10):", Equal(10, 10))
	fmt.Println("Equal(\"hello\", \"world\"):", Equal("hello", "world"))

	// カスタム型でも動作する（型の基底型が制約に含まれているため）
	type MyInt int
	var x, y MyInt = 100, 200
	fmt.Println("Max with custom type:", Max(x, y))
}
