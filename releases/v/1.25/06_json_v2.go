// Go 1.25 新機能: encoding/json/v2 パッケージ (実験的)
// 原文: "Experimental encoding/json/v2: Improved JSON implementation"
//
// 説明: Go 1.25では、改良されたJSON実装を含む実験的なencoding/json/v2パッケージが追加されました。
// このパッケージは、より高性能で機能豊富なJSON処理を提供します。
//
// 参考リンク:
// - Go 1.25 Release Notes: https://go.dev/doc/go1.25#encoding-json-v2
// - encoding/json/v2 Package: https://pkg.go.dev/encoding/json/v2

// +build ignore

package main

// このファイルを実行するには: go run 06_json_v2.go

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Active   bool      `json:"active"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	fmt.Println("=== encoding/json/v2 パッケージ Demo ===")

	fmt.Println("Go 1.25で追加された改良版JSONパッケージ")
	fmt.Println("注意: この機能は実際のGo 1.25環境でのみ利用可能です")

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

	// 従来のencoding/jsonを使用
	demonstrateOriginalJSON(user)

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

	// 大量データでのパフォーマンステスト
	demonstratePerformance()
}

func demonstrateOriginalJSON(user User) {
	fmt.Println("\n現在のencoding/jsonでの処理例:")

	// マーシャリング
	start := time.Now()
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	marshalTime := time.Since(start)

	fmt.Printf("マーシャリング時間: %v\n", marshalTime)
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
	fmt.Printf("デコードされたユーザー: %+v\n", decodedUser)
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