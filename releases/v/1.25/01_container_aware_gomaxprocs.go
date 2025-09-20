// Go 1.25 新機能: Container-aware GOMAXPROCS
// 原文: "The runtime now considers CPU bandwidth limits when setting the default value for GOMAXPROCS in Linux containers"
//
// 説明: Go 1.25では、Linuxコンテナ内でのCPU帯域制限を考慮してGOMAXPROCSのデフォルト値を設定するようになりました。
// これにより、コンテナのCPU制限に応じて適切なgoroutineの並行度が自動設定されます。
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#runtime
// - Runtime Package: https://pkg.go.dev/runtime#GOMAXPROCS

//go:build ignore
// +build ignore

package main

// このファイルを実行するには: go run 01_container_aware_gomaxprocs.go

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== Container-aware GOMAXPROCS Demo ===")

	// 現在のGOMAXPROCSの値を取得
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("現在のGOMAXPROCS: %d\n", maxProcs)

	// CPU数を取得
	numCPU := runtime.NumCPU()
	fmt.Printf("システムのCPU数: %d\n", numCPU)

	// Go 1.25では、コンテナのCPU制限がある場合、
	// GOMAXPROCSがそれに応じて調整される

	// 実際のCPU使用状況をシミュレート
	fmt.Println("\nCPU集約的なタスクを実行中...")

	start := time.Now()
	done := make(chan bool, maxProcs)

	// GOMAXPROCSの数だけgoroutineを起動
	for i := 0; i < maxProcs; i++ {
		go func(id int) {
			// CPU集約的な処理をシミュレート
			for j := 0; j < 1000000; j++ {
				_ = j * j
			}
			fmt.Printf("Goroutine %d 完了\n", id)
			done <- true
		}(i)
	}

	// すべてのgoroutineの完了を待機
	for i := 0; i < maxProcs; i++ {
		<-done
	}

	elapsed := time.Since(start)
	fmt.Printf("\n実行時間: %v\n", elapsed)

	fmt.Println("\n--- コンテナでの動作について ---")
	fmt.Println("Docker等のコンテナ環境で以下のようにCPU制限を設定した場合:")
	fmt.Println("  docker run --cpus=2.5 your-go-app")
	fmt.Println("Go 1.25では、GOMAXPROCSが自動的に2または3に調整されます")
	fmt.Println("(従来は物理CPU数がそのまま使用されていました)")

	// 環境変数の確認
	fmt.Println("\n--- 関連する環境変数 ---")
	fmt.Println("GOMAXPROCS環境変数で手動設定も可能:")
	fmt.Println("  export GOMAXPROCS=4")
	fmt.Println("  go run 01_container_aware_gomaxprocs.go")
}

// % go run 01_container_aware_gomaxprocs.go
// === Container-aware GOMAXPROCS Demo ===
// 現在のGOMAXPROCS: 8
// システムのCPU数: 8

// CPU集約的なタスクを実行中...
// Goroutine 7 完了
// Goroutine 0 完了
// Goroutine 1 完了
// Goroutine 4 完了
// Goroutine 5 完了
// Goroutine 2 完了
// Goroutine 6 完了
// Goroutine 3 完了

// 実行時間: 1.389042ms

// --- コンテナでの動作について ---
// Docker等のコンテナ環境で以下のようにCPU制限を設定した場合:
//   docker run --cpus=2.5 your-go-app
// Go 1.25では、GOMAXPROCSが自動的に2または3に調整されます
// (従来は物理CPU数がそのまま使用されていました)

// --- 関連する環境変数 ---
// GOMAXPROCS環境変数で手動設定も可能:
//   export GOMAXPROCS=4
//   go run 01_container_aware_gomaxprocs.go
