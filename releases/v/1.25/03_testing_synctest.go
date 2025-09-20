// Go 1.25 新機能: testing/synctest パッケージ
// 原文: "New testing/synctest package supports testing concurrent code"
//
// 説明: Go 1.25では、並行処理コードのテストを支援する新しいtesting/synctestパッケージが追加されました。
// このパッケージにより、goroutineやchannelを使った並行処理のテストがより簡単で確実に行えるようになります。

//go:build ignore
// +build ignore

package main

// このファイルを実行するには: go run 03_testing_synctest.go

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== testing/synctest Package Demo ===")

	fmt.Println("Go 1.25で追加された並行処理テスト支援パッケージ")
	fmt.Println("注意: この機能は実際のGo 1.25環境でのみ利用可能です")

	// 並行処理のサンプルコード
	demonstrateConcurrentCode()

	fmt.Println("\n--- testing/synctestの特徴 ---")
	fmt.Println("1. 決定論的な並行処理テスト")
	fmt.Println("   - タイミングに依存しないテストが可能")
	fmt.Println("   - レースコンディションの確実な検出")

	fmt.Println("\n2. 仮想時間の制御")
	fmt.Println("   - テスト内で時間を制御可能")
	fmt.Println("   - タイムアウトやスリープのテストが高速化")

	fmt.Println("\n3. goroutineの状態監視")
	fmt.Println("   - デッドロックの検出")
	fmt.Println("   - goroutineリークの防止")

	fmt.Println("\n--- 使用例（Go 1.25以降） ---")
	fmt.Println("import \"testing/synctest\"")
	fmt.Println("")
	fmt.Println("func TestConcurrentFunction(t *testing.T) {")
	fmt.Println("    synctest.Run(func() {")
	fmt.Println("        // 並行処理のテストコード")
	fmt.Println("        // 時間やgoroutineが制御された環境でテスト")
	fmt.Println("    })")
	fmt.Println("}")

	fmt.Println("\n--- 従来の問題と解決策 ---")
	fmt.Println("従来の問題:")
	fmt.Println("  - time.Sleep()で待機する不安定なテスト")
	fmt.Println("  - レースコンディションの不確実な再現")
	fmt.Println("  - 長時間実行されるテスト")
	fmt.Println("")
	fmt.Println("synctest による解決:")
	fmt.Println("  - 仮想時間による高速テスト")
	fmt.Println("  - 決定論的な実行順序")
	fmt.Println("  - 確実なgoroutine完了待機")
}

func demonstrateConcurrentCode() {
	fmt.Println("\n並行処理サンプル実行...")

	var wg sync.WaitGroup
	ch := make(chan int, 3)

	// データ生成goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := range 3 {
			ch <- i + 1
			fmt.Printf("送信: %d\n", i+1)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// データ消費goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Printf("受信: %d\n", value)
			time.Sleep(30 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("並行処理完了")

	fmt.Println("\n※ このコードをtesting/synctestでテストする場合:")
	fmt.Println("  - time.Sleep()は仮想時間で瞬時に処理")
	fmt.Println("  - goroutineの実行順序が制御可能")
	fmt.Println("  - デッドロックやリークの自動検出")
}

// % go run 03_testing_synctest.go
// === testing/synctest Package Demo ===
// Go 1.25で追加された並行処理テスト支援パッケージ
// 注意: この機能は実際のGo 1.25環境でのみ利用可能です

// 並行処理サンプル実行...
// 送信: 1
// 受信: 1
// 送信: 2
// 受信: 2
// 送信: 3
// 受信: 3
// 並行処理完了

// ※ このコードをtesting/synctestでテストする場合:
//   - time.Sleep()は仮想時間で瞬時に処理
//   - goroutineの実行順序が制御可能
//   - デッドロックやリークの自動検出

// --- testing/synctestの特徴 ---
// 1. 決定論的な並行処理テスト
//    - タイミングに依存しないテストが可能
//    - レースコンディションの確実な検出

// 2. 仮想時間の制御
//    - テスト内で時間を制御可能
//    - タイムアウトやスリープのテストが高速化

// 3. goroutineの状態監視
//    - デッドロックの検出
//    - goroutineリークの防止

// --- 使用例（Go 1.25以降） ---
// import "testing/synctest"

// func TestConcurrentFunction(t *testing.T) {
//     synctest.Run(func() {
//         // 並行処理のテストコード
//         // 時間やgoroutineが制御された環境でテスト
//     })
// }

// --- 従来の問題と解決策 ---
// 従来の問題:
//   - time.Sleep()で待機する不安定なテスト
//   - レースコンディションの不確実な再現
//   - 長時間実行されるテスト

// synctest による解決:
//   - 仮想時間による高速テスト
//   - 決定論的な実行順序
//   - 確実なgoroutine完了待機
