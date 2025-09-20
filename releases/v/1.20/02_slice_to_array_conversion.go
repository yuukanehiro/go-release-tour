// Go 1.20 新機能: Slice to Array Conversion
// 原文: "Go 1.20 extends slice-to-array-pointer conversions"
//
// 説明: Go 1.20では、スライスから配列への直接変換が可能になり、
// より安全で読みやすいコードが書けるようになりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Slice to Array Conversion Demo ===")

	// 基本的なスライス
	slice := []byte{1, 2, 3, 4, 5, 6}
	fmt.Printf("元のスライス: %v (長さ: %d)\n", slice, len(slice))

	fmt.Println("\n--- Go 1.20の新機能 ---")

	// 1. スライスから配列への直接変換
	array4 := [4]byte(slice) // 最初の4要素を配列に変換
	fmt.Printf("配列変換 [4]byte(slice): %v\n", array4)

	// 2. 異なるサイズの配列変換
	array2 := [2]byte(slice) // 最初の2要素
	fmt.Printf("配列変換 [2]byte(slice): %v\n", array2)

	// 3. 文字列での例
	str := "Hello, World!"
	strSlice := []byte(str)
	fmt.Printf("\n文字列からバイトスライス: %v\n", strSlice)

	// 文字列の最初の5文字を配列に
	hello := [5]byte(strSlice)
	fmt.Printf("最初の5文字を配列に: %v ('%s')\n", hello, string(hello[:]))

	fmt.Println("\n--- 従来の方法との比較 ---")

	// 従来の方法（Go 1.19以前）
	fmt.Println("従来の方法:")
	oldWay := *(*[4]byte)(slice) // ポインタ変換が必要
	fmt.Printf("  *(*[4]byte)(slice): %v\n", oldWay)

	fmt.Println("Go 1.20の方法:")
	newWay := [4]byte(slice) // 直接変換
	fmt.Printf("  [4]byte(slice): %v\n", newWay)

	fmt.Println("\n--- 実用例 ---")

	// 1. バイト配列の処理
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello"
	fmt.Printf("データ: %v\n", data)

	// ヘッダーとして最初の4バイトを取得
	header := [4]byte(data)
	fmt.Printf("ヘッダー（最初の4バイト）: %v\n", header)

	// 2. ハッシュ値の処理
	hashSlice := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	hash := [8]byte(hashSlice)
	fmt.Printf("ハッシュ値: %x\n", hash)

	// 3. ネットワークアドレスの処理
	ipBytes := []byte{192, 168, 1, 1}
	ipArray := [4]byte(ipBytes)
	fmt.Printf("IPアドレス: %d.%d.%d.%d\n", ipArray[0], ipArray[1], ipArray[2], ipArray[3])

	fmt.Println("\n--- エラーハンドリング ---")

	// スライスが短すぎる場合の動作
	shortSlice := []byte{1, 2}
	fmt.Printf("短いスライス: %v (長さ: %d)\n", shortSlice, len(shortSlice))

	// 実際の使用では以下のようなチェックが必要
	if len(shortSlice) >= 4 {
		safe := [4]byte(shortSlice)
		fmt.Printf("安全な変換: %v\n", safe)
	} else {
		fmt.Printf("スライスが短すぎます（長さ %d < 4）\n", len(shortSlice))
	}

	fmt.Println("\n--- パフォーマンスの比較 ---")

	// 変換のパフォーマンステスト用データ
	largeSlice := make([]byte, 1000)
	for i := range largeSlice {
		largeSlice[i] = byte(i % 256)
	}

	fmt.Println("大きなスライスでの変換:")
	fmt.Printf("  スライス長: %d\n", len(largeSlice))

	// 新しい方法（Go 1.20）
	chunk := [16]byte(largeSlice)
	fmt.Printf("  最初の16バイト: %v...\n", chunk[:8])

	fmt.Println("\n利点:")
	fmt.Println("  - より読みやすい構文")
	fmt.Println("  - 型安全性の向上")
	fmt.Println("  - ポインタ操作が不要")
	fmt.Println("  - コンパイル時の型チェック")
	fmt.Println("  - パフォーマンスの向上")

	fmt.Println("\n注意点:")
	fmt.Println("  - スライスの長さが配列のサイズ以上である必要")
	fmt.Println("  - 実行時パニックを避けるため事前チェック推奨")
	fmt.Println("  - 元のスライスのコピーが作成される（参照ではない）")
}

// 実用的な関数例
func parseHeader(data []byte) ([8]byte, error) {
	if len(data) < 8 {
		return [8]byte{}, fmt.Errorf("データが短すぎます: %d < 8", len(data))
	}
	return [8]byte(data), nil
}

func ipFromBytes(data []byte) ([4]byte, bool) {
	if len(data) < 4 {
		return [4]byte{}, false
	}
	return [4]byte(data), true
}

// ベンチマーク用の例
func processChunks(data []byte) [][16]byte {
	var chunks [][16]byte
	for i := 0; i+16 <= len(data); i += 16 {
		chunk := [16]byte(data[i:])
		chunks = append(chunks, chunk)
	}
	return chunks
}

// % go run 02_slice_to_array_conversion.go
// === Slice to Array Conversion Demo ===
// 元のスライス: [1 2 3 4 5 6] (長さ: 6)
//
// --- Go 1.20の新機能 ---
// 配列変換 [4]byte(slice): [1 2 3 4]
// 配列変換 [2]byte(slice): [1 2]
//
// 文字列からバイトスライス: [72 101 108 108 111 44 32 87 111 114 108 100 33]
// 最初の5文字を配列に: [72 101 108 108 111] ('Hello')
//
// --- 従来の方法との比較 ---
// 従来の方法:
//   *(*[4]byte)(slice): [1 2 3 4]
// Go 1.20の方法:
//   [4]byte(slice): [1 2 3 4]