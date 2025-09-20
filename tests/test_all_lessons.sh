#!/bin/bash

# シェル設定を読み込み
source ~/.zshrc

# Go Release Tour レッスンファイル一括テストスクリプト

# 結果ファイルのクリア
> test_results.txt
> test_errors.txt

echo "=== Go Release Tour レッスンファイル一括テスト開始 ===" | tee -a test_results.txt
echo "開始時刻: $(date)" | tee -a test_results.txt
echo "" | tee -a test_results.txt

# カウンター
total_files=0
success_files=0
error_files=0

# バージョンリスト
versions=("1.25" "1.24" "1.23" "1.22" "1.21" "1.20" "1.19" "1.18")

for version in "${versions[@]}"; do
    echo "=== Go $version のテスト ===" | tee -a test_results.txt
    version_dir="releases/v/$version"

    if [ ! -d "$version_dir" ]; then
        echo "❌ ディレクトリが存在しません: $version_dir" | tee -a test_results.txt test_errors.txt
        continue
    fi

    # .goファイルを検索
    go_files=$(find "$version_dir" -name "*.go" -type f)

    if [ -z "$go_files" ]; then
        echo "⚠️  .goファイルが見つかりません: $version_dir" | tee -a test_results.txt
        continue
    fi

    echo "テストファイル数: $(echo "$go_files" | wc -l)" | tee -a test_results.txt

    # 各ファイルをテスト
    while IFS= read -r file; do
        if [ -n "$file" ]; then
            total_files=$((total_files + 1))
            filename=$(basename "$file")
            echo -n "  テスト中: $filename ... " | tee -a test_results.txt

            # goを実行（timeoutコマンドはmacOSにないため直接実行）
            if go run "$file" > /tmp/go_output_$$ 2>&1; then
                echo "✅ 成功" | tee -a test_results.txt
                success_files=$((success_files + 1))
                # 成功時の出力の最初の数行を保存
                echo "    出力:" >> test_results.txt
                head -n 3 /tmp/go_output_$$ | sed 's/^/      /' >> test_results.txt
            else
                echo "❌ エラー" | tee -a test_results.txt test_errors.txt
                error_files=$((error_files + 1))
                echo "ファイル: $file" >> test_errors.txt
                echo "エラー内容:" >> test_errors.txt
                cat /tmp/go_output_$$ | sed 's/^/  /' >> test_errors.txt
                echo "---" >> test_errors.txt
            fi

            # 一時ファイルを削除
            rm -f /tmp/go_output_$$
        fi
    done <<< "$go_files"

    echo "" | tee -a test_results.txt
done

# サマリー
echo "=== テスト結果サマリー ===" | tee -a test_results.txt
echo "総ファイル数: $total_files" | tee -a test_results.txt
echo "成功: $success_files" | tee -a test_results.txt
echo "エラー: $error_files" | tee -a test_results.txt
echo "成功率: $(echo "scale=1; $success_files * 100 / $total_files" | bc -l)%" | tee -a test_results.txt
echo "" | tee -a test_results.txt
echo "終了時刻: $(date)" | tee -a test_results.txt

echo ""
echo "=== 実行完了 ==="
echo "詳細結果: test_results.txt"
echo "エラー詳細: test_errors.txt"

if [ $error_files -gt 0 ]; then
    echo ""
    echo "❌ エラーのあるファイル一覧:"
    grep "ファイル:" test_errors.txt
    exit 1
else
    echo ""
    echo "✅ 全てのテストが成功しました！"
    exit 0
fi