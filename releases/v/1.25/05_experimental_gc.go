// Go 1.25 新機能: 実験的ガベージコレクター (Green Tea GC)
// 原文: "An experimental garbage collector is available, with potential 10-40% reduction in garbage collection overhead"
//
// 説明: Go 1.25では、Green Tea GCと呼ばれる実験的なガベージコレクターが利用可能になりました。
// このGCは10-40%のガベージコレクションオーバーヘッド削減の可能性があります。

// +build ignore

package main

// このファイルを実行するには: go run 05_experimental_gc.go

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== 実験的ガベージコレクター (Green Tea GC) Demo ===")

	// 現在のGC統計を取得
	var gcStats runtime.MemStats
	runtime.ReadMemStats(&gcStats)

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GC実行回数: %d\n", gcStats.NumGC)

	fmt.Println("\n--- Green Tea GCの特徴 ---")
	fmt.Println("1. 10-40%のGCオーバーヘッド削減")
	fmt.Println("   - より効率的なメモリ管理")
	fmt.Println("   - レイテンシの改善")

	fmt.Println("\n2. 改善されたローカリティとCPUスケーラビリティ")
	fmt.Println("   - メモリアクセスパターンの最適化")
	fmt.Println("   - マルチコアでの性能向上")

	fmt.Println("\n3. 実験的機能の有効化")
	fmt.Println("   環境変数: GOEXPERIMENT=greenteagc")

	// メモリ使用状況のデモンストレーション
	demonstrateMemoryUsage()

	fmt.Println("\n--- 使用方法 ---")
	fmt.Println("1. 環境変数での有効化:")
	fmt.Println("   export GOEXPERIMENT=greenteagc")
	fmt.Println("   go run 05_experimental_gc.go")

	fmt.Println("\n2. ビルド時の指定:")
	fmt.Println("   GOEXPERIMENT=greenteagc go build main.go")

	fmt.Println("\n3. 複数実験機能の組み合わせ:")
	fmt.Println("   GOEXPERIMENT=greenteagc,otherflag go run main.go")

	fmt.Println("\n--- 注意事項 ---")
	fmt.Println("⚠️  実験的機能のため本番環境での使用は非推奨")
	fmt.Println("⚠️  将来のバージョンで仕様が変更される可能性")
	fmt.Println("⚠️  パフォーマンステストでの効果測定を推奨")

	// 最終的なGC統計を表示
	runtime.GC() // 手動でGCを実行
	runtime.ReadMemStats(&gcStats)
	fmt.Printf("\nGC実行後の統計: %d回実行\n", gcStats.NumGC)
}

func demonstrateMemoryUsage() {
	fmt.Println("\nメモリ使用量デモンストレーション...")

	// 初期メモリ統計
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("初期ヒープサイズ: %d KB\n", m1.HeapAlloc/1024)

	// 大量のメモリ割り当てをシミュレート
	const iterations = 1000
	data := make([][]int, iterations)

	start := time.Now()
	for i := range iterations {
		// 各イテレーションで1MBのスライスを作成
		data[i] = make([]int, 250000) // 250000 * 4 bytes ≈ 1MB

		// 定期的にGCを促す
		if i%100 == 0 {
			runtime.GC()
		}
	}
	elapsed := time.Since(start)

	// 最終メモリ統計
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("最終ヒープサイズ: %d KB\n", m2.HeapAlloc/1024)
	fmt.Printf("総GC実行回数: %d\n", m2.NumGC-m1.NumGC)
	fmt.Printf("GC合計時間: %v\n", time.Duration(m2.PauseTotalNs-m1.PauseTotalNs))
	fmt.Printf("メモリ割り当て処理時間: %v\n", elapsed)

	fmt.Println("\n※ Green Tea GCを使用した場合:")
	fmt.Println("  - GC実行時間がより短縮される可能性")
	fmt.Println("  - メモリ割り当て処理が高速化される可能性")
	fmt.Println("  - 特に大規模なアプリケーションで効果が期待される")

	// メモリをクリア
	data = nil
	runtime.GC()
}