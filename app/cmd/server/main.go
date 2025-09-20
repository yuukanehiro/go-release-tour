// Go Release Tour - Go 1.25の新機能をインタラクティブに学習するWebチュートリアル
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-release-tour/app/internal/handlers"
	"go-release-tour/app/internal/lessons"
	"go-release-tour/app/internal/templates"
	"go-release-tour/app/internal/types"
)


func main() {
	server := &types.Server{
		Lessons: make(map[string][]types.Lesson),
	}
	lessons.LoadLessons(server)

	// 静的ファイル
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// APIエンドポイント
	http.HandleFunc("/api/versions", handlers.HandleVersions(server))
	http.HandleFunc("/api/lessons", handlers.HandleLessons(server))
	http.HandleFunc("/api/run", handlers.HandleRun)

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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}



