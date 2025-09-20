//go:build ignore
// +build ignore

// 説明: Generic Data Structures - ジェネリクスを使ったデータ構造の実装
package main

import (
	"fmt"
)

// Stack はジェネリックなスタックの実装
type Stack[T any] struct {
	items []T
}

// NewStack は新しいスタックを作成
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push は要素をスタックに追加
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop はスタックから要素を取り出し
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

// IsEmpty はスタックが空かどうかを確認
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size はスタックのサイズを返す
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Pair はジェネリックなペア構造体
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair は新しいペアを作成
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		First:  first,
		Second: second,
	}
}

// Swap はペアの要素を入れ替える
func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{
		First:  p.Second,
		Second: p.First,
	}
}

func main() {
	// 整数スタックの使用例
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Integer Stack:")
	for !intStack.IsEmpty() {
		if value, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d, Stack size: %d\n", value, intStack.Size())
		}
	}

	// 文字列スタックの使用例
	strStack := NewStack[string]()
	strStack.Push("Hello")
	strStack.Push("World")
	strStack.Push("Generics")

	fmt.Println("\nString Stack:")
	for !strStack.IsEmpty() {
		if value, ok := strStack.Pop(); ok {
			fmt.Printf("Popped: %s, Stack size: %d\n", value, strStack.Size())
		}
	}

	// ペアの使用例
	fmt.Println("\nPair Examples:")

	// 異なる型のペア
	namePair := NewPair("Alice", 25)
	fmt.Printf("Name-Age pair: %+v\n", namePair)

	// ペアの入れ替え
	swapped := namePair.Swap()
	fmt.Printf("Swapped pair: %+v\n", swapped)

	// 同じ型のペア
	coordPair := NewPair(10.5, 20.3)
	fmt.Printf("Coordinate pair: %+v\n", coordPair)

	// 構造体のペア
	type Person struct {
		Name string
		Age  int
	}

	person1 := Person{"Bob", 30}
	person2 := Person{"Carol", 28}
	peoplePair := NewPair(person1, person2)
	fmt.Printf("People pair: %+v\n", peoplePair)
}
