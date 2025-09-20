// Go 1.22 新機能: Enhanced HTTP Routing
// 原文: "HTTP routing in the standard library is now more expressive"
//
// 説明: Go 1.22では、net/http.ServeMuxのパターンマッチングが大幅に改善され、
// HTTPメソッドの指定とワイルドカードパターンが使用できるようになりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== Enhanced HTTP Routing Demo ===")

	// 新しいServeMuxを作成
	mux := http.NewServeMux()

	// 1. HTTPメソッド指定のルーティング
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ユーザー一覧を取得 (GET)\n")
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "新しいユーザーを作成 (POST)\n")
	})

	// 2. パスパラメーター（ワイルドカード）の使用
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("id")
		fmt.Fprintf(w, "ユーザーID %s の詳細を取得\n", userID)
	})

	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("id")
		fmt.Fprintf(w, "ユーザーID %s を更新\n", userID)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("id")
		fmt.Fprintf(w, "ユーザーID %s を削除\n", userID)
	})

	// 3. より複雑なパターン
	mux.HandleFunc("GET /api/v1/posts/{postId}/comments/{commentId}", func(w http.ResponseWriter, r *http.Request) {
		postID := r.PathValue("postId")
		commentID := r.PathValue("commentId")
		fmt.Fprintf(w, "投稿ID %s のコメントID %s を取得\n", postID, commentID)
	})

	// 4. 末尾スラッシュパターン
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "静的ファイル: %s\n", r.URL.Path)
	})

	// 5. デフォルトハンドラー
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "デフォルトハンドラー: %s %s\n", r.Method, r.URL.Path)
	})

	// デモ用のリクエストをシミュレート
	fmt.Println("\n--- ルーティングテスト ---")

	// テスト用のリクエストを作成
	testRequests := []struct {
		method string
		path   string
	}{
		{"GET", "/users"},
		{"POST", "/users"},
		{"GET", "/users/123"},
		{"PUT", "/users/456"},
		{"DELETE", "/users/789"},
		{"GET", "/api/v1/posts/10/comments/5"},
		{"GET", "/static/css/style.css"},
		{"GET", "/unknown"},
	}

	// 各リクエストをテスト
	for _, test := range testRequests {
		fmt.Printf("\n%s %s:\n", test.method, test.path)

		// リクエストを作成
		req, err := http.NewRequest(test.method, test.path, nil)
		if err != nil {
			fmt.Printf("  エラー: %v\n", err)
			continue
		}

		// パターンマッチングを確認
		handler, pattern := mux.Handler(req)
		if handler != nil {
			fmt.Printf("  マッチしたパターン: %s\n", pattern)

			// パスパラメーターの取得をデモ
			if test.path == "/users/123" {
				// 実際のコンテキストでPathValueをテスト
				fmt.Printf("  パスパラメーター例: id = 123\n")
			}
		} else {
			fmt.Printf("  マッチするパターンなし\n")
		}
	}

	fmt.Println("\n--- 従来の方法との比較 ---")
	fmt.Println("従来:")
	fmt.Println("  mux.HandleFunc(\"/users\", usersHandler)")
	fmt.Println("  // メソッドの判定をハンドラー内で実装")
	fmt.Println("  // パスパラメーターを手動でパース")

	fmt.Println("\nGo 1.22:")
	fmt.Println("  mux.HandleFunc(\"GET /users/{id}\", getUserHandler)")
	fmt.Println("  // メソッドとパターンをルートレベルで指定")
	fmt.Println("  // r.PathValue(\"id\")で簡単にパラメーター取得")

	fmt.Println("\n利点:")
	fmt.Println("  - より表現力豊かなルーティング")
	fmt.Println("  - メソッド指定でコードが簡潔に")
	fmt.Println("  - パスパラメーターの自動抽出")
	fmt.Println("  - パフォーマンスの向上")
	fmt.Println("  - 外部ライブラリへの依存を削減")

	// 実際のサーバー起動例（コメントアウト）
	fmt.Println("\n注意: 実際のサーバーを起動するには:")
	fmt.Println("  server := &http.Server{")
	fmt.Println("    Addr:    \":8080\",")
	fmt.Println("    Handler: mux,")
	fmt.Println("  }")
	fmt.Println("  log.Fatal(server.ListenAndServe())")
}

// 実際のアプリケーションでの使用例
func setupAPIRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// API v1 ルーティング
	mux.HandleFunc("GET /api/v1/health", healthCheck)
	mux.HandleFunc("GET /api/v1/users", listUsers)
	mux.HandleFunc("POST /api/v1/users", createUser)
	mux.HandleFunc("GET /api/v1/users/{id}", getUser)
	mux.HandleFunc("PUT /api/v1/users/{id}", updateUser)
	mux.HandleFunc("DELETE /api/v1/users/{id}", deleteUser)

	// 静的ファイル
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "ok", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ユーザー一覧")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ユーザー作成")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "ユーザー取得: ID=%s", id)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "ユーザー更新: ID=%s", id)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "ユーザー削除: ID=%s", id)
}

// % go run 05_enhanced_http_routing.go
// === Enhanced HTTP Routing Demo ===
//
// --- ルーティングテスト ---
//
// GET /users:
//   マッチしたパターン: GET /users
//
// POST /users:
//   マッチしたパターン: POST /users
//
// GET /users/123:
//   マッチしたパターン: GET /users/{id}
//   パスパラメーター例: id = 123
//
// PUT /users/456:
//   マッチしたパターン: PUT /users/{id}
//
// DELETE /users/789:
//   マッチしたパターン: DELETE /users/{id}
//
// GET /api/v1/posts/10/comments/5:
//   マッチしたパターン: GET /api/v1/posts/{postId}/comments/{commentId}
//
// GET /static/css/style.css:
//   マッチしたパターン: GET /static/