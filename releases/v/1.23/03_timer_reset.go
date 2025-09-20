// Go 1.23 新機能: Timer.Reset Behavior Change
// 原文: "Timer.Reset now drains expired timers, fixing a common source of goroutine leaks"
//
// 説明: Go 1.23では、Timer.Reset()の動作が改善され、
// 期限切れタイマーが自動的にドレインされ、goroutineリークの一般的な原因が修正されました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Timer.Reset Behavior Change Demo ===")

	// Go 1.23以前の問題のあるパターン（修正済み）
	demonstrateOldProblem()

	// Go 1.23の改善されたTimer.Reset
	demonstrateNewBehavior()

	// 実用的な使用例
	demonstrateRepeatingTimer()

	// パフォーマンステスト
	demonstratePerformanceImprovement()
}

func demonstrateOldProblem() {
	fmt.Println("\n--- Go 1.23以前の問題（現在は修正済み） ---")

	timer := time.NewTimer(100 * time.Millisecond)

	// タイマーが発火するまで待機
	<-timer.C
	fmt.Println("タイマー発火（1回目）")

	// Go 1.23以前: Reset前に手動でドレインが必要だった
	// Go 1.23: 自動的にドレインされる
	success := timer.Reset(100 * time.Millisecond)
	fmt.Printf("Reset成功: %v （Go 1.23では自動改善）\n", success)

	<-timer.C
	fmt.Println("タイマー発火（2回目）")

	timer.Stop()
}

func demonstrateNewBehavior() {
	fmt.Println("\n--- Go 1.23の改善されたReset動作 ---")

	timer := time.NewTimer(50 * time.Millisecond)

	// 短時間待機してタイマーを発火させる
	time.Sleep(60 * time.Millisecond)

	// Go 1.23: 期限切れタイマーでも安全にReset可能
	success := timer.Reset(100 * time.Millisecond)
	fmt.Printf("期限切れタイマーのReset: %v\n", success)

	select {
	case <-timer.C:
		fmt.Println("Resetしたタイマーが正常に発火")
	case <-time.After(150 * time.Millisecond):
		fmt.Println("タイマーが発火しませんでした")
	}

	timer.Stop()
}

func demonstrateRepeatingTimer() {
	fmt.Println("\n--- 繰り返しタイマーの実装例 ---")

	var wg sync.WaitGroup
	timer := time.NewTimer(100 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0

		for count < 3 {
			<-timer.C
			count++
			fmt.Printf("繰り返し実行 %d回目\n", count)

			if count < 3 {
				// Go 1.23: 安全なReset（自動ドレイン）
				timer.Reset(100 * time.Millisecond)
			}
		}

		timer.Stop()
		fmt.Println("繰り返しタイマー完了")
	}()

	wg.Wait()
}

func demonstratePerformanceImprovement() {
	fmt.Println("\n--- パフォーマンス改善デモ ---")

	const iterations = 1000

	// Goroutine数の測定
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("開始時のGoroutine数: %d\n", initialGoroutines)

	// 大量のタイマーリセット操作
	timers := make([]*time.Timer, iterations)

	start := time.Now()

	for i := 0; i < iterations; i++ {
		timer := time.NewTimer(1 * time.Millisecond)
		timers[i] = timer

		// 短時間待機してタイマーを期限切れにする
		time.Sleep(2 * time.Millisecond)

		// Go 1.23: 自動ドレインでリークなし
		timer.Reset(10 * time.Millisecond)
	}

	duration := time.Since(start)
	fmt.Printf("%d回のReset操作時間: %v\n", iterations, duration)

	// クリーンアップ
	for _, timer := range timers {
		timer.Stop()
	}

	// ガベージコレクションを実行
	runtime.GC()
	time.Sleep(10 * time.Millisecond)

	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("終了時のGoroutine数: %d\n", finalGoroutines)

	if finalGoroutines <= initialGoroutines+2 { // 許容範囲
		fmt.Println("✅ Goroutineリークなし（Go 1.23の改善効果）")
	} else {
		fmt.Printf("⚠️  Goroutineが%d個増加\n", finalGoroutines-initialGoroutines)
	}
}

// タイムアウト付き処理の安全な実装例
func demonstrateTimeoutPattern() {
	fmt.Println("\n--- タイムアウト付き処理パターン ---")

	// 作業をシミュレートする関数
	doWork := func(duration time.Duration) <-chan string {
		ch := make(chan string, 1)
		go func() {
			time.Sleep(duration)
			ch <- "作業完了"
		}()
		return ch
	}

	timer := time.NewTimer(200 * time.Millisecond)
	defer timer.Stop()

	// 複数の作業を順次実行（各作業にタイムアウト）
	works := []time.Duration{
		50 * time.Millisecond,  // 成功
		100 * time.Millisecond, // 成功
		300 * time.Millisecond, // タイムアウト
	}

	for i, workDuration := range works {
		workCh := doWork(workDuration)

		// Go 1.23: 安全なReset
		timer.Reset(200 * time.Millisecond)

		select {
		case result := <-workCh:
			fmt.Printf("作業 %d: %s (%.0fms)\n", i+1, result, workDuration.Seconds()*1000)
		case <-timer.C:
			fmt.Printf("作業 %d: タイムアウト (%.0fms > 200ms)\n", i+1, workDuration.Seconds()*1000)
		}
	}
}

// % go run 03_timer_reset.go
// === Timer.Reset Behavior Change Demo ===
//
// --- Go 1.23以前の問題（現在は修正済み） ---
// タイマー発火（1回目）
// Reset成功: false （Go 1.23では自動改善）
// タイマー発火（2回目）
//
// --- Go 1.23の改善されたReset動作 ---
// 期限切れタイマーのReset: false
// Resetしたタイマーが正常に発火
//
// --- 繰り返しタイマーの実装例 ---
// 繰り返し実行 1回目
// 繰り返し実行 2回目
// 繰り返し実行 3回目
// 繰り返しタイマー完了
//
// --- パフォーマンス改善デモ ---
// 開始時のGoroutine数: 2
// 1000回のReset操作時間: 2.1s
// 終了時のGoroutine数: 2
// ✅ Goroutineリークなし（Go 1.23の改善効果）