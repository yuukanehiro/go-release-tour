// Go 1.19 新機能: New Atomic Types
// 原文: "The sync/atomic package defines new atomic types"
//
// 説明: Go 1.19では、sync/atomicパッケージに新しいアトミック型が追加され、
// より安全で使いやすい並行プログラミングが可能になりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== New Atomic Types Demo ===")

	fmt.Println("\n--- 基本的なアトミック型 ---")

	// 新しいアトミック型の使用
	var counter atomic.Int64
	var flag atomic.Bool
	var name atomic.Pointer[string]

	fmt.Printf("初期値 - counter: %d, flag: %t\n", counter.Load(), flag.Load())

	// 値の設定
	counter.Store(42)
	flag.Store(true)

	str := "Hello, Atomic!"
	name.Store(&str)

	fmt.Printf("設定後 - counter: %d, flag: %t, name: %s\n",
		counter.Load(), flag.Load(), *name.Load())

	fmt.Println("\n--- アトミック操作 ---")

	// カウンターの増加
	fmt.Printf("Add前: %d\n", counter.Load())
	newValue := counter.Add(10)
	fmt.Printf("Add(10)後: %d (戻り値: %d)\n", counter.Load(), newValue)

	// Compare-And-Swap
	fmt.Printf("CAS前: %d\n", counter.Load())
	swapped := counter.CompareAndSwap(52, 100)
	fmt.Printf("CAS(52, 100): %t, 現在値: %d\n", swapped, counter.Load())

	// Swap
	oldValue := counter.Swap(200)
	fmt.Printf("Swap(200): 古い値=%d, 新しい値=%d\n", oldValue, counter.Load())

	fmt.Println("\n--- 従来の方法との比較 ---")

	// 従来の方法（Go 1.18以前）
	var oldCounter int64
	fmt.Println("従来の方法:")
	atomic.StoreInt64(&oldCounter, 10)
	fmt.Printf("  atomic.StoreInt64: %d\n", atomic.LoadInt64(&oldCounter))
	atomic.AddInt64(&oldCounter, 5)
	fmt.Printf("  atomic.AddInt64: %d\n", atomic.LoadInt64(&oldCounter))

	// 新しい方法（Go 1.19+）
	var newCounter atomic.Int64
	fmt.Println("新しい方法:")
	newCounter.Store(10)
	fmt.Printf("  newCounter.Store: %d\n", newCounter.Load())
	newCounter.Add(5)
	fmt.Printf("  newCounter.Add: %d\n", newCounter.Load())

	fmt.Println("\n--- ポインター型の使用 ---")

	// 文字列ポインターのアトミック操作
	var atomicStr atomic.Pointer[string]

	str1 := "最初の値"
	str2 := "更新された値"

	atomicStr.Store(&str1)
	fmt.Printf("初期値: %s\n", *atomicStr.Load())

	atomicStr.Store(&str2)
	fmt.Printf("更新後: %s\n", *atomicStr.Load())

	// Compare-And-Swap でポインター更新
	str3 := "CASで設定"
	swapped = atomicStr.CompareAndSwap(&str2, &str3)
	fmt.Printf("CAS結果: %t, 現在値: %s\n", swapped, *atomicStr.Load())

	fmt.Println("\n--- 実用例: カウンターサービス ---")

	// 並行アクセスのテスト
	service := NewCounterService()

	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	// 複数のgoroutineでカウンターを増加
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				service.Increment()
			}
		}()
	}

	wg.Wait()

	expected := int64(numGoroutines * incrementsPerGoroutine)
	actual := service.Value()
	fmt.Printf("期待値: %d, 実際の値: %d, 正確性: %t\n",
		expected, actual, expected == actual)

	fmt.Println("\n--- Boolアトミック型の使用 ---")

	// フラグ管理
	var done atomic.Bool
	var processed atomic.Int64

	// ワーカーgoroutine
	go func() {
		for !done.Load() {
			processed.Add(1)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	// 少し待ってから停止
	time.Sleep(10 * time.Millisecond)
	done.Store(true)
	time.Sleep(1 * time.Millisecond) // ワーカーが停止するまで待機

	fmt.Printf("処理数: %d\n", processed.Load())

	fmt.Println("\n--- 型安全性の利点 ---")

	// コンパイル時エラーの例（実際にはコメントアウト）
	// var unsafeCounter int64
	// counter.Store(unsafeCounter) // コンパイルエラー: 型が一致しない

	fmt.Println("新しいアトミック型の利点:")
	fmt.Println("  - 型安全性: 間違った型での操作を防止")
	fmt.Println("  - アクセス制御: 非アトミックアクセスを防止")
	fmt.Println("  - 使いやすさ: メソッド形式で直感的")
	fmt.Println("  - アライメント: 自動的に適切にアライメント")
	fmt.Println("  - ポインター安全性: unsafe.Pointerが不要")

	fmt.Println("\n--- パフォーマンス比較デモ ---")

	performanceComparison()

	fmt.Println("\n--- 複数の型を持つ構造体 ---")

	stats := &Statistics{}
	stats.IncrementRequests()
	stats.IncrementRequests()
	stats.SetLastError("エラーメッセージ")
	stats.SetActive(true)

	fmt.Printf("統計情報:\n")
	fmt.Printf("  リクエスト数: %d\n", stats.GetRequests())
	fmt.Printf("  アクティブ: %t\n", stats.IsActive())
	fmt.Printf("  最新エラー: %s\n", stats.GetLastError())
}

// カウンターサービスの例
type CounterService struct {
	count atomic.Int64
}

func NewCounterService() *CounterService {
	return &CounterService{}
}

func (c *CounterService) Increment() {
	c.count.Add(1)
}

func (c *CounterService) Value() int64 {
	return c.count.Load()
}

func (c *CounterService) Reset() {
	c.count.Store(0)
}

// 統計情報の管理
type Statistics struct {
	requests atomic.Int64
	active   atomic.Bool
	lastErr  atomic.Pointer[string]
}

func (s *Statistics) IncrementRequests() {
	s.requests.Add(1)
}

func (s *Statistics) GetRequests() int64 {
	return s.requests.Load()
}

func (s *Statistics) SetActive(active bool) {
	s.active.Store(active)
}

func (s *Statistics) IsActive() bool {
	return s.active.Load()
}

func (s *Statistics) SetLastError(err string) {
	s.lastErr.Store(&err)
}

func (s *Statistics) GetLastError() string {
	if ptr := s.lastErr.Load(); ptr != nil {
		return *ptr
	}
	return "エラーなし"
}

// パフォーマンス比較
func performanceComparison() {
	const iterations = 1000000

	// 従来の方法
	var oldCounter int64
	start := time.Now()
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&oldCounter, 1)
	}
	oldDuration := time.Since(start)

	// 新しい方法
	var newCounter atomic.Int64
	start = time.Now()
	for i := 0; i < iterations; i++ {
		newCounter.Add(1)
	}
	newDuration := time.Since(start)

	fmt.Printf("パフォーマンス比較 (%d回の操作):\n", iterations)
	fmt.Printf("  従来の方法: %v\n", oldDuration)
	fmt.Printf("  新しい方法: %v\n", newDuration)
	fmt.Printf("  最終値確認: 従来=%d, 新=%d\n", oldCounter, newCounter.Load())
}

// 設定管理の例
type Config struct {
	debug    atomic.Bool
	maxConns atomic.Int64
	version  atomic.Pointer[string]
}

func NewConfig() *Config {
	c := &Config{}
	c.debug.Store(false)
	c.maxConns.Store(100)

	version := "1.0.0"
	c.version.Store(&version)

	return c
}

func (c *Config) SetDebug(debug bool) {
	c.debug.Store(debug)
}

func (c *Config) IsDebug() bool {
	return c.debug.Load()
}

func (c *Config) SetMaxConnections(max int64) {
	c.maxConns.Store(max)
}

func (c *Config) GetMaxConnections() int64 {
	return c.maxConns.Load()
}

func (c *Config) UpdateVersion(version string) {
	c.version.Store(&version)
}

func (c *Config) GetVersion() string {
	return *c.version.Load()
}

// % go run 02_atomic_types.go
// === New Atomic Types Demo ===
//
// --- 基本的なアトミック型 ---
// 初期値 - counter: 0, flag: false
// 設定後 - counter: 42, flag: true, name: Hello, Atomic!
//
// --- アトミック操作 ---
// Add前: 42
// Add(10)後: 52 (戻り値: 52)
// CAS前: 52
// CAS(52, 100): true, 現在値: 100
// Swap(200): 古い値=100, 新しい値=200