#!/bin/bash

# Go Release Tour - API経由レッスンファイル一括テストスクリプト
# Go言語で書き直されたテストランナーのラッパー

set -e

# デフォルト設定
API_URL="${API_URL:-http://localhost:8080/api/run}"
OUTPUT_DIR="${OUTPUT_DIR:-../results}"
VERBOSE="${VERBOSE:-false}"

# Goテストランナーのディレクトリに移動
cd "$(dirname "$0")"

# Go依存関係の確認とダウンロード
if [ ! -f "go.mod" ]; then
    echo "Error: go.mod not found. Please ensure you're in the correct directory."
    exit 1
fi

# Goプログラムをビルドして実行
echo "Building Go integration test runner..."
go build -o integration-test-runner .

echo "Running Go integration tests..."
if [ "$VERBOSE" = "true" ]; then
    ./integration-test-runner -url="$API_URL" -output="$OUTPUT_DIR" -v
else
    ./integration-test-runner -url="$API_URL" -output="$OUTPUT_DIR"
fi

# 終了コードを保持
exit_code=$?

# クリーンアップ
rm -f integration-test-runner

exit $exit_code