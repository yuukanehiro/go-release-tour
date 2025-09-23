// Go Release Tour - Interactive Web Tutorial for Go Language Features
//
// This application provides an interactive web-based tutorial for learning
// new features introduced in Go versions 1.18 through 1.25. Each lesson
// includes:
// - Official documentation references
// - Executable code examples
// - Practical use cases
// - Performance comparisons
// - Best practices
//
// Architecture:
// - Backend: Go HTTP server with lesson management
// - Frontend: Vanilla JavaScript with CodeMirror integration
// - Storage: File-based lesson content with dynamic loading
// - Development: Docker Compose with hot reload support
//
// API Endpoints:
// - GET /api/versions: Available Go versions
// - GET /api/lessons?version=X.XX: Lessons for specific version
// - POST /api/run: Execute Go code snippets
//
// Static Assets:
// - /static/: CSS, JS, images, and other static resources
//
// Environment Variables:
// - APP_PORT: Server port (default: 8080)
// - GO_VERSION: Go version for display purposes
//
// Usage:
//
//	go run ./app/cmd/server
//	# or with Docker Compose:
//	docker-compose up
//
// Author: Go Release Tour Project
// License: MIT
// Go Version: 1.24+
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-release-tour/app/internal/handlers"
	"go-release-tour/app/internal/lessons"
	"go-release-tour/app/internal/templates"
	"go-release-tour/app/internal/types"
)

// addNoCacheHeaders は開発環境でキャッシュを無効化するミドルウェア
func addNoCacheHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

func main() {
	appServer := &types.Server{
		Lessons: make(map[string][]types.Lesson),
	}
	lessons.LoadLessons(appServer)

	// 静的ファイル（開発環境用キャッシュ無効化）
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", addNoCacheHeaders(staticHandler))

	// テストファイル（開発環境専用）
	http.Handle("/tests/", http.StripPrefix("/tests/", http.FileServer(http.Dir("tests"))))

	// APIエンドポイント
	http.HandleFunc("/api/versions", handlers.HandleVersions(appServer))
	http.HandleFunc("/api/lessons", handlers.HandleLessons(appServer))
	http.HandleFunc("/api/run", handlers.HandleRun)
	http.HandleFunc("/api/version-info", handlers.HandleVersionInfo)

	// メインページ
	http.HandleFunc("/", templates.HandleIndex)

	// 環境変数からポートを取得、デフォルトは8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	goVersion := os.Getenv("GO_VERSION")
	if goVersion != "" {
		fmt.Printf("Go Release Tour server (Go %s) starting on :%s\n", goVersion, port)
	} else {
		fmt.Printf("Go Release Tour server starting on :%s\n", port)
	}
	fmt.Printf("Visit http://localhost:%s\n", port)

	httpServer := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
}
