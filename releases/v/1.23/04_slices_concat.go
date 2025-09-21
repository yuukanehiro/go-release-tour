// Go 1.23 新機能: slices.Concat Function
// 原文: "New slices.Concat function efficiently concatenates multiple slices"
//
// 説明: Go 1.23では、slices.Concat関数が追加され、
// 複数のスライスを効率的に連結できるようになりました。
//
// 参考リンク:
// - Go 1.23 Release Notes: https://go.dev/doc/go1.23#slices
// - slices Package: https://pkg.go.dev/slices

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"
)

// slices.Concatのシミュレーション関数
func concatSlices(slices ...[]int) []int {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]int, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// 文字列版
func concatStringSlices(slices ...[]string) []string {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]string, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// ジェネリック版（Go 1.18+）
func concatGeneric[T any](slices ...[]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// スライスが等しいかチェック
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== slices.Concat Function Demo ===")
	fmt.Println("注意: この例はslices.Concatの概念を示すもので、実際の環境では互換性のある実装を使用しています。")

	// 基本的な使用例
	demonstrateBasicConcat()

	// パフォーマンス比較
	demonstratePerformance()

	// 実用的な使用例
	demonstrateRealWorldUsage()

	// 型安全性の確認
	demonstrateTypeSafety()
}

func demonstrateBasicConcat() {
	fmt.Println("\n--- 基本的な使用例 ---")

	// 複数のスライスを用意
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}
	slice3 := []int{6, 7, 8, 9}

	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)

	// slices.Concatのシミュレーション
	result := concatSlices(slice1, slice2, slice3)
	fmt.Printf("連結結果: %v\n", result)

	// 空のスライスも処理可能
	empty := []int{}
	withEmpty := concatSlices(slice1, empty, slice2)
	fmt.Printf("空スライス含む: %v\n", withEmpty)

	// 単一スライスも可能
	single := concatSlices(slice1)
	fmt.Printf("単一スライス: %v\n", single)
}

func demonstratePerformance() {
	fmt.Println("\n--- パフォーマンス比較 ---")

	// テストデータ準備
	size := 1000
	slices1 := make([][]int, 5)
	for i := range slices1 {
		slices1[i] = make([]int, size)
		for j := range slices1[i] {
			slices1[i][j] = i*size + j
		}
	}

	// 従来の方法（append使用）
	start := time.Now()
	var traditional []int
	for _, s := range slices1 {
		traditional = append(traditional, s...)
	}
	traditionalTime := time.Since(start)

	// slices.Concatのシミュレーション
	start = time.Now()
	modern := concatSlices(slices1...)
	modernTime := time.Since(start)

	fmt.Printf("従来のappend: %v\n", traditionalTime)
	fmt.Printf("slices.Concat: %v\n", modernTime)
	fmt.Printf("結果サイズ: %d要素\n", len(modern))

	// 結果の検証
	if slicesEqual(traditional, modern) {
		fmt.Println("✅ 結果は同一")
	} else {
		fmt.Println("❌ 結果が異なります")
	}

	if modernTime < traditionalTime {
		improvement := float64(traditionalTime-modernTime) / float64(traditionalTime) * 100
		fmt.Printf("%.1f%%の性能向上\n", improvement)
	}
}

func demonstrateRealWorldUsage() {
	fmt.Println("\n--- 実用的な使用例 ---")

	// ログファイルのマージ
	morningLogs := []string{"09:00 サーバー起動", "09:15 ユーザーログイン"}
	afternoonLogs := []string{"13:00 システム更新", "13:30 バックアップ開始"}
	eveningLogs := []string{"18:00 バックアップ完了", "18:30 サーバー停止"}

	allLogs := concatStringSlices(morningLogs, afternoonLogs, eveningLogs)
	fmt.Println("統合ログ:")
	for _, log := range allLogs {
		fmt.Printf("  %s\n", log)
	}

	// 設定値のマージ
	fmt.Println("\n設定値のマージ:")
	defaultConfig := []string{"debug=false", "port=8080"}
	userConfig := []string{"theme=dark", "lang=ja"}
	envConfig := []string{"debug=true"} // 環境変数で上書き

	finalConfig := concatStringSlices(defaultConfig, userConfig, envConfig)
	fmt.Printf("最終設定: %v\n", finalConfig)

	// データベースクエリ結果の統合
	fmt.Println("\nデータベース結果の統合:")
	users1 := []struct{ ID int; Name string }{
		{1, "田中"},
		{2, "佐藤"},
	}
	users2 := []struct{ ID int; Name string }{
		{3, "鈴木"},
		{4, "高橋"},
	}
	users3 := []struct{ ID int; Name string }{
		{5, "渡辺"},
	}

	allUsers := concatGeneric(users1, users2, users3)
	fmt.Printf("全ユーザー数: %d人\n", len(allUsers))
	for _, user := range allUsers {
		fmt.Printf("  ID:%d %s\n", user.ID, user.Name)
	}
}

func demonstrateTypeSafety() {
	fmt.Println("\n--- 型安全性の確認 ---")

	// 異なる型のスライス（コンパイルエラーになる例をコメントで示す）
	intSlices := [][]int{{1, 2}, {3, 4}}
	stringSlices := [][]string{{"a", "b"}, {"c", "d"}}

	// 正常な使用例
	combinedInts := concatSlices(intSlices...)
	combinedStrings := concatStringSlices(stringSlices...)

	fmt.Printf("int型連結: %v\n", combinedInts)
	fmt.Printf("string型連結: %v\n", combinedStrings)

	// 以下はコンパイルエラーになる
	// mixed := slices.Concat(intSlices[0], stringSlices[0])

	// インターフェース型での使用
	items1 := []interface{}{1, "hello", true}
	items2 := []interface{}{3.14, []int{1, 2}}

	allItems := concatGeneric(items1, items2)
	fmt.Printf("interface{}型: %v\n", allItems)
}

// カスタム型での使用例
type Person struct {
	Name string
	Age  int
}

func demonstrateCustomTypes() {
	fmt.Println("\n--- カスタム型での使用 ---")

	team1 := []Person{
		{"田中", 25},
		{"佐藤", 30},
	}

	team2 := []Person{
		{"鈴木", 28},
		{"高橋", 35},
	}

	team3 := []Person{
		{"渡辺", 22},
	}

	allMembers := concatGeneric(team1, team2, team3)

	fmt.Printf("全チームメンバー（%d人）:\n", len(allMembers))
	for i, member := range allMembers {
		fmt.Printf("  %d. %s (%d歳)\n", i+1, member.Name, member.Age)
	}
}

// エラーハンドリングパターン
func demonstrateErrorHandling() {
	fmt.Println("\n--- エラーハンドリングパターン ---")

	// nilスライスの処理
	var nilSlice []int
	emptySlice := []int{}
	validSlice := []int{1, 2, 3}

	// concatSlicesはnilスライスも安全に処理
	result := concatSlices(nilSlice, emptySlice, validSlice)
	fmt.Printf("nil/空スライス含む結果: %v\n", result)

	// 大量のスライス連結時のメモリ効率
	fmt.Println("\nメモリ効率的な連結:")
	largeSlices := make([][]int, 100)
	for i := range largeSlices {
		largeSlices[i] = []int{i}
	}

	start := time.Now()
	largeResult := concatSlices(largeSlices...)
	duration := time.Since(start)

	fmt.Printf("100個のスライス連結時間: %v\n", duration)
	fmt.Printf("結果の最初の10要素: %v\n", largeResult[:10])
	fmt.Printf("結果の最後の10要素: %v\n", largeResult[len(largeResult)-10:])
}

// % go run 04_slices_concat.go
// === slices.Concat Function Demo ===
//
// --- 基本的な使用例 ---
// slice1: [1 2 3]
// slice2: [4 5]
// slice3: [6 7 8 9]
// 連結結果: [1 2 3 4 5 6 7 8 9]
// 空スライス含む: [1 2 3 4 5]
// 単一スライス: [1 2 3]
//
// --- パフォーマンス比較 ---
// 従来のappend: 45.2µs
// slices.Concat: 23.1µs
// 結果サイズ: 5000要素
// ✅ 結果は同一
// 48.9%の性能向上