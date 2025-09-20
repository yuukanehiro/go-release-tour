// Go 1.22 新機能: math/rand/v2 Package
// 原文: "New math/rand/v2 package with improved API and better random number generation"
//
// 説明: Go 1.22では、改良されたAPIとより良い乱数生成アルゴリズムを持つ
// math/rand/v2パッケージが追加されました。
//
// 参考リンク:
// - Go 1.22 Release Notes: https://go.dev/doc/go1.22#math-rand-v2
// - math/rand/v2 Package: https://pkg.go.dev/math/rand/v2

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== math/rand/v2 Package Demo ===")

	// 注意: Go 1.22の新機能ですが、現在の環境では動作しません
	// 以下は概念を示すための疑似コードです

	// 基本的な使用例
	demonstrateBasicUsage()

	// 型安全性の改善
	demonstrateTypeSafety()

	// パフォーマンスの改善
	demonstratePerformance()

	// 従来との比較
	demonstrateComparison()
}

func demonstrateBasicUsage() {
	fmt.Println("\n--- 基本的な使用例 ---")

	// 現在のGoバージョンでの従来の方法
	oldRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("従来のmath/rand:")
	for i := 0; i < 5; i++ {
		value := oldRand.Intn(100)
		fmt.Printf("  乱数 %d: %d\n", i+1, value)
	}

	// Go 1.22のmath/rand/v2概念例:
	fmt.Println("\nGo 1.22のmath/rand/v2（概念例）:")
	fmt.Println("  import \"math/rand/v2\"")
	fmt.Println("  r := rand.New(rand.NewPCG(seed))")
	fmt.Println("  value := r.IntN(100)  // より明確なAPI")

	// シミュレーション
	fmt.Println("\n改善されたAPI（シミュレーション）:")
	for i := 0; i < 5; i++ {
		value := oldRand.Intn(100) // 実際のv2では IntN() メソッド
		fmt.Printf("  乱数 %d: %d（より高品質な乱数）\n", i+1, value)
	}
}

func demonstrateTypeSafety() {
	fmt.Println("\n--- 型安全性の改善 ---")

	// 従来の方法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("従来のAPIの問題:")
	fmt.Println("  Intn(0) → パニック！")
	fmt.Println("  Intn(-1) → パニック！")

	fmt.Println("\nGo 1.22のrand/v2の改善:")
	fmt.Println("  IntN(n uint) → 負数でコンパイルエラー")
	fmt.Println("  UintN(n uint) → 符号なし整数用")
	fmt.Println("  Int32N(n uint32) → 32bit専用")
	fmt.Println("  Int64N(n uint64) → 64bit専用")

	// より安全な使用例（シミュレーション）
	fmt.Println("\n型安全な乱数生成（シミュレーション）:")

	// 0-99の範囲
	value := r.Intn(100)
	fmt.Printf("  0-99の乱数: %d\n", value)

	// より大きな範囲（シミュレーション）
	largeValue := r.Int63n(1000000)
	fmt.Printf("  大きな乱数: %d\n", largeValue)

	// 浮動小数点
	floatValue := r.Float64()
	fmt.Printf("  0.0-1.0の乱数: %.6f\n", floatValue)
}

func demonstratePerformance() {
	fmt.Println("\n--- パフォーマンスの改善 ---")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// パフォーマンステスト（シミュレーション）
	iterations := 100000

	start := time.Now()
	sum := 0
	for i := 0; i < iterations; i++ {
		sum += r.Intn(1000)
	}
	duration := time.Since(start)

	fmt.Printf("従来のmath/rand: %d回の乱数生成に%v\n", iterations, duration)
	fmt.Printf("合計値: %d\n", sum)

	fmt.Println("\nGo 1.22のrand/v2の改善点:")
	fmt.Println("  - PCG（Permuted Congruential Generator）採用")
	fmt.Println("  - より高品質な乱数")
	fmt.Println("  - より高速な生成")
	fmt.Println("  - より良い統計的性質")

	// アルゴリズムの比較例
	fmt.Println("\n乱数生成アルゴリズム:")
	fmt.Println("  従来: Linear Congruential Generator (LCG)")
	fmt.Println("  v2: Permuted Congruential Generator (PCG)")
	fmt.Println("  PCGの利点:")
	fmt.Println("    - より長い周期")
	fmt.Println("    - より良い分布")
	fmt.Println("    - 予測しにくい")
	fmt.Println("    - より高速")
}

func demonstrateComparison() {
	fmt.Println("\n--- 従来との比較 ---")

	fmt.Println("Go 1.21以前（math/rand）:")
	fmt.Println("  import \"math/rand\"")
	fmt.Println("  r := rand.New(rand.NewSource(seed))")
	fmt.Println("  value := r.Intn(100)  // int, パニックの可能性")
	fmt.Println("  問題: 型安全性の欠如、品質の限界")

	fmt.Println("\nGo 1.22（math/rand/v2）:")
	fmt.Println("  import \"math/rand/v2\"")
	fmt.Println("  r := rand.New(rand.NewPCG(seed))")
	fmt.Println("  value := r.IntN(100)  // より明確、型安全")
	fmt.Println("  改善: より良いAPI、高品質な乱数")

	// 実用的な使用例の比較
	fmt.Println("\n実用例の比較:")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// パスワード生成例
	fmt.Println("\nパスワード生成（従来の方法）:")
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, 8)
	for i := range password {
		password[i] = chars[r.Intn(len(chars))]
	}
	fmt.Printf("  生成されたパスワード: %s\n", password)

	fmt.Println("\nGo 1.22のrand/v2では:")
	fmt.Println("  password[i] = chars[r.IntN(uint(len(chars)))]")
	fmt.Println("  利点: 型安全、より高品質な乱数")

	// 配列のシャッフル例
	fmt.Println("\n配列のシャッフル:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("  元の配列: %v\n", numbers)

	// Fisher-Yates シャッフル
	for i := len(numbers) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	fmt.Printf("  シャッフル後: %v\n", numbers)

	fmt.Println("\nGo 1.22のrand/v2では:")
	fmt.Println("  j := r.IntN(uint(i + 1))  // より安全")
}

// セキュリティに関する注意
func demonstrateSecurityNote() {
	fmt.Println("\n--- セキュリティに関する注意 ---")

	fmt.Println("math/rand/v2の注意点:")
	fmt.Println("  - 暗号学的に安全ではない")
	fmt.Println("  - セキュリティ用途にはcrypto/randを使用")
	fmt.Println("  - ゲーム、シミュレーション用途に最適")

	fmt.Println("\n用途別の選択:")
	fmt.Println("  ゲーム・シミュレーション → math/rand/v2")
	fmt.Println("  暗号・セキュリティ → crypto/rand")
	fmt.Println("  テスト・ベンチマーク → math/rand/v2")
}

// % go run 03_math_rand_v2.go
// === math/rand/v2 Package Demo ===
//
// --- 基本的な使用例 ---
// 従来のmath/rand:
//   乱数 1: 81
//   乱数 2: 47
//   乱数 3: 25
//   乱数 4: 56
//   乱数 5: 94
//
// Go 1.22のmath/rand/v2（概念例）:
//   import "math/rand/v2"
//   r := rand.New(rand.NewPCG(seed))
//   value := r.IntN(100)  // より明確なAPI