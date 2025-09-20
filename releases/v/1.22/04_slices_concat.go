// Go 1.22 新機能: slices package enhancements
// 原文: "Enhanced slices package with new utility functions"
//
// 説明: Go 1.22では、slicesパッケージに多くの便利な関数が追加され、
// スライス操作がより効率的になりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("=== slices Package Enhancements Demo ===")

	// 基本的な操作
	demonstrateBasicOperations()

	// 検索と比較
	demonstrateSearchAndCompare()

	// ソートと操作
	demonstrateSortAndManipulate()
}

func demonstrateBasicOperations() {
	fmt.Println("\n--- 基本的な操作 ---")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("元のスライス: %v\n", numbers)

	// Contains - 要素の存在確認
	fmt.Printf("3が含まれているか: %v\n", slices.Contains(numbers, 3))
	fmt.Printf("10が含まれているか: %v\n", slices.Contains(numbers, 10))

	// Index - 要素のインデックス検索
	index := slices.Index(numbers, 4)
	fmt.Printf("4のインデックス: %d\n", index)

	notFoundIndex := slices.Index(numbers, 10)
	fmt.Printf("10のインデックス: %d（見つからない場合は-1）\n", notFoundIndex)

	// Max と Min
	fmt.Printf("最大値: %d\n", slices.Max(numbers))
	fmt.Printf("最小値: %d\n", slices.Min(numbers))
}

func demonstrateSearchAndCompare() {
	fmt.Println("\n--- 検索と比較 ---")

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}

	// Equal - スライスの比較
	fmt.Printf("slice1とslice2は等しい: %v\n", slices.Equal(slice1, slice2))
	fmt.Printf("slice1とslice3は等しい: %v\n", slices.Equal(slice1, slice3))

	// Compare - スライスの辞書順比較
	result := slices.Compare(slice1, slice3)
	fmt.Printf("slice1とslice3の比較結果: %d\n", result)
	if result < 0 {
		fmt.Println("  slice1 < slice3")
	}

	// 文字列での例
	words1 := []string{"apple", "banana", "cherry"}
	words2 := []string{"apple", "banana", "date"}

	fmt.Printf("文字列スライスの比較: %d\n", slices.Compare(words1, words2))
}

func demonstrateSortAndManipulate() {
	fmt.Println("\n--- ソートと操作 ---")

	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("ソート前: %v\n", numbers)

	// Sort - インプレースソート
	slices.Sort(numbers)
	fmt.Printf("ソート後: %v\n", numbers)

	// Reverse - 逆順
	slices.Reverse(numbers)
	fmt.Printf("逆順: %v\n", numbers)

	// IsSorted - ソート済みかチェック
	fmt.Printf("ソート済みか: %v\n", slices.IsSorted(numbers))

	// もう一度ソート
	slices.Sort(numbers)
	fmt.Printf("再ソート後: %v\n", numbers)
	fmt.Printf("ソート済みか: %v\n", slices.IsSorted(numbers))

	// Clone - スライスの複製
	cloned := slices.Clone(numbers)
	fmt.Printf("複製されたスライス: %v\n", cloned)

	// 元のスライスを変更
	numbers[0] = 999
	fmt.Printf("元のスライス変更後: %v\n", numbers)
	fmt.Printf("複製は影響なし: %v\n", cloned)
}

// % go run 04_slices_concat.go
// === slices Package Enhancements Demo ===
//
// --- 基本的な操作 ---
// 元のスライス: [1 2 3 4 5]
// 3が含まれているか: true
// 10が含まれているか: false
// 4のインデックス: 3
// 10のインデックス: -1（見つからない場合は-1）
// 最大値: 5
// 最小値: 1