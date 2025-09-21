# Go Release Tour - Testing

このディレクトリには、Go Release Tourの包括的なテストスイートが含まれています。

## ディレクトリ構造

```
tests/
├── README.md               # このファイル
├── e2e/                   # End-to-End テスト
│   ├── e2e_api_test.sh    # APIテスト
│   └── e2e_frontend_test.html # フロントエンドテスト
├── integration/           # 統合テスト
│   └── test_all_lessons.sh # 全レッスンテスト
├── unit/                  # ユニットテスト（将来追加予定）
└── results/               # テスト結果ファイル
    ├── e2e_test_results.json
    ├── test_results.txt
    └── test_errors.txt
```

## テスト種別

### 1. End-to-End テスト (`e2e/`)

#### APIテスト (`e2e/e2e_api_test.sh`)

**概要**: バックエンドAPIの各Goバージョンでの実行をテストする自動化スクリプト

**実行方法**:
```bash
# 全テスト実行（推奨）
make test

# 個別実行
chmod +x tests/e2e/e2e_api_test.sh
./tests/e2e/e2e_api_test.sh
```

**テスト内容**:
- Go 1.25: 基本的なHello World
- Go 1.24: Generic Type Aliases機能
- Go 1.23: Structured Logging概念
- Go 1.22: For-Range over Integers
- Go 1.21: slicesパッケージ
- Go 1.20, 1.19: 基本機能（互換性確認）
- エラーケース: 無効なバージョン、バージョン未指定

**結果確認**:
```bash
# JSON形式の詳細結果
cat tests/e2e_test_results.json | jq .

# 結果サマリー
cat tests/e2e_test_results.json | jq '.tests[] | {name, version, status}'
```

### 2. フロントエンドテスト (`e2e_frontend_test.html`)

**概要**: ブラウザでの実際のフロントエンド動作をテストするインタラクティブページ

**アクセス方法**:
```
http://localhost:8080/tests/e2e_frontend_test.html
```

**テスト機能**:
- 手動コード実行テスト
- バージョンセレクター動作確認
- 自動化された全バージョンテスト
- ブラウザコンソールでのデバッグログ確認

**使用方法**:
1. バージョンセレクターで任意のGoバージョンを選択
2. コードエディターにGoコードを入力
3. 「実行」ボタンでテスト
4. 「全バージョンテスト実行」で自動テスト

## テスト結果例

### APIテスト成功例
```json
{
  "test_run": "2025-09-21T19:22:48+09:00",
  "tests": [
    {
      "name": "Basic Hello World",
      "version": "1.25",
      "status": "PASSED",
      "execution_time": "2.00s",
      "timestamp": "2025-09-21T10:22:50Z"
    },
    {
      "name": "For-Range over Integers",
      "version": "1.22",
      "status": "PASSED",
      "execution_time": "0.80s",
      "timestamp": "2025-09-21T10:22:52Z"
    }
  ]
}
```

### 期待される結果

**成功するバージョン**:
- Go 1.25: 最新機能
- Go 1.24: Generic Type Aliases
- Go 1.23: Structured Logging
- Go 1.22: For-Range over Integers
- Go 1.21: slicesパッケージ

**予期される制限**:
- Go 1.18-1.20: Docker環境でのRosetta互換性問題（既知の制限）

## デバッグ

### ブラウザでのデバッグ
1. フロントエンドテストページでF12を開く
2. コンソールでデバッグログを確認:
   ```
   Debug: versionSelect element = <select>
   Debug: selectedVersion = 1.25
   Debug: Final payload = {"code":"...", "version":"1.25", "auto_detect":false}
   ```

### APIデバッグ
```bash
# 直接APIをテスト
curl -X POST http://localhost:8080/api/run \
  -H "Content-Type: application/json" \
  -d '{"code":"package main\nimport \"fmt\"\nfunc main(){fmt.Println(\"test\")}", "version":"1.25", "auto_detect":false}'
```

### サーバーログ確認
```bash
docker-compose -f docker-compose.dev.yml logs go-release-tour-dev
```

## CI/CD統合

### GitHub Actions例
```yaml
- name: Run E2E Tests
  run: |
    docker-compose -f docker-compose.dev.yml up -d
    sleep 10
    ./tests/e2e_api_test.sh
    docker-compose -f docker-compose.dev.yml down
```

### 継続的監視
- 新しいGoバージョンリリース時の対応確認
- パフォーマンス回帰テスト
- 各バージョンでの機能動作確認

## テスト追加ガイド

新しいテストケースを追加する場合:

1. **APIテスト**: `e2e_api_test.sh`の`run_api_test`呼び出しを追加
2. **フロントエンドテスト**: `e2e_frontend_test.html`の`testCases`配列に追加
3. **期待される動作**: 各バージョンで正確に実行されることを確認

## 品質保証

これらのE2Eテストにより以下が保証されます:

- 真のマルチバージョン実行（欺瞞なし）
- バージョンセレクターの正確な動作
- フロントエンド・バックエンド連携
- 各Goバージョンの新機能が適切に動作
- エラーハンドリングの適切性
- パフォーマンスの確認（実行時間測定）