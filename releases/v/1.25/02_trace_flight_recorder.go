// Go 1.25 新機能: Trace Flight Recorder
// 原文: "The runtime now supports a trace flight recorder, which allows capturing traces in memory and writing out only the most significant segments"
//
// 説明: Go 1.25では、メモリ内でトレースをキャプチャし、重要なセグメントのみを書き出すトレース フライト レコーダー機能が追加されました。
// これにより、パフォーマンス問題の発生時のみトレースデータを保存できるため、オーバーヘッドを最小限に抑えられます。
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#trace
// - Trace Package: https://pkg.go.dev/runtime/trace

//go:build ignore
// +build ignore

package main

// このファイルを実行するには: go run 02_trace_flight_recorder.go

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Trace Flight Recorder Demo ===")

	fmt.Println("Go 1.25のTrace Flight Recorder機能をデモンストレーション")
	fmt.Println("注意: この機能は実際のGo 1.25環境でのみ利用可能です")

	// シンプルなワークロードをシミュレート
	simulateWorkload()

	fmt.Println("\n--- Trace Flight Recorderの特徴 ---")
	fmt.Println("1. メモリ内でのトレース収集")
	fmt.Println("   - 継続的にトレースデータをメモリに保持")
	fmt.Println("   - ディスクI/Oのオーバーヘッドを削減")

	fmt.Println("\n2. 選択的な書き出し")
	fmt.Println("   - 重要なイベント発生時のみファイルに保存")
	fmt.Println("   - パフォーマンス問題の瞬間をキャプチャ")

	fmt.Println("\n3. 軽量なオーバーヘッド")
	fmt.Println("   - 本番環境での常時実行が可能")
	fmt.Println("   - 必要な時だけ詳細トレースを取得")

	fmt.Println("\n--- 使用方法（Go 1.25以降） ---")
	fmt.Println("環境変数での設定:")
	fmt.Println("  GOTRACEBACK=crash")
	fmt.Println("  GOTRACE=flightrecorder")
	fmt.Println("")
	fmt.Println("プログラム内での制御:")
	fmt.Println("  trace.Start() // フライトレコーダー開始")
	fmt.Println("  trace.WriteToFile() // 重要イベント時に保存")
	fmt.Println("  trace.Stop() // レコーダー停止")

	fmt.Println("\n--- 実用例 ---")
	fmt.Println("1. Webサーバーの遅延監視")
	fmt.Println("2. バッチ処理のボトルネック検出")
	fmt.Println("3. 本番環境でのGC分析")
}

func simulateWorkload() {
	fmt.Println("\nワークロードシミュレーション開始...")

	var wg sync.WaitGroup
	const numWorkers = 5

	// 複数のgoroutineでワークロードをシミュレート
	for i := range numWorkers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	// 別のgoroutineで「重要なイベント」をシミュレート
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("重要イベント発生: パフォーマンス問題を検出")
		fmt.Println("   → この時点でTrace Flight Recorderがトレースを保存")
	}()

	wg.Wait()
	fmt.Println("ワークロード完了")
}

func worker(id int) {
	fmt.Printf("Worker %d 開始\n", id)

	// CPU集約的な処理をシミュレート
	for i := range 100000 {
		_ = i * i
	}

	// メモリ割り当てをシミュレート
	data := make([]int, 1000)
	for i := range len(data) {
		data[i] = i
	}

	fmt.Printf("Worker %d 完了\n", id)
}

// % go run 02_trace_flight_recorder.go
// === Trace Flight Recorder Demo ===
// Go 1.25のTrace Flight Recorder機能をデモンストレーション
// 注意: この機能は実際のGo 1.25環境でのみ利用可能です

// ワークロードシミュレーション開始...
// Worker 3 開始
// Worker 0 開始
// Worker 2 開始
// Worker 3 完了
// Worker 4 開始
// Worker 0 完了
// Worker 2 完了
// Worker 4 完了
// Worker 1 開始
// Worker 1 完了
// ワークロード完了

// --- Trace Flight Recorderの特徴 ---
// 1. メモリ内でのトレース収集
//    - 継続的にトレースデータをメモリに保持
//    - ディスクI/Oのオーバーヘッドを削減

// 2. 選択的な書き出し
//    - 重要なイベント発生時のみファイルに保存
//    - パフォーマンス問題の瞬間をキャプチャ

// 3. 軽量なオーバーヘッド
//    - 本番環境での常時実行が可能
//    - 必要な時だけ詳細トレースを取得

// --- 使用方法（Go 1.25以降） ---
// 環境変数での設定:
//   GOTRACEBACK=crash
//   GOTRACE=flightrecorder

// プログラム内での制御:
//   trace.Start() // フライトレコーダー開始
//   trace.WriteToFile() // 重要イベント時に保存
//   trace.Stop() // レコーダー停止

// --- 実用例 ---
// 1. Webサーバーの遅延監視
// 2. バッチ処理のボトルネック検出
// 3. 本番環境でのGC分析
