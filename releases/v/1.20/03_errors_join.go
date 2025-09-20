// Go 1.20 新機能: errors.Join Function
// 原文: "The new function errors.Join returns an error wrapping a list of errors"
//
// 説明: Go 1.20では、複数のエラーをまとめる errors.Join 関数が追加され、
// より柔軟なエラーハンドリングが可能になりました。
//
// 参考リンク:
// - Go 1.20 Release Notes: https://go.dev/doc/go1.20#errors
// - errors Package: https://pkg.go.dev/errors

//go:build ignore
// +build ignore

package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// カスタムエラー型の定義
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("検証エラー [%s]: %s", e.Field, e.Message)
}

// Unwrap() []error メソッドを実装したカスタムエラー型
type MultiError struct {
	Errors []error
}

func (me MultiError) Error() string {
	if len(me.Errors) == 0 {
		return "エラーなし"
	}
	if len(me.Errors) == 1 {
		return me.Errors[0].Error()
	}
	return fmt.Sprintf("%s (他 %d 個のエラー)", me.Errors[0].Error(), len(me.Errors)-1)
}

func (me MultiError) Unwrap() []error {
	return me.Errors
}

func main() {
	fmt.Println("=== errors.Join Function Demo ===")

	// 基本的な使用例
	fmt.Println("\n--- 基本的な使用例 ---")

	err1 := errors.New("データベース接続エラー")
	err2 := errors.New("キャッシュ接続エラー")
	err3 := errors.New("外部API接続エラー")

	// 複数のエラーを結合
	combinedErr := errors.Join(err1, err2, err3)
	fmt.Printf("結合されたエラー:\n%v\n", combinedErr)

	fmt.Println("\n--- errors.Is の使用 ---")

	// 特定のエラーが含まれているかチェック
	if errors.Is(combinedErr, err1) {
		fmt.Println("✅ データベース接続エラーが含まれています")
	}

	if errors.Is(combinedErr, err2) {
		fmt.Println("✅ キャッシュ接続エラーが含まれています")
	}

	// 含まれていないエラーのチェック
	notIncluded := errors.New("含まれていないエラー")
	if !errors.Is(combinedErr, notIncluded) {
		fmt.Println("✅ 含まれていないエラーは正しく検出されません")
	}

	fmt.Println("\n--- errors.As の使用 ---")

	// 複数の検証エラーを作成
	valErr1 := ValidationError{Field: "email", Message: "無効な形式"}
	valErr2 := ValidationError{Field: "password", Message: "短すぎます"}
	valErr3 := errors.New("一般的なエラー")

	validationErr := errors.Join(valErr1, valErr2, valErr3)
	fmt.Printf("検証エラー:\n%v\n", validationErr)

	// errors.As を使用して特定の型のエラーを取得
	var ve ValidationError
	if errors.As(validationErr, &ve) {
		fmt.Printf("✅ 検証エラーが見つかりました: %s\n", ve.Field)
	}

	fmt.Println("\n--- 実用的な例: 並行処理でのエラー収集 ---")

	// 複数の処理を模擬実行
	results := processMultipleTasks()
	if results != nil {
		fmt.Printf("処理中にエラーが発生しました:\n%v\n", results)

		// 個別のエラーを確認
		if errors.Is(results, io.EOF) {
			fmt.Println("  ➤ EOF エラーが含まれています")
		}
		if errors.Is(results, strconv.ErrSyntax) {
			fmt.Println("  ➤ 構文エラーが含まれています")
		}
	}

	fmt.Println("\n--- エラーのアンラップ ---")

	// カスタムマルチエラーの使用
	multiErr := MultiError{
		Errors: []error{
			errors.New("エラー1"),
			errors.New("エラー2"),
			errors.New("エラー3"),
		},
	}

	fmt.Printf("カスタムマルチエラー: %v\n", multiErr)

	// errors.Is でチェック
	target := errors.New("エラー2")
	if errors.Is(multiErr, target) {
		fmt.Println("✅ エラー2が含まれています")
	}

	fmt.Println("\n--- nilエラーの処理 ---")

	// nil エラーを含む場合
	errWithNil := errors.Join(err1, nil, err2, nil)
	fmt.Printf("nil含むエラー結合: %v\n", errWithNil)

	// 全てnil の場合
	allNil := errors.Join(nil, nil, nil)
	fmt.Printf("全てnil: %v\n", allNil) // nil が返される

	fmt.Println("\n--- 実際のアプリケーションでの使用例 ---")

	// Webアプリケーションでの検証例
	user := User{
		Email:    "invalid-email",
		Password: "123",
		Age:      -5,
	}

	if err := validateUser(user); err != nil {
		fmt.Printf("ユーザー検証エラー:\n%v\n", err)

		// 特定のエラータイプをチェック
		var fieldErr FieldError
		if errors.As(err, &fieldErr) {
			fmt.Printf("フィールドエラーが見つかりました: %s\n", fieldErr.Field)
		}
	}

	fmt.Println("\n利点:")
	fmt.Println("  - 複数のエラーを1つにまとめられる")
	fmt.Println("  - errors.Is/As が複数エラーに対応")
	fmt.Println("  - 並行処理でのエラー収集が簡単")
	fmt.Println("  - エラーの詳細な分析が可能")
	fmt.Println("  - 既存のエラーハンドリングパターンと互換")
}

// 並行処理のシミュレーション
func processMultipleTasks() error {
	var errs []error

	// タスク1: ファイル読み込みエラー
	errs = append(errs, io.EOF)

	// タスク2: 数値変換エラー
	_, err := strconv.Atoi("invalid")
	if err != nil {
		errs = append(errs, err)
	}

	// タスク3: 成功（エラーなし）
	// errs = append(errs, nil) // nilは追加しない

	// タスク4: カスタムエラー
	errs = append(errs, errors.New("外部サービス応答なし"))

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// 検証用の構造体とエラー型
type User struct {
	Email    string
	Password string
	Age      int
}

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateUser(user User) error {
	var errs []error

	// メール検証
	if !isValidEmail(user.Email) {
		errs = append(errs, FieldError{
			Field:   "email",
			Message: "無効なメールアドレス形式",
		})
	}

	// パスワード検証
	if len(user.Password) < 8 {
		errs = append(errs, FieldError{
			Field:   "password",
			Message: "パスワードは8文字以上である必要があります",
		})
	}

	// 年齢検証
	if user.Age < 0 {
		errs = append(errs, FieldError{
			Field:   "age",
			Message: "年齢は0以上である必要があります",
		})
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func isValidEmail(email string) bool {
	// 簡単なメール検証（実用的ではない）
	return len(email) > 0 && email[0] != '@' && email[len(email)-1] != '@'
}

// ファイル処理でのエラー収集例
func processFiles(filenames []string) error {
	var errs []error

	for _, filename := range filenames {
		if err := processFile(filename); err != nil {
			errs = append(errs, fmt.Errorf("ファイル %s: %w", filename, err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func processFile(filename string) error {
	// ファイル処理のシミュレーション
	if filename == "nonexistent.txt" {
		return errors.New("ファイルが見つかりません")
	}
	if filename == "readonly.txt" {
		return errors.New("読み取り専用ファイル")
	}
	return nil
}

// % go run 03_errors_join.go
// === errors.Join Function Demo ===
//
// --- 基本的な使用例 ---
// 結合されたエラー:
// データベース接続エラー
// キャッシュ接続エラー
// 外部API接続エラー
//
// --- errors.Is の使用 ---
// ✅ データベース接続エラーが含まれています
// ✅ キャッシュ接続エラーが含まれています
// ✅ 含まれていないエラーは正しく検出されません