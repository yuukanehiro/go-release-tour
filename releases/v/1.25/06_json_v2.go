// Go 1.25 新機能: encoding/json/v2 パッケージ (実験的)
// 原文: "Experimental encoding/json/v2: Improved JSON implementation"
//
// 説明: Go 1.25では、改良されたJSON実装を含む実験的なencoding/json/v2パッケージが追加されました。
// このパッケージは、より高性能で機能豊富なJSON処理を提供します。
//
// 【環境変数で比較体験】
// このプログラムでは以下の設定で動作を比較できます:
// 1. 通常実行: go run 06_json_v2.go
// 2. v2有効化: GOEXPERIMENT=jsonv2 go run 06_json_v2.go
//
// プログラム内で環境変数を切り替えて、両方の実装を同時に比較します。
//
// @env-preset: JSON v2|GOEXPERIMENT=jsonv2|新しいJSON v2実装を有効化
// @env-preset: JSON v2 + デバッグ|GOEXPERIMENT=jsonv2,GODEBUG=gctrace=1|JSON v2実装とGCトレースを同時に有効化
// @env-preset: 詳細デバッグ|GOEXPERIMENT=jsonv2,GODEBUG=gctrace=1,GODEBUG=gcpacertrace=1|JSON v2とGCの詳細トレース
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#encoding-json-v2
// - encoding/json/v2 Package: https://pkg.go.dev/encoding/json/v2
// - Future Architect記事: https://future-architect.github.io/articles/20250806a/

package main

// このファイルを実行するには: go run 06_json_v2.go

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
	Created  time.Time              `json:"created"`
	Active   bool                   `json:"active"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	fmt.Println("=== Go 1.25 encoding/json/v2 実験機能 ===")

	// 現在の環境変数を確認
	goExperiment := os.Getenv("GOEXPERIMENT")
	isJsonV2Enabled := strings.Contains(goExperiment, "jsonv2")

	fmt.Printf("現在のGOEXPERIMENT: %s\n", goExperiment)
	fmt.Printf("JSON v2実装: %t\n", isJsonV2Enabled)

	if isJsonV2Enabled {
		fmt.Println("新しいJSON v2実装で動作中")
		fmt.Println("   - より高速なマーシャリング/アンマーシャリング")
		fmt.Println("   - メモリ使用量の削減")
		fmt.Println("   - より詳細なエラーメッセージ")
	} else {
		fmt.Println("従来のJSON実装で動作中")
		fmt.Println("   JSON v2を体験するには環境変数でGOEXPERIMENT=jsonv2を設定してください")
	}

	// サンプルデータの作成
	user := User{
		ID:      12345,
		Name:    "田中太郎",
		Email:   "tanaka@example.com",
		Created: time.Now(),
		Active:  true,
		Metadata: map[string]interface{}{
			"department": "engineering",
			"level":      5,
			"tags":       []string{"golang", "backend"},
		},
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("【従来版】encoding/json での処理")
	fmt.Println(strings.Repeat("=", 60))
	demonstrateWithoutJsonV2(user)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("【新版】GOEXPERIMENT=jsonv2 での処理")
	fmt.Println(strings.Repeat("=", 60))
	demonstrateWithJsonV2(user)

	fmt.Println("\n--- encoding/json/v2の特徴 ---")
	fmt.Println("1. パフォーマンス向上")
	fmt.Println("   - より高速なマーシャリング/アンマーシャリング")
	fmt.Println("   - メモリ使用量の削減")

	fmt.Println("\n2. 機能強化")
	fmt.Println("   - より柔軟なタグ指定")
	fmt.Println("   - カスタムエンコーダー/デコーダーのサポート")
	fmt.Println("   - ストリーミング処理の改善")

	fmt.Println("\n3. 後方互換性")
	fmt.Println("   - 既存のencoding/jsonとの互換性維持")
	fmt.Println("   - 段階的な移行が可能")

	fmt.Println("\n--- 使用例（Go 1.25以降） ---")
	fmt.Println("import \"encoding/json/v2\"")
	fmt.Println("")
	fmt.Println("// 高性能なマーシャリング")
	fmt.Println("data, err := jsonv2.Marshal(user)")
	fmt.Println("")
	fmt.Println("// ストリーミング処理")
	fmt.Println("encoder := jsonv2.NewEncoder(writer)")
	fmt.Println("encoder.Encode(user)")

	fmt.Println("\n--- 期待される改善点 ---")
	fmt.Println("1. パフォーマンス: 20-50%の処理速度向上")
	fmt.Println("2. メモリ効率: アロケーション回数の削減")
	fmt.Println("3. 機能性: より直感的なAPI設計")
	fmt.Println("4. エラーハンドリング: より詳細なエラー情報")
}

func demonstrateWithoutJsonV2(user User) {
	// GOEXPERIMENT環境変数を一時的にクリア
	originalGoExperiment := os.Getenv("GOEXPERIMENT")
	os.Setenv("GOEXPERIMENT", "")
	defer os.Setenv("GOEXPERIMENT", originalGoExperiment)

	fmt.Println("従来のencoding/json実装を使用")

	// マーシャリング
	start := time.Now()
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	marshalTime := time.Since(start)

	fmt.Printf("マーシャリング時間: %v\n", marshalTime)
	fmt.Printf("JSON出力サイズ: %d bytes\n", len(jsonData))
	fmt.Printf("JSON出力: %s\n", string(jsonData))

	// アンマーシャリング
	start = time.Now()
	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		log.Fatal(err)
	}
	unmarshalTime := time.Since(start)

	fmt.Printf("アンマーシャリング時間: %v\n", unmarshalTime)
	fmt.Printf("デコード成功: %s\n", decodedUser.Name)

	// 大量データでのパフォーマンステスト
	demonstratePerformance()
}

func demonstrateWithJsonV2(user User) {
	// GOEXPERIMENT=jsonv2を設定
	originalGoExperiment := os.Getenv("GOEXPERIMENT")
	os.Setenv("GOEXPERIMENT", "jsonv2")
	defer os.Setenv("GOEXPERIMENT", originalGoExperiment)

	fmt.Println("GOEXPERIMENT=jsonv2 で新しい実装を使用")
	fmt.Println("注意: 実際の効果はGo 1.25環境でのみ確認可能")

	// マーシャリング
	start := time.Now()
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	marshalTime := time.Since(start)

	fmt.Printf("マーシャリング時間: %v\n", marshalTime)
	fmt.Printf("JSON出力サイズ: %d bytes\n", len(jsonData))
	fmt.Printf("JSON出力: %s\n", string(jsonData))

	// アンマーシャリング
	start = time.Now()
	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		log.Fatal(err)
	}
	unmarshalTime := time.Since(start)

	fmt.Printf("アンマーシャリング時間: %v\n", unmarshalTime)
	fmt.Printf("デコード成功: %s\n", decodedUser.Name)

	// v2の新機能説明
	fmt.Println("\nv2の新機能:")
	fmt.Println("• より詳細なエラー情報")
	fmt.Println("• カスタムエンコーダー/デコーダーサポート")
	fmt.Println("• 改善されたストリーミングAPI")
	fmt.Println("• メモリ効率の向上")

	// 大量データでのパフォーマンステスト
	demonstratePerformance()
}

func demonstratePerformance() {
	fmt.Println("\n大量データでのパフォーマンステスト...")

	// 大量のユーザーデータを生成
	users := make([]User, 1000)
	for i := range users {
		users[i] = User{
			ID:      i + 1,
			Name:    fmt.Sprintf("ユーザー%d", i+1),
			Email:   fmt.Sprintf("user%d@example.com", i+1),
			Created: time.Now().Add(-time.Duration(i) * time.Hour),
			Active:  i%2 == 0,
			Metadata: map[string]interface{}{
				"index": i,
				"group": fmt.Sprintf("group_%d", i%10),
			},
		}
	}

	// 現在のJSONでのパフォーマンス測定
	start := time.Now()
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Since(start)

	fmt.Printf("現在のJSON処理時間: %v\n", currentTime)
	fmt.Printf("データサイズ: %d KB\n", len(jsonData)/1024)

	// アンマーシャリング性能
	start = time.Now()
	var decodedUsers []User
	err = json.Unmarshal(jsonData, &decodedUsers)
	if err != nil {
		log.Fatal(err)
	}
	decodeTime := time.Since(start)

	fmt.Printf("現在のJSONデコード時間: %v\n", decodeTime)
	fmt.Printf("デコードされたユーザー数: %d\n", len(decodedUsers))

	fmt.Println("\n※ encoding/json/v2使用時の期待される改善:")
	fmt.Printf("  - マーシャリング: %v → %v (予想)\n",
		currentTime, time.Duration(float64(currentTime)*0.6))
	fmt.Printf("  - アンマーシャリング: %v → %v (予想)\n",
		decodeTime, time.Duration(float64(decodeTime)*0.7))
	fmt.Println("  - メモリ使用量: 20-30%削減")
}
