// Go 1.23 新機能: Range over Function Types (Iterators)
// 原文: "Range over function types enables custom iteration patterns with for-range loops"
//
// 説明: Go 1.23では、関数型に対するrange構文が追加され、
// カスタムイテレーターの実装が可能になりました。
//
// 参考リンク:
// - Go 1.23 Release Notes: https://go.dev/doc/go1.23#iterators
// - iter Package: https://pkg.go.dev/iter
//
// 注意: この例はイテレーターの概念を示すもので、実際の環境では従来の反復処理を使用しています。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sort"
)

// イテレーター関数型の概念例
type Iterator[T any] struct {
	values []T
	index  int
}

func NewIterator[T any](values []T) *Iterator[T] {
	return &Iterator[T]{values: values, index: 0}
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < len(it.values)
}

func (it *Iterator[T]) Next() T {
	if !it.HasNext() {
		var zero T
		return zero
	}
	value := it.values[it.index]
	it.index++
	return value
}

func main() {
	fmt.Println("=== Range over Function Types Demo ===")
	fmt.Println("注意: これはイテレーターの概念例です。実際の環境では従来の反復処理を使用します。")

	// 基本的なカスタムイテレーター（シミュレーション）
	numbers := func(callback func(int)) {
		for i := 1; i <= 5; i++ {
			callback(i)
		}
	}

	fmt.Println("\nカスタムイテレーターの使用:")
	numbers(func(n int) {
		fmt.Printf("数値: %d\n", n)
	})

	// キー・バリューペアのイテレーター（シミュレーション）
	fmt.Println("\n--- キー・バリューイテレーター ---")
	keyValue := func(callback func(string, int)) {
		pairs := map[string]int{
			"apple":  100,
			"banana": 200,
			"orange": 150,
		}
		for k, v := range pairs {
			callback(k, v)
		}
	}

	keyValue(func(key string, value int) {
		fmt.Printf("%s: %d円\n", key, value)
	})

	// フィルタリングイテレーター（シミュレーション）
	fmt.Println("\n--- フィルタリングイテレーター ---")
	evenNumbers := func(callback func(int)) {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				callback(i)
			}
		}
	}

	evenNumbers(func(num int) {
		fmt.Printf("偶数: %d\n", num)
	})

	// 無限イテレーター（制限付き）シミュレーション
	fmt.Println("\n--- 無限イテレーター（制限付き） ---")
	fibonacci := func(limit int, callback func(int)) {
		a, b := 0, 1
		for i := 0; i < limit; i++ {
			callback(a)
			a, b = b, a+b
		}
	}

	fibonacci(8, func(fib int) {
		fmt.Printf("フィボナッチ数: %d\n", fib)
	})

	// 遅延評価イテレーター（シミュレーション）
	fmt.Println("\n--- 遅延評価イテレーター ---")
	lazyRange := func(start, end int, callback func(int) bool) {
		fmt.Printf("  遅延評価開始: %d から %d まで\n", start, end)
		for i := start; i < end; i++ {
			fmt.Printf("  値を生成: %d\n", i)
			if !callback(i) {
				fmt.Println("  イテレーター中断")
				return
			}
		}
		fmt.Println("  遅延評価完了")
	}

	lazyRange(10, 13, func(val int) bool {
		fmt.Printf("受信した値: %d\n", val)
		if val == 11 {
			fmt.Println("  早期終了")
			return false // 中断
		}
		return true // 継続
	})

	// エラーハンドリング付きイテレーター（シミュレーション）
	fmt.Println("\n--- エラーハンドリング付きイテレーター ---")
	safeRange := func(data []int, callback func(int)) {
		if len(data) == 0 {
			fmt.Println("警告: 空のスライスです")
			return
		}
		for _, item := range data {
			callback(item)
		}
	}

	// 空のスライスでテスト
	safeRange([]int{}, func(val int) {
		fmt.Printf("値: %d\n", val)
	})

	// 正常なスライスでテスト
	safeRange([]int{100, 200, 300}, func(val int) {
		fmt.Printf("値: %d\n", val)
	})

	// チェーン可能なイテレーター（シミュレーション）
	fmt.Println("\n--- チェーン可能なイテレーター ---")

	// 基本データ
	baseData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map変換 → Filter → 結果表示
	var results []int
	for _, val := range baseData {
		doubled := val * 2  // Map: 2倍
		if doubled >= 10 {  // Filter: 10以上のみ
			results = append(results, doubled)
		}
	}

	for _, val := range results {
		fmt.Printf("チェーン結果: %d\n", val)
	}

	fmt.Println("\n--- 従来の方法との比較 ---")
	fmt.Println("従来: スライスを事前に作成して反復")
	fmt.Println("新機能: オンデマンドで値を生成するイテレーター")
	fmt.Println("利点:")
	fmt.Println("  - メモリ効率（大量データの遅延評価）")
	fmt.Println("  - 無限シーケンスの表現")
	fmt.Println("  - 早期終了による処理中断")
	fmt.Println("  - 関数型プログラミングパターン")
}

// スライス処理との連携例
func demonstrateSlicesIntegration() {
	fmt.Println("\n--- スライス処理との連携 ---")

	data := []int{1, 3, 2, 5, 4}
	fmt.Printf("元データ: %v\n", data)

	// ソート後の反復
	sort.Ints(data)
	fmt.Printf("ソート後: %v\n", data)

	// カスタム処理
	fmt.Println("処理結果:")
	for _, val := range data {
		processed := fmt.Sprintf("値_%d", val*10)
		fmt.Println(processed)
	}
}

// % go run 02_iterators.go
// === Range over Function Types Demo ===
// カスタムイテレーターの使用:
// 数値: 1
// 数値: 2
// 数値: 3
// 数値: 4
// 数値: 5
//
// --- キー・バリューイテレーター ---
// apple: 100円
// banana: 200円
// orange: 150円
//
// --- フィルタリングイテレーター ---
// 偶数: 2
// 偶数: 4
// 偶数: 6
// 偶数: 8
// 偶数: 10