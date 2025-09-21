#!/bin/bash

# Go Release Tour E2E API Tests
# このスクリプトは各Goバージョンでの基本的なAPI動作を確認します

set -e

BASE_URL="http://localhost:8080"
TEST_DIR=$(dirname "$0")
RESULTS_DIR="$(dirname "$TEST_DIR")/results"
RESULTS_FILE="$RESULTS_DIR/e2e_test_results.json"

# 結果ディレクトリの作成
mkdir -p "$RESULTS_DIR"

echo "Go Release Tour E2E API Tests Starting..."
echo "Base URL: $BASE_URL"
echo "Test Results: $RESULTS_FILE"

# 結果格納用JSONファイル初期化
echo '{
  "test_run": "'$(date -Iseconds)'",
  "tests": []
}' > "$RESULTS_FILE"

# テスト結果を記録する関数
record_test() {
    local test_name="$1"
    local version="$2"
    local status="$3"
    local execution_time="$4"
    local error_message="$5"

    # 現在の結果を読み込み
    local temp_file=$(mktemp)
    cat "$RESULTS_FILE" | jq --arg name "$test_name" \
                            --arg version "$version" \
                            --arg status "$status" \
                            --arg exec_time "$execution_time" \
                            --arg error "$error_message" \
        '.tests += [{
            "name": $name,
            "version": $version,
            "status": $status,
            "execution_time": $exec_time,
            "error": $error,
            "timestamp": now | strftime("%Y-%m-%dT%H:%M:%SZ")
        }]' > "$temp_file"
    mv "$temp_file" "$RESULTS_FILE"
}

# APIテスト実行関数
run_api_test() {
    local test_name="$1"
    local version="$2"
    local code="$3"
    local expected_output="$4"

    echo "Testing: $test_name (Go $version)"

    # テスト用ペイロード作成
    local payload=$(jq -n \
        --arg code "$code" \
        --arg version "$version" \
        '{
            "code": $code,
            "version": $version,
            "auto_detect": false
        }')

    local start_time=$(date +%s.%3N)

    # APIリクエスト実行
    local response=$(curl -s -X POST "$BASE_URL/api/run" \
        -H "Content-Type: application/json" \
        -d "$payload" || echo '{"error": "Request failed"}')

    local end_time=$(date +%s.%3N)
    local execution_time=$(echo "$end_time - $start_time" | bc)

    # レスポンス解析
    local api_error=$(echo "$response" | jq -r '.error // empty')
    local output=$(echo "$response" | jq -r '.output // empty')
    local used_version=$(echo "$response" | jq -r '.used_version // empty')
    local go_version=$(echo "$response" | jq -r '.go_version // empty')

    if [ -n "$api_error" ]; then
        echo "FAILED: $api_error"
        record_test "$test_name" "$version" "FAILED" "${execution_time}s" "$api_error"
        return 1
    fi

    if [ "$used_version" != "$version" ]; then
        echo "FAILED: Expected version $version, got $used_version"
        record_test "$test_name" "$version" "FAILED" "${execution_time}s" "Version mismatch: expected $version, got $used_version"
        return 1
    fi

    if [ -n "$expected_output" ] && [[ "$output" != *"$expected_output"* ]]; then
        echo "FAILED: Expected output containing '$expected_output', got: $output"
        record_test "$test_name" "$version" "FAILED" "${execution_time}s" "Output mismatch"
        return 1
    fi

    echo "PASSED: Go $go_version ($used_version) - ${execution_time}s"
    record_test "$test_name" "$version" "PASSED" "${execution_time}s" ""
    return 0
}

# サーバー接続確認
echo "Checking server availability..."
if ! curl -s "$BASE_URL/" > /dev/null; then
    echo "Server not available at $BASE_URL"
    echo "Please start the server with: docker-compose -f docker-compose.dev.yml up"
    exit 1
fi
echo "Server is running"

# テスト実行開始
echo "Running E2E API Tests..."

# Go 1.25 テスト
run_api_test "Basic Hello World" "1.25" \
'package main

import "fmt"

func main() {
    fmt.Println("Hello Go 1.25!")
}' \
"Hello Go 1.25!"

# Go 1.24 テスト（Generic Type Aliases）
run_api_test "Generic Type Aliases" "1.24" \
'package main

import "fmt"

type StringList[T ~string] []T

func main() {
    var names StringList[string] = []string{"Alice", "Bob"}
    fmt.Println("Names:", names)
}' \
"Names:"

# Go 1.23 テスト（Structured Logging概念）
run_api_test "Structured Logging Concept" "1.23" \
'package main

import (
    "encoding/json"
    "fmt"
    "time"
)

func main() {
    logEntry := map[string]interface{}{
        "time":    time.Now().Format(time.RFC3339),
        "level":   "INFO",
        "message": "Go 1.23 structured logging concept",
        "version": "1.23",
    }
    jsonData, _ := json.Marshal(logEntry)
    fmt.Println(string(jsonData))
}' \
"structured logging concept"

# Go 1.22 テスト（For-Range over Integers）
run_api_test "For-Range over Integers" "1.22" \
'package main

import "fmt"

func main() {
    fmt.Println("Go 1.22 for-range over integers:")
    for i := range 3 {
        fmt.Printf("Index: %d\n", i)
    }
}' \
"Go 1.22 for-range"

# Go 1.21 テスト（slices package）
run_api_test "Slices Package" "1.21" \
'package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{3, 1, 4, 1, 5}
    slices.Sort(nums)
    fmt.Println("Sorted:", nums)
}' \
"Sorted:"

# Go 1.20 テスト（基本機能）
run_api_test "Basic Go 1.20" "1.20" \
'package main

import "fmt"

func main() {
    fmt.Println("Go 1.20 basic test")
}' \
"Go 1.20 basic test"

# Go 1.19 テスト（基本機能）
run_api_test "Basic Go 1.19" "1.19" \
'package main

import "fmt"

func main() {
    fmt.Println("Go 1.19 basic test")
}' \
"Go 1.19 basic test"

# エラーケーステスト
echo "Testing error cases..."

# 無効なバージョン
run_api_test "Invalid Version" "999.999" \
'package main

import "fmt"

func main() {
    fmt.Println("This should fail")
}' \
""

# バージョン指定なし
echo "Testing: No Version Specified"
response=$(curl -s -X POST "$BASE_URL/api/run" \
    -H "Content-Type: application/json" \
    -d '{"code":"package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"test\") }", "auto_detect": true}')

error=$(echo "$response" | jq -r '.error // empty')
if [ -n "$error" ]; then
    echo "PASSED: Correctly rejected request without version"
    record_test "No Version Specified" "none" "PASSED" "0s" ""
else
    echo "FAILED: Should have rejected request without version"
    record_test "No Version Specified" "none" "FAILED" "0s" "Should have rejected request"
fi

# 結果サマリー表示
echo ""
echo "Test Results Summary:"
total_tests=$(cat "$RESULTS_FILE" | jq '.tests | length')
passed_tests=$(cat "$RESULTS_FILE" | jq '.tests | map(select(.status == "PASSED")) | length')
failed_tests=$(cat "$RESULTS_FILE" | jq '.tests | map(select(.status == "FAILED")) | length')

echo "Total Tests: $total_tests"
echo "Passed: $passed_tests"
echo "Failed: $failed_tests"

if [ "$failed_tests" -eq 0 ]; then
    echo "All tests passed!"
    exit 0
else
    echo "Some tests failed. Check $RESULTS_FILE for details."
    exit 1
fi