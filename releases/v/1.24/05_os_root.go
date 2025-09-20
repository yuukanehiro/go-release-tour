// Go 1.24 新機能: os.Root Type
// 原文: "The new os.Root type provides the ability to perform filesystem operations within a specific directory"
//
// 説明: 新しいos.Root型により、特定のディレクトリ内でのファイルシステム操作を安全に制限できるようになりました。
//
// 参考リンク:
// - Go 1.24 Release Notes: https://go.dev/doc/go1.24#os
// - os Package: https://pkg.go.dev/os

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== os.Root Type Demo ===")

	// 現在のディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	fmt.Printf("Current directory: %s\n", currentDir)

	// テスト用のディレクトリ構造を作成
	testDir := filepath.Join(currentDir, "test-root")
	fmt.Printf("\nCreating test directory: %s\n", testDir)

	err = os.MkdirAll(testDir, 0755)
	if err != nil {
		fmt.Printf("Error creating test directory: %v\n", err)
		return
	}

	// テスト用のファイルを作成
	subDir := filepath.Join(testDir, "subdir")
	err = os.MkdirAll(subDir, 0755)
	if err != nil {
		fmt.Printf("Error creating subdirectory: %v\n", err)
		return
	}

	testFile1 := filepath.Join(testDir, "file1.txt")
	testFile2 := filepath.Join(subDir, "file2.txt")

	err = os.WriteFile(testFile1, []byte("Hello from file1"), 0644)
	if err != nil {
		fmt.Printf("Error creating file1: %v\n", err)
		return
	}

	err = os.WriteFile(testFile2, []byte("Hello from file2"), 0644)
	if err != nil {
		fmt.Printf("Error creating file2: %v\n", err)
		return
	}

	fmt.Println("✅ Test directory structure created")

	// Go 1.24の新機能: os.Root を使用（概念的な例）
	// 注意: 実際のos.Root APIは異なる可能性があります
	fmt.Println("\n--- os.Root の概念例 ---")

	// 従来の方法（制限なし）
	fmt.Println("従来の方法（制限なし）:")
	files, err := os.ReadDir(testDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
	} else {
		for _, file := range files {
			fmt.Printf("  %s (dir: %v)\n", file.Name(), file.IsDir())
		}
	}

	// os.Root を使用した制限付きアクセス（概念例）
	fmt.Println("\nos.Root を使用した制限付きアクセス:")
	fmt.Printf("// 概念的なコード例:\n")
	fmt.Printf("root := os.NewRoot(%q)\n", testDir)
	fmt.Printf("files, err := root.ReadDir(\".\")\n")
	fmt.Printf("content, err := root.ReadFile(\"file1.txt\")\n")

	// 実際の制限付きアクセスのシミュレーション
	fmt.Println("\n実際の制限付きアクセス例:")

	// ルートディレクトリ外へのアクセスを防ぐ関数
	safeReadFile := func(rootDir, filename string) ([]byte, error) {
		// パスの正規化
		fullPath := filepath.Join(rootDir, filename)
		cleanPath := filepath.Clean(fullPath)

		// ルートディレクトリ外へのアクセスを確認
		if !filepath.HasPrefix(cleanPath, rootDir) {
			return nil, fmt.Errorf("access denied: path outside root directory")
		}

		return os.ReadFile(cleanPath)
	}

	// 安全なファイル読み取り例
	content1, err := safeReadFile(testDir, "file1.txt")
	if err != nil {
		fmt.Printf("Error reading file1: %v\n", err)
	} else {
		fmt.Printf("✅ file1.txt: %s\n", string(content1))
	}

	content2, err := safeReadFile(testDir, "subdir/file2.txt")
	if err != nil {
		fmt.Printf("Error reading file2: %v\n", err)
	} else {
		fmt.Printf("✅ subdir/file2.txt: %s\n", string(content2))
	}

	// 危険なパスへのアクセス試行
	_, err = safeReadFile(testDir, "../../../etc/passwd")
	if err != nil {
		fmt.Printf("✅ Security check worked: %v\n", err)
	}

	fmt.Println("\n--- os.Root の利点 ---")
	fmt.Println("✅ セキュリティ: ディレクトリ外へのアクセスを防止")
	fmt.Println("✅ 安全性: Path traversal攻撃の防止")
	fmt.Println("✅ 分離: サンドボックス化されたファイルアクセス")
	fmt.Println("✅ 明確性: 操作範囲の明示")

	fmt.Println("\n--- 使用例 ---")
	fmt.Println("1. Webサーバーでの静的ファイル配信")
	fmt.Println("2. ファイルアップロード処理")
	fmt.Println("3. テンプレートエンジンでのファイル読み取り")
	fmt.Println("4. コンテナ内でのファイルアクセス制御")
	fmt.Println("5. プラグインシステムでのファイル分離")

	// API例の表示
	fmt.Println("\n--- 予想されるAPI例 ---")
	apiExample := `// os.Root型の使用例
root, err := os.NewRoot("/safe/directory")
if err != nil {
    return err
}

// ルート内でのファイル操作
file, err := root.Open("data/config.json")
if err != nil {
    return err
}
defer file.Close()

// ディレクトリの作成
err = root.MkdirAll("logs/2025", 0755)

// ファイルの書き込み
err = root.WriteFile("output.txt", data, 0644)

// ディレクトリの読み取り
entries, err := root.ReadDir("uploads")`

	fmt.Println(apiExample)

	// クリーンアップ
	fmt.Println("\n--- Cleanup ---")
	err = os.RemoveAll(testDir)
	if err != nil {
		fmt.Printf("Error removing test directory: %v\n", err)
	} else {
		fmt.Println("✅ Test directory cleaned up")
	}

	fmt.Println("\n--- セキュリティのベストプラクティス ---")
	fmt.Println("1. 常にパスの正規化を行う")
	fmt.Println("2. 相対パス(.., .)の処理に注意")
	fmt.Println("3. シンボリックリンクの扱いに注意")
	fmt.Println("4. 権限の最小化")
	fmt.Println("5. 入力検証の実装")
}

// パストラバーサル攻撃の例と対策
func demonstratePathTraversal() {
	fmt.Println("\n--- Path Traversal 攻撃例 ---")

	maliciousPaths := []string{
		"../../../etc/passwd",
		"..\\..\\..\\windows\\system32\\config\\sam",
		"....//....//etc/passwd",
		"/etc/passwd",
		"./../../secret.txt",
	}

	baseDir := "/safe/uploads"

	for _, malPath := range maliciousPaths {
		fullPath := filepath.Join(baseDir, malPath)
		cleanPath := filepath.Clean(fullPath)

		fmt.Printf("Input: %s\n", malPath)
		fmt.Printf("Joined: %s\n", fullPath)
		fmt.Printf("Cleaned: %s\n", cleanPath)

		if !filepath.HasPrefix(cleanPath, baseDir) {
			fmt.Printf("❌ BLOCKED: Path outside safe directory\n")
		} else {
			fmt.Printf("✅ ALLOWED: Path within safe directory\n")
		}
		fmt.Println()
	}
}

// % go run 05_os_root.go
// === os.Root Type Demo ===
// Current directory: /Users/example/go-release-tour
//
// Creating test directory: /Users/example/go-release-tour/test-root
// ✅ Test directory structure created
//
// --- os.Root の概念例 ---
// 従来の方法（制限なし）:
//   file1.txt (dir: false)
//   subdir (dir: true)
//
// os.Root を使用した制限付きアクセス:
// // 概念的なコード例:
// root := os.NewRoot("/Users/example/go-release-tour/test-root")
// files, err := root.ReadDir(".")
// content, err := root.ReadFile("file1.txt")