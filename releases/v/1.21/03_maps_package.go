// Go 1.21 新機能: maps package
// 原文: "New maps package for common map operations"
//
// 説明: Go 1.21では、マップ操作のためのmapsパッケージが標準ライブラリに追加されました。
//
// 参考リンク:
// - Go 1.21 Release Notes: https://go.dev/doc/go1.21#maps
// - maps Package: https://pkg.go.dev/maps

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("=== maps Package Demo ===")

	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("元のマップ: %v\n", m1)

	// コピー
	m2 := maps.Clone(m1)
	fmt.Printf("複製: %v\n", m2)

	// 比較
	fmt.Printf("等しい: %v\n", maps.Equal(m1, m2))

	// キーの取得
	keys := maps.Keys(m1)
	fmt.Printf("キー: %v\n", keys)

	// 値の取得
	values := maps.Values(m1)
	fmt.Printf("値: %v\n", values)

	// マージ
	m3 := map[string]int{"d": 4, "e": 5}
	maps.Copy(m2, m3)
	fmt.Printf("マージ後: %v\n", m2)

	fmt.Println("\n利点:")
	fmt.Println("  - 一般的なマップ操作の標準化")
	fmt.Println("  - メモリ安全な実装")
	fmt.Println("  - ジェネリクスによる型安全性")
}

// % go run 03_maps_package.go
// === maps Package Demo ===
// 元のマップ: map[a:1 b:2 c:3]
// 複製: map[a:1 b:2 c:3]
// 等しい: true