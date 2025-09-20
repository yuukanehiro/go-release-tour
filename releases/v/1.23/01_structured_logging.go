// Go 1.23 新機能: Structured Logging (log/slog)
// 原文: "The log/slog package provides structured logging with levels, contexts, and structured output"
//
// 説明: Go 1.21でlog/slogパッケージが追加され、Go 1.23でさらに改良されました。
// 構造化ログ、ログレベル、コンテキスト対応の高度なロギング機能を提供します。
//
// 参考リンク:
// - Go 1.23 Release Notes: https://go.dev/doc/go1.23#log-slog
// - log/slog Package: https://pkg.go.dev/log/slog
//
// 注意: この例は log/slog の概念を示すもので、実際の環境では標準のlogパッケージを使用しています。

//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("=== Structured Logging Demo ===")
	fmt.Println("注意: これはlog/slogの概念例です。実際の環境では標準logパッケージを使用します。")

	// 構造化ログのシミュレーション
	logEntry := func(level, message string, fields ...interface{}) {
		entry := map[string]interface{}{
			"time":    time.Now().Format(time.RFC3339),
			"level":   level,
			"message": message,
		}

		// フィールドをペアで処理
		for i := 0; i < len(fields)-1; i += 2 {
			if key, ok := fields[i].(string); ok {
				entry[key] = fields[i+1]
			}
		}

		jsonData, _ := json.MarshalIndent(entry, "", "  ")
		fmt.Println(string(jsonData))
	}

	// 基本的なログ出力
	logEntry("INFO", "アプリケーション開始")
	logEntry("DEBUG", "デバッグ情報", "version", "1.23", "env", "development")

	// 構造化データでのログ出力
	user := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{1, "田中太郎"}

	logEntry("INFO", "ユーザーログイン",
		"user_id", user.ID,
		"user_name", user.Name,
		"login_time", time.Now().Format(time.RFC3339),
	)

	// エラーログ
	err := fmt.Errorf("データベース接続エラー")
	logEntry("ERROR", "システムエラー発生",
		"error", err.Error(),
		"component", "database",
		"retry_count", 3,
	)

	// コンテキスト付きログ（シミュレーション）
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	requestID := ctx.Value("request_id")
	logEntry("INFO", "リクエスト処理完了",
		"request_id", requestID,
		"duration_ms", 150,
		"status", "success",
	)

	// ログレベル別出力例
	fmt.Println("\n--- ログレベル別出力 ---")

	logEntry("DEBUG", "デバッグメッセージ（開発時のみ）")
	logEntry("INFO", "情報メッセージ（一般的な動作）")
	logEntry("WARN", "警告メッセージ（注意が必要）")
	logEntry("ERROR", "エラーメッセージ（問題発生）")

	// テキスト形式ログ
	fmt.Println("\n--- テキスト形式ログ ---")
	log.SetPrefix("[APP] ")
	log.Println("テキスト形式でのログ出力: service=go-release-tour version=1.23")

	// グループ化されたログ（シミュレーション）
	fmt.Println("\n--- グループ化ログ ---")
	requestData := map[string]interface{}{
		"request": map[string]interface{}{
			"method": "GET",
			"path":   "/api/lessons",
			"ip":     "192.168.1.100",
		},
		"response": map[string]interface{}{
			"status":   200,
			"size":     1024,
			"duration": "45ms",
		},
	}

	entry := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"level":   "INFO",
		"message": "HTTP リクエスト",
	}

	for k, v := range requestData {
		entry[k] = v
	}

	jsonData, _ := json.MarshalIndent(entry, "", "  ")
	fmt.Println(string(jsonData))

	fmt.Println("\n--- 従来のlogパッケージとの比較 ---")
	fmt.Println("従来: log.Printf(\"User %s logged in at %v\", name, time)")
	fmt.Println("新機能: logger.Info(\"ユーザーログイン\", \"user\", name, \"time\", time)")
	fmt.Println("利点:")
	fmt.Println("  - 構造化データで検索・分析が容易")
	fmt.Println("  - JSON形式で機械読み取り可能")
	fmt.Println("  - ログレベルによる出力制御")
	fmt.Println("  - コンテキスト情報の自動追加")
}

// % go run 01_structured_logging.go
// === Structured Logging Demo ===
// {"time":"2024-01-15T10:30:45.123Z","level":"INFO","msg":"アプリケーション開始"}
// {"time":"2024-01-15T10:30:45.124Z","level":"DEBUG","msg":"デバッグ情報","version":"1.23","env":"development"}
// {"time":"2024-01-15T10:30:45.125Z","level":"INFO","msg":"ユーザーログイン","user_id":1,"user_name":"田中太郎","login_time":"2024-01-15T10:30:45.125Z"}
// {"time":"2024-01-15T10:30:45.126Z","level":"ERROR","msg":"システムエラー発生","error":"データベース接続エラー","component":"database","retry_count":3}
// {"time":"2024-01-15T10:30:45.127Z","level":"INFO","msg":"リクエスト処理完了","duration_ms":150,"status":"success"}
//
// --- ログレベル別出力 ---
// {"time":"2024-01-15T10:30:45.128Z","level":"DEBUG","msg":"デバッグメッセージ（開発時のみ）"}
// {"time":"2024-01-15T10:30:45.129Z","level":"INFO","msg":"情報メッセージ（一般的な動作）"}
// {"time":"2024-01-15T10:30:45.130Z","level":"WARN","msg":"警告メッセージ（注意が必要）"}
// {"time":"2024-01-15T10:30:45.131Z","level":"ERROR","msg":"エラーメッセージ（問題発生）"}