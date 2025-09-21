#!/bin/bash

# Go Release Tour - API経由レッスンファイル一括テストスクリプト
# 各Goバージョンで正確にコードを実行してテストします

# 設定
API_URL="http://localhost:8080/api/run"
RESULTS_DIR="tests/results"
RESULTS_FILE="$RESULTS_DIR/integration_test_results.txt"
ERRORS_FILE="$RESULTS_DIR/integration_test_errors.txt"

# 結果ディレクトリの作成
mkdir -p "$RESULTS_DIR"

# 結果ファイルのクリア
> "$RESULTS_FILE"
> "$ERRORS_FILE"

echo "=== Go Release Tour API経由統合テスト開始 ===" | tee -a "$RESULTS_FILE"
echo "開始時刻: $(date)" | tee -a "$RESULTS_FILE"
echo "API URL: $API_URL" | tee -a "$RESULTS_FILE"
echo "" | tee -a "$RESULTS_FILE"

# カウンター
total_tests=0
success_tests=0
error_tests=0

# バージョンリスト
versions=("1.25" "1.24" "1.23" "1.22" "1.21" "1.20" "1.19" "1.18")

# APIテスト関数
test_api_with_code() {
    local version="$1"
    local code="$2"
    local test_name="$3"

    total_tests=$((total_tests + 1))
    echo -n "  API テスト: $test_name (Go $version) ... " | tee -a "$RESULTS_FILE"

    # APIリクエストのペイロード作成
    local payload=$(jq -n \
        --arg code "$code" \
        --arg version "$version" \
        '{code: $code, version: $version}')

    # API呼び出し
    local response=$(curl -s -X POST "$API_URL" \
        -H "Content-Type: application/json" \
        -d "$payload" \
        --max-time 30)

    # レスポンスの確認
    if [ $? -eq 0 ] && echo "$response" | jq -e '.error == null or (.error | length) == 0' > /dev/null 2>&1; then
        echo "[PASS]" | tee -a "$RESULTS_FILE"
        success_tests=$((success_tests + 1))

        # 成功時の出力を保存
        local output=$(echo "$response" | jq -r '.output // "No output"')
        local used_version=$(echo "$response" | jq -r '.used_version // "Unknown"')
        echo "    実行バージョン: $used_version" >> "$RESULTS_FILE"
        echo "    出力: $(echo "$output" | head -c 100)..." >> "$RESULTS_FILE"
    else
        echo "[FAIL]" | tee -a "$RESULTS_FILE" "$ERRORS_FILE"
        error_tests=$((error_tests + 1))

        # エラー詳細を保存
        echo "テスト: $test_name (Go $version)" >> "$ERRORS_FILE"
        echo "レスポンス: $response" >> "$ERRORS_FILE"
        echo "---" >> "$ERRORS_FILE"
    fi
}

# 各バージョンのテスト実行
for version in "${versions[@]}"; do
    echo "=== Go $version のテスト ===" | tee -a "$RESULTS_FILE"
    version_dir="releases/v/$version"

    if [ ! -d "$version_dir" ]; then
        echo "[WARN] ディレクトリが存在しません: $version_dir" | tee -a "$RESULTS_FILE"
        continue
    fi

    # .goファイルを検索
    go_files=$(find "$version_dir" -name "*.go" -type f)

    if [ -z "$go_files" ]; then
        echo "[WARN] .goファイルが見つかりません: $version_dir" | tee -a "$RESULTS_FILE"
        continue
    fi

    echo "テストファイル数: $(echo "$go_files" | wc -l)" | tee -a "$RESULTS_FILE"

    # 各ファイルをテスト
    while IFS= read -r file; do
        if [ -n "$file" ]; then
            filename=$(basename "$file")

            # ファイルの内容を読み取り
            if [ -r "$file" ]; then
                file_content=$(cat "$file")
                test_api_with_code "$version" "$file_content" "$filename"
            else
                echo "  [SKIP] ファイルが読み取れません: $filename" | tee -a "$RESULTS_FILE"
            fi
        fi
    done <<< "$go_files"

    echo "" | tee -a "$RESULTS_FILE"
done

# 基本的なAPIテストも追加
echo "=== 基本APIテスト ===" | tee -a "$RESULTS_FILE"

# 基本的なHello Worldテスト
hello_world='package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}'

for version in "${versions[@]}"; do
    test_api_with_code "$version" "$hello_world" "Basic Hello World"
done

echo "" | tee -a "$RESULTS_FILE"

# サマリー
echo "=== テスト結果サマリー ===" | tee -a "$RESULTS_FILE"
echo "総テスト数: $total_tests" | tee -a "$RESULTS_FILE"
echo "成功: $success_tests" | tee -a "$RESULTS_FILE"
echo "エラー: $error_tests" | tee -a "$RESULTS_FILE"

if [ $total_tests -gt 0 ]; then
    success_rate=$(echo "scale=1; $success_tests * 100 / $total_tests" | bc -l 2>/dev/null || echo "N/A")
    echo "成功率: $success_rate%" | tee -a "$RESULTS_FILE"
fi

echo "" | tee -a "$RESULTS_FILE"
echo "終了時刻: $(date)" | tee -a "$RESULTS_FILE"

echo ""
echo "=== 実行完了 ==="
echo "詳細結果: $RESULTS_FILE"
echo "エラー詳細: $ERRORS_FILE"

if [ $error_tests -gt 0 ]; then
    echo ""
    echo "[ERROR] エラーのあるテスト一覧:"
    grep "テスト:" "$ERRORS_FILE"
    exit 1
else
    echo ""
    echo "[SUCCESS] 全てのテストが成功しました！"
    exit 0
fi