// Go 1.24 新機能: Swiss Tables Map Implementation
// 原文: "New map implementation based on Swiss Tables, more efficient memory allocation"
//
// 説明: Swiss Tablesベースの新しいマップ実装により、大きなマップ（>1024エントリ）で約30%の性能向上を実現。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== Swiss Tables Map Performance Demo ===")
	fmt.Println("注意: Swiss Tables実装はGo 1.24の内部変更で、ユーザーコードは変更不要")
	fmt.Println("以下は通常のmap操作ですが、Go 1.24では自動的に高速化されます")
	fmt.Println()

	// メモリ使用量の測定
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	fmt.Println("--- Large Map Performance Test ---")

	// 大きなマップの作成とアクセステスト
	start := time.Now()
	largeMap := make(map[int]string, 10000) // Pre-sized map

	// マップに大量のデータを追加
	for i := 0; i < 10000; i++ {
		largeMap[i] = fmt.Sprintf("value-%d", i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Large map creation (10,000 entries): %v\n", elapsed)

	// アクセス性能テスト
	start = time.Now()
	accessCount := 0
	for i := 0; i < 50000; i++ {
		if _, exists := largeMap[i%10000]; exists {
			accessCount++
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("Large map access (50,000 lookups): %v\n", elapsed)
	fmt.Printf("Successful lookups: %d\n", accessCount)

	runtime.GC()
	runtime.ReadMemStats(&m2)
	fmt.Printf("Memory used by large map: %d KB\n", (m2.HeapInuse-m1.HeapInuse)/1024)

	fmt.Println("\n--- Map Iteration Performance ---")

	// イテレーション性能テスト
	start = time.Now()
	iterCount := 0
	for k, v := range largeMap {
		if k%1000 == 0 {
			_ = v // Use the value to prevent optimization
		}
		iterCount++
	}
	elapsed = time.Since(start)
	fmt.Printf("Map iteration (%d entries): %v\n", iterCount, elapsed)

	fmt.Println("\n--- Sparse Map Performance ---")

	// スパースマップ（低負荷率）のテスト
	sparseMap := make(map[int]string, 100000) // Large size, few entries
	for i := 0; i < 1000; i += 100 { // Only 10 entries in a 100k-sized map
		sparseMap[i] = fmt.Sprintf("sparse-%d", i)
	}

	start = time.Now()
	sparseIterCount := 0
	for k, v := range sparseMap {
		_ = k
		_ = v
		sparseIterCount++
	}
	elapsed = time.Since(start)
	fmt.Printf("Sparse map iteration (%d entries): %v\n", sparseIterCount, elapsed)

	fmt.Println("\n--- Pre-sized vs Dynamic Map ---")

	// Pre-sizedマップの性能
	start = time.Now()
	presizedMap := make(map[int]string, 5000)
	for i := 0; i < 5000; i++ {
		presizedMap[i] = fmt.Sprintf("presized-%d", i)
	}
	presizedTime := time.Since(start)

	// 動的マップの性能
	start = time.Now()
	dynamicMap := make(map[int]string)
	for i := 0; i < 5000; i++ {
		dynamicMap[i] = fmt.Sprintf("dynamic-%d", i)
	}
	dynamicTime := time.Since(start)

	fmt.Printf("Pre-sized map creation: %v\n", presizedTime)
	fmt.Printf("Dynamic map creation: %v\n", dynamicTime)
	improvement := float64(dynamicTime-presizedTime) / float64(dynamicTime) * 100
	fmt.Printf("Pre-sized map improvement: %.1f%%\n", improvement)

	fmt.Println("\n--- String Key Performance ---")

	// 文字列キーマップのテスト
	stringMap := make(map[string]int, 1000)
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	start = time.Now()
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%s-%d", words[i%len(words)], i)
		stringMap[key] = i
	}
	elapsed = time.Since(start)
	fmt.Printf("String key map creation (1,000 entries): %v\n", elapsed)

	// 文字列キーアクセス
	start = time.Now()
	foundCount := 0
	for i := 0; i < 5000; i++ {
		key := fmt.Sprintf("%s-%d", words[i%len(words)], i%1000)
		if _, exists := stringMap[key]; exists {
			foundCount++
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("String key lookups (5,000 attempts): %v\n", elapsed)
	fmt.Printf("Found entries: %d\n", foundCount)

	fmt.Println("\n--- Swiss Tables Benefits ---")
	fmt.Println("✅ Large maps (>1024 entries): ~30% faster access")
	fmt.Println("✅ Pre-sized maps: ~35% faster assignment")
	fmt.Println("✅ Iteration: ~10% faster overall, ~60% for sparse maps")
	fmt.Println("✅ Better memory locality")
	fmt.Println("✅ Reduced memory overhead")
	fmt.Println("✅ Backward compatible - no code changes needed")

	fmt.Println("\n--- Memory Stats ---")
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	fmt.Printf("Total heap allocated: %d KB\n", m3.HeapAlloc/1024)
	fmt.Printf("Number of GC cycles: %d\n", m3.NumGC)
}

// ベンチマーク例（実際のベンチマーク用）
func benchmarkMapAccess(size int) time.Duration {
	m := make(map[int]string, size)
	for i := 0; i < size; i++ {
		m[i] = fmt.Sprintf("value-%d", i)
	}

	start := time.Now()
	for i := 0; i < size*10; i++ {
		_ = m[i%size]
	}
	return time.Since(start)
}

// % go run 03_swiss_tables_maps.go
// === Swiss Tables Map Performance Demo ===
// --- Large Map Performance Test ---
// Large map creation (10,000 entries): 2.234ms
// Large map access (50,000 lookups): 1.845ms
// Successful lookups: 50000
// Memory used by large map: 892 KB
//
// --- Map Iteration Performance ---
// Map iteration (10000 entries): 234µs
//
// --- Sparse Map Performance ---
// Sparse map iteration (10 entries): 12µs
//
// --- Pre-sized vs Dynamic Map ---
// Pre-sized map creation: 1.123ms
// Dynamic map creation: 1.789ms
// Pre-sized map improvement: 37.2%