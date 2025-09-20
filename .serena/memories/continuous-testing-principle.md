# 継続的テスト実行の原則

## 基本原則
**プログラムを書いたら一定間隔で全て動かしてデバッグすること**

## 実行方法
1. **直接実行可能な場合**: そのまま実行
2. **直接実行できない場合**: スクリプトなどを通して実行

## 具体的な実装パターン

### 1. 一括テストスクリプトの活用
```bash
#!/bin/bash
# プロジェクト全体のテストスクリプト例

source ~/.zshrc  # 環境設定の読み込み

# 全ファイルのテスト
find . -name "*.go" | while read file; do
    echo "テスト中: $file"
    if go run "$file" > /dev/null 2>&1; then
        echo "✅ 成功: $file"
    else
        echo "❌ エラー: $file"
        go run "$file"  # エラー詳細を表示
    fi
done
```

### 2. 定期実行のタイミング
- **新機能追加後**
- **バグ修正後**
- **リファクタリング後**
- **依存関係変更後**
- **コミット前**

### 3. 実行できないファイルへの対応

#### A. build制約のあるファイル
```go
//go:build ignore
// +build ignore
```
→ `go run` で直接実行

#### B. パッケージとしてのみ動作するファイル
→ テスト用の main関数を含むラッパーファイル作成

#### C. 特定環境依存のファイル
→ Docker等を使用した環境構築スクリプト

### 4. 自動化の例

#### Makefileを使用
```makefile
.PHONY: test-all
test-all:
	@echo "全ファイルテスト開始"
	@./test_all_lessons.sh
	@echo "テスト完了"

.PHONY: test-quick
test-quick:
	@echo "クイックテスト"
	@find . -name "*.go" -exec go build {} \;
```

#### Git hooksの活用
```bash
# .git/hooks/pre-commit
#!/bin/bash
echo "コミット前テスト実行..."
./test_all_lessons.sh
if [ $? -ne 0 ]; then
    echo "テストに失敗しました。コミットを中止します。"
    exit 1
fi
```

## 重要なポイント

### 1. 早期発見の価値
- 小さなエラーは早く見つけて修正
- 複数の変更が絡む前に問題を特定
- デバッグ時間の大幅短縮

### 2. 全体への影響確認
- 一部の修正が他の部分に与える影響をチェック
- 依存関係のあるファイル群の整合性確認
- リグレッション（退行バグ）の防止

### 3. 開発リズムの確立
- 「書く→テスト→修正」のサイクル確立
- 継続的な品質維持
- 安心してリファクタリングできる環境構築

## 実践例（Go Release Tourプロジェクト）

### 使用したテストスクリプト
```bash
#!/bin/bash
source ~/.zshrc

# 結果ファイルのクリア
> test_results.txt
> test_errors.txt

total_files=0
success_files=0

# 各バージョンをテスト
for version in 1.25 1.24 1.23 1.22 1.21 1.20 1.19 1.18; do
    echo "=== Go $version のテスト ==="

    while IFS= read -r file; do
        if [ -n "$file" ]; then
            total_files=$((total_files + 1))
            filename=$(basename "$file")
            echo -n "  テスト中: $filename ... "

            if go run "$file" > /tmp/go_output_$$ 2>&1; then
                echo "✅ 成功"
                success_files=$((success_files + 1))
            else
                echo "❌ エラー"
                echo "ファイル: $file" >> test_errors.txt
                cat /tmp/go_output_$$ >> test_errors.txt
            fi
            rm -f /tmp/go_output_$$
        fi
    done < <(find "releases/v/$version" -name "*.go" 2>/dev/null | sort)
done

# 結果サマリー
echo "=== テスト結果サマリー ==="
echo "総ファイル数: $total_files"
echo "成功: $success_files"
echo "エラー: $((total_files - success_files))"
success_rate=$(awk "BEGIN {printf \"%.1f\", $success_files/$total_files*100}")
echo "成功率: $success_rate%"
```

### 成果
- 初期: 0% → 最終: 100%
- 系統的なエラー分類と修正
- 全30ファイルの動作保証

## 注意事項
- **パフォーマンスとのバランス**: 大規模プロジェクトでは部分テストも併用
- **環境依存への対応**: CI/CD環境でも同様のテストが実行できるよう設計
- **テスト結果の保存**: ログファイルでエラーの傾向を分析

## まとめ
継続的なテスト実行により、品質の高いコードを効率的に維持できる。特に複数ファイルが相互に関連するプロジェクトでは、全体テストの価値が非常に高い。