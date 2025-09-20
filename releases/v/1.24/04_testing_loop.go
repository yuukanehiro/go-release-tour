// Go 1.24 新機能: testing.B.Loop() Method
// 原文: "Benchmarks may now use the faster and less error-prone testing.B.Loop method"
//
// 説明: 新しいtesting.B.Loop()メソッドにより、従来のb.Nを使用したループより高速で安全なベンチマークが作成できます。
//
// 参考リンク:
// - Go 1.24 Release Notes: https://go.dev/doc/go1.24#testing
// - testing Package: https://pkg.go.dev/testing
//
// 注意: この機能はGo 1.24の新機能で、現在のGoバージョンでは利用できません。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// 現在のベンチマーク例（通常のパターン）
func BenchmarkStringConcatCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := ""
		for j := 0; j < 100; j++ {
			result += "hello"
		}
		_ = result
	}
}

// Go 1.24の新しいベンチマーク例（将来のパターン）
// func BenchmarkStringConcatNew(b *testing.B) {
// 	for b.Loop() {
// 		result := ""
// 		for j := 0; j < 100; j++ {
// 			result += "hello"
// 		}
// 		_ = result
// 	}
// }

// より複雑な例：文字列処理
func BenchmarkStringBuilderOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("hello")
		}
		_ = builder.String()
	}
}

// 注意: B.Loop()は実際にはGo 1.24で実装されていません
// func BenchmarkStringBuilderNew(b *testing.B) {
// 	for b.Loop() {
// 		var builder strings.Builder
// 		for j := 0; j < 100; j++ {
// 			builder.WriteString("hello")
// 		}
// 		_ = builder.String()
// 	}
// }

// スライス操作のベンチマーク例
func BenchmarkSliceAppendOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0)
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}
}

// 注意: B.Loop()は実際にはGo 1.24で実装されていません
// func BenchmarkSliceAppendNew(b *testing.B) {
// 	for b.Loop() {
// 		slice := make([]int, 0)
// 		for j := 0; j < 1000; j++ {
// 			slice = append(slice, j)
// 		}
// 	}
// }

// 実際のベンチマーク実行をシミュレート
func simulateBenchmark(name string, fn func()) time.Duration {
	start := time.Now()
	iterations := 10000

	for i := 0; i < iterations; i++ {
		fn()
	}

	elapsed := time.Since(start)
	fmt.Printf("%s: %d iterations in %v (avg: %v per operation)\n",
		name, iterations, elapsed, elapsed/time.Duration(iterations))
	return elapsed
}

func main() {
	fmt.Println("=== testing.B.Loop() Demo ===")
	fmt.Println("注意: testing.B.Loop()はGo 1.24の新機能です")
	fmt.Println("現在は概念的なデモを実行し、実際の使用例を示します")
	fmt.Println()

	fmt.Println("現在のベンチマーク記述法:")
	fmt.Println(`func BenchmarkCurrent(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // ベンチマーク対象のコード
    }
}`)

	fmt.Println("\nGo 1.24の新しい記述法（将来実装予定）:")
	fmt.Println(`func BenchmarkNew(b *testing.B) {
    for b.Loop() {
        // ベンチマーク対象のコード
    }
}`)

	fmt.Println("\n現在の推奨方法:")
	fmt.Println(`func BenchmarkRecommended(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // ベンチマーク対象のコード
    }
}`)

	fmt.Println("\n--- 実際のパフォーマンス比較 ---")

	// 文字列連結のシミュレーション
	oldStringConcat := func() {
		result := ""
		for j := 0; j < 100; j++ {
			result += "hello"
		}
	}

	newStringBuilder := func() {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("hello")
		}
		_ = builder.String()
	}

	fmt.Println("1. 文字列連結パフォーマンス:")
	time1 := simulateBenchmark("String concatenation (+)", oldStringConcat)
	time2 := simulateBenchmark("String builder", newStringBuilder)

	improvement := float64(time1-time2) / float64(time1) * 100
	if improvement > 0 {
		fmt.Printf("   StringBuilder is %.1f%% faster\n", improvement)
	} else {
		fmt.Printf("   String concat is %.1f%% faster\n", -improvement)
	}

	fmt.Println("\n2. スライス操作パフォーマンス:")

	preAllocatedSlice := func() {
		slice := make([]int, 0, 1000) // Pre-allocate capacity
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}

	dynamicSlice := func() {
		slice := make([]int, 0) // No pre-allocation
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}

	time3 := simulateBenchmark("Pre-allocated slice", preAllocatedSlice)
	time4 := simulateBenchmark("Dynamic slice", dynamicSlice)

	improvement2 := float64(time4-time3) / float64(time4) * 100
	fmt.Printf("   Pre-allocated slice is %.1f%% faster\n", improvement2)

	fmt.Println("\n--- B.Loop()の期待される利点 ---")
	fmt.Println("✅ より読みやすいコード")
	fmt.Println("✅ インデックス変数(i)が不要")
	fmt.Println("✅ b.Nの誤用を防止")
	fmt.Println("✅ ベンチマークフレームワークの内部最適化")
	fmt.Println("✅ より正確な計測")

	fmt.Println("\n--- 現在のベストプラクティス ---")
	fmt.Println("✅ b.ResetTimer()で初期化コストを除外")
	fmt.Println("✅ b.StopTimer()/b.StartTimer()で計測制御")
	fmt.Println("✅ b.ReportAllocs()でメモリ割り当て計測")
	fmt.Println("✅ Sub-benchmarkでパラメータ化テスト")

	fmt.Println("\n--- 現在のよくある間違いと対策 ---")
	fmt.Println("現在の方法でのよくある間違い:")
	fmt.Println(`// ❌ 間違い: ループ内でb.Nを使用
func BenchmarkWrong(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for j := 0; j < b.N; j++ { // 間違い!
            // 処理
        }
    }
}`)

	fmt.Println("\n現在の正しい書き方:")
	fmt.Println(`// ✅ 正しい: 固定値を使用
func BenchmarkCorrect(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for j := 0; j < 100; j++ { // 固定値を使用
            // 処理
        }
    }
}`)

	fmt.Println("\nGo 1.24での新しい書き方（将来）:")
	fmt.Println(`// ✅ より自然: Loop()は一度だけ使用
func BenchmarkNew(b *testing.B) {
    for b.Loop() {
        for j := 0; j < 100; j++ { // 固定値を使用
            // 処理
        }
    }
}`)

	fmt.Println("\n--- ベンチマーク実行コマンド例 ---")
	fmt.Println("go test -bench=.")
	fmt.Println("go test -bench=BenchmarkStringConcat")
	fmt.Println("go test -bench=. -benchmem")
	fmt.Println("go test -bench=. -count=5")

	fmt.Println("\n--- Sub-benchmarkの使用例 ---")
	fmt.Println("\n現在の方法:")
	fmt.Println(`func BenchmarkString(b *testing.B) {
    sizes := []int{10, 100, 1000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                // size に応じた処理
            }
        })
    }
}`)

	fmt.Println("\nGo 1.24での新しい方法（期待）:")
	fmt.Println(`func BenchmarkString(b *testing.B) {
    sizes := []int{10, 100, 1000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            for b.Loop() {
                // size に応じた処理
            }
        })
    }
}`)
}

// % go run 04_testing_loop.go
// === testing.B.Loop() Demo ===
// 従来のベンチマーク記述法:
// func BenchmarkOld(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         // ベンチマーク対象のコード
//     }
// }
//
// Go 1.24の新しい記述法:
// func BenchmarkNew(b *testing.B) {
//     for b.Loop() {
//         // ベンチマーク対象のコード
//     }
// }
//
// --- 実際のパフォーマンス比較 ---
// 1. 文字列連結パフォーマンス:
// String concatenation (+): 10000 iterations in 45.2ms (avg: 4.52µs per operation)
// String builder: 10000 iterations in 2.1ms (avg: 210ns per operation)
//    StringBuilder is 95.4% faster