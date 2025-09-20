// Go 1.23 新機能: cmp.Or Function
// 原文: "New cmp.Or function returns the first non-zero value from its arguments"
//
// 説明: Go 1.23では、cmp.Or関数が追加され、
// 引数から最初の非ゼロ値を返すことでデフォルト値の設定が簡単になりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"
)

// cmp.Or のシミュレーション関数
func orString(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func orInt(values ...int) int {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}

func orBool(values ...bool) bool {
	for _, v := range values {
		if v != false {
			return v
		}
	}
	return false
}

func orDuration(values ...time.Duration) time.Duration {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}

func main() {
	fmt.Println("=== cmp.Or Function Demo ===")
	fmt.Println("注意: この例はcmp.Orの概念を示すもので、実際の環境では互換性のある実装を使用しています。")

	// 基本的な使用例
	demonstrateBasicUsage()

	// 文字列でのデフォルト値設定
	demonstrateStringDefaults()

	// 数値でのデフォルト値設定
	demonstrateNumericDefaults()

	// 構造体での使用例
	demonstrateStructUsage()

	// 実用的なパターン
	demonstratePracticalPatterns()
}

func demonstrateBasicUsage() {
	fmt.Println("\n--- 基本的な使用例 ---")

	// 文字列のデフォルト値
	var name string // ゼロ値（空文字列）
	defaultName := orString(name, "匿名ユーザー")
	fmt.Printf("名前: '%s' → デフォルト適用: '%s'\n", name, defaultName)

	name = "田中太郎"
	actualName := orString(name, "匿名ユーザー")
	fmt.Printf("名前: '%s' → 値使用: '%s'\n", name, actualName)

	// 数値のデフォルト値
	var port int // ゼロ値（0）
	defaultPort := orInt(port, 8080)
	fmt.Printf("ポート: %d → デフォルト適用: %d\n", port, defaultPort)

	port = 3000
	actualPort := orInt(port, 8080)
	fmt.Printf("ポート: %d → 値使用: %d\n", port, actualPort)
}

func demonstrateStringDefaults() {
	fmt.Println("\n--- 文字列でのデフォルト値設定 ---")

	config := struct {
		Host     string
		Database string
		Username string
	}{
		Host:     "", // 未設定
		Database: "myapp",
		Username: "", // 未設定
	}

	// 複数のデフォルト値を適用
	finalHost := orString(config.Host, "localhost")
	finalDB := orString(config.Database, "default")
	finalUser := orString(config.Username, "admin")

	fmt.Printf("Host: '%s' → '%s'\n", config.Host, finalHost)
	fmt.Printf("Database: '%s' → '%s'\n", config.Database, finalDB)
	fmt.Printf("Username: '%s' → '%s'\n", config.Username, finalUser)

	// 環境変数パターンのシミュレーション
	fmt.Println("\n環境変数パターン:")
	getEnv := func(key string) string {
		envs := map[string]string{
			"APP_PORT": "9000",
			// "APP_HOST" は未設定
		}
		return envs[key]
	}

	appHost := orString(getEnv("APP_HOST"), "0.0.0.0")
	appPort := orString(getEnv("APP_PORT"), "8080")

	fmt.Printf("APP_HOST: %s\n", appHost)
	fmt.Printf("APP_PORT: %s\n", appPort)
}

func demonstrateNumericDefaults() {
	fmt.Println("\n--- 数値でのデフォルト値設定 ---")

	settings := struct {
		Timeout     int
		MaxRetries  int
		BufferSize  int
		MaxWorkers  int
	}{
		Timeout:    0, // 未設定
		MaxRetries: 3,
		BufferSize: 0, // 未設定
		MaxWorkers: 0, // 未設定
	}

	// デフォルト値の適用
	finalTimeout := orInt(settings.Timeout, 30)
	finalRetries := orInt(settings.MaxRetries, 5)
	finalBuffer := orInt(settings.BufferSize, 1024)
	finalWorkers := orInt(settings.MaxWorkers, 4)

	fmt.Printf("Timeout: %d秒 → %d秒\n", settings.Timeout, finalTimeout)
	fmt.Printf("MaxRetries: %d → %d\n", settings.MaxRetries, finalRetries)
	fmt.Printf("BufferSize: %d → %d\n", settings.BufferSize, finalBuffer)
	fmt.Printf("MaxWorkers: %d → %d\n", settings.MaxWorkers, finalWorkers)

	// 複数候補からの選択
	fmt.Println("\n複数候補からの選択:")
	userLimit := 0      // ユーザー設定なし
	orgLimit := 100     // 組織設定
	systemLimit := 1000 // システムデフォルト

	effectiveLimit := orInt(userLimit, orgLimit, systemLimit)
	fmt.Printf("有効な制限値: %d (ユーザー:%d, 組織:%d, システム:%d)\n",
		effectiveLimit, userLimit, orgLimit, systemLimit)
}

func demonstrateStructUsage() {
	fmt.Println("\n--- 構造体での使用例 ---")

	type Config struct {
		ServerName string
		Port       int
		Debug      bool
		Timeout    time.Duration
	}

	// ユーザー設定（一部のみ指定）
	userConfig := Config{
		Port: 3000,
		// ServerName, Debug, Timeout は未設定（ゼロ値）
	}

	// デフォルト設定
	defaultConfig := Config{
		ServerName: "MyApp",
		Port:       8080,
		Debug:      false,
		Timeout:    30 * time.Second,
	}

	// 設定をマージ
	finalConfig := Config{
		ServerName: orString(userConfig.ServerName, defaultConfig.ServerName),
		Port:       orInt(userConfig.Port, defaultConfig.Port),
		Debug:      orBool(userConfig.Debug, defaultConfig.Debug),
		Timeout:    orDuration(userConfig.Timeout, defaultConfig.Timeout),
	}

	fmt.Printf("最終設定:\n")
	fmt.Printf("  ServerName: %s\n", finalConfig.ServerName)
	fmt.Printf("  Port: %d\n", finalConfig.Port)
	fmt.Printf("  Debug: %v\n", finalConfig.Debug)
	fmt.Printf("  Timeout: %v\n", finalConfig.Timeout)
}

func demonstratePracticalPatterns() {
	fmt.Println("\n--- 実用的なパターン ---")

	// ログレベルの設定
	fmt.Println("1. ログレベル設定:")
	var logLevel string // 環境変数から取得（空の場合）
	effectiveLogLevel := orString(logLevel, "INFO")
	fmt.Printf("   ログレベル: %s\n", effectiveLogLevel)

	// ファイルパスの設定
	fmt.Println("\n2. ファイルパス設定:")
	var configPath string // コマンドライン引数（指定されない場合）
	var homeConfigPath = "/home/user/.config/app.yaml"
	var systemConfigPath = "/etc/app/config.yaml"

	finalConfigPath := orString(configPath, homeConfigPath, systemConfigPath)
	fmt.Printf("   設定ファイル: %s\n", finalConfigPath)

	// API応答での使用
	fmt.Println("\n3. API応答での使用:")
	type User struct {
		Name     string
		Email    string
		Language string
	}

	user := User{
		Name:  "", // 未設定
		Email: "user@example.com",
		// Language 未設定
	}

	// レスポンス用に整形
	response := map[string]string{
		"name":     orString(user.Name, "匿名ユーザー"),
		"email":    orString(user.Email, "未設定"),
		"language": orString(user.Language, "ja"),
	}

	fmt.Printf("   APIレスポンス: %+v\n", response)

	// チェーン的な使用
	fmt.Println("\n4. チェーン的な設定解決:")
	var cmdLineValue string
	var envValue string = "env-setting"
	var configValue string
	var defaultValue string = "default-value"

	result := orString(cmdLineValue, envValue, configValue, defaultValue)
	fmt.Printf("   解決された値: %s\n", result)
	fmt.Printf("   解決順序: コマンドライン → 環境変数 → 設定ファイル → デフォルト\n")
}

// 従来の方法との比較
func demonstrateComparison() {
	fmt.Println("\n--- 従来の方法との比較 ---")

	var username string

	// 従来の方法
	var traditional string
	if username != "" {
		traditional = username
	} else {
		traditional = "guest"
	}

	// orStringを使用
	modern := orString(username, "guest")

	fmt.Printf("従来の方法: %s\n", traditional)
	fmt.Printf("orString: %s\n", modern)

	// より複雑な従来の方法
	var value1, value2, value3 string
	value2 = "second"

	// 従来（三項演算子がないGoでは冗長）
	var traditionalComplex string
	if value1 != "" {
		traditionalComplex = value1
	} else if value2 != "" {
		traditionalComplex = value2
	} else if value3 != "" {
		traditionalComplex = value3
	} else {
		traditionalComplex = "default"
	}

	// orStringでスッキリ
	modernComplex := orString(value1, value2, value3, "default")

	fmt.Printf("\n複雑なケース:\n")
	fmt.Printf("従来: %s\n", traditionalComplex)
	fmt.Printf("orString: %s\n", modernComplex)
}

// カスタム型での使用
type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "低"
	case Medium:
		return "中"
	case High:
		return "高"
	default:
		return "未設定"
	}
}

func demonstrateCustomTypes() {
	fmt.Println("\n--- カスタム型での使用 ---")

	var userPriority Priority // ゼロ値（0）
	var defaultPriority Priority
	if userPriority != 0 {
		defaultPriority = userPriority
	} else {
		defaultPriority = Medium
	}

	fmt.Printf("ユーザー優先度: %v → デフォルト適用: %v\n", userPriority, defaultPriority)

	userPriority = High
	var actualPriority Priority
	if userPriority != 0 {
		actualPriority = userPriority
	} else {
		actualPriority = Medium
	}
	fmt.Printf("ユーザー優先度: %v → 値使用: %v\n", userPriority, actualPriority)
}

// % go run 05_cmp_or.go
// === cmp.Or Function Demo ===
//
// --- 基本的な使用例 ---
// 名前: '' → デフォルト適用: '匿名ユーザー'
// 名前: '田中太郎' → 値使用: '田中太郎'
// ポート: 0 → デフォルト適用: 8080
// ポート: 3000 → 値使用: 3000
//
// --- 文字列でのデフォルト値設定 ---
// Host: '' → 'localhost'
// Database: 'myapp' → 'myapp'
// Username: '' → 'admin'
//
// 環境変数パターン:
// APP_HOST: 0.0.0.0
// APP_PORT: 9000