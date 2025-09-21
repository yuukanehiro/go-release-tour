<div align="center">
  <img src="static/header-logo.png" alt="Go Release Tour" width="300">
</div>

# Go Release Tour

[Go Tour](https://go-tour-jp.appspot.com/)風のインターフェースで、Goの各バージョンの新機能を実際にコードを実行しながら学べる学習プラットフォームです。

## 機能

- **マルチバージョン対応**: Go 1.18〜1.25の全バージョンに対応
- **インタラクティブなコードエディター**: ブラウザ上でGoコードを編集・実行
- **真のバージョン実行**: 各Goバージョンで実際にコードを実行
- **包括的なテスト**: E2Eテスト・統合テストで品質保証
- **レスポンシブデザイン**: デスクトップ・モバイル対応

## 対応バージョンと主要機能

### Go 1.25 (最新)
- **Container-aware GOMAXPROCS** - コンテナ環境でのCPU最適化
- **Trace Flight Recorder** - 本番環境でのパフォーマンス分析
- **testing/synctest Package** - 並行処理テストの品質向上
- **go.mod ignore Directive** - 依存関係管理の向上
- **Experimental Green Tea GC** - ガベージコレクション最適化
- **encoding/json/v2 Package** - 改良されたJSON処理
- **go doc -http Option** - ローカル開発効率化

### Go 1.24
- **Generic Type Aliases** - ジェネリック型エイリアス
- **Tool Dependencies** - go.modでのツール依存関係管理
- **Swiss Tables Maps** - 高性能マップ実装
- **testing Loop** - テストループの改善
- **os.Root Type** - ファイルシステム操作の制限
- **crypto/mlkem** - ML-KEM暗号化
- **weak Package** - 弱参照の実装

### Go 1.23
- **Structured Logging** - 構造化ログ
- **Iterators** - イテレーター機能
- **Timer Reset** - タイマーリセットの改善
- **Slices Concat** - スライス結合関数
- **CMP Or** - 比較関数の拡張
- **Maps Collect** - マップ収集関数

### Go 1.22〜1.18
- **For-Range over Integers** (1.22)
- **Loop Variables** (1.22)
- **Enhanced HTTP Routing** (1.22)
- **Built-in Functions** (1.21)
- **Slices Package** (1.21)
- **Comparable Types** (1.20)
- **Generics** (1.18)
- その他多数の機能

## クイックスタート

### 前提条件
- Docker & Docker Compose
- Webブラウザ（Chrome, Firefox, Safari等）

### Make コマンド（推奨）

```bash
# リポジトリをクローン
git clone <repository-url>
cd go-release-tour

# 初期化とビルド
make init

# アプリケーション起動
make app

# ブラウザで http://localhost:8080 にアクセス
```

### 利用可能な Make コマンド

```bash
make help        # ヘルプ表示
make drop        # 全コンテナ停止・削除とクリーンアップ
make build       # Docker イメージビルド
make init        # プロジェクト初期化
make app         # アプリケーション起動（デタッチ）
make dev         # 開発環境起動（ログ表示）
make test        # 全テスト実行（E2E・統合テスト）
make logs        # アプリケーションログ表示
make status      # コンテナ状態確認
make clean       # Docker アーティファクト削除
```

### 詳細セットアップ

#### 1. プロジェクト準備
```bash
# 全て削除して最初から（必要に応じて）
make drop

# プロジェクト初期化
make init
```

#### 2. アプリケーション起動
```bash
# バックグラウンドで起動
make app

# または開発モード（ログ表示）
make dev
```

#### 3. 品質確認
```bash
# E2E・統合テスト実行
make test

# テスト結果確認
cat tests/results/e2e_test_results.json
cat tests/results/integration_test_results.txt
```

### アクセス方法
- **メインアプリ**: http://localhost:8080
- **フロントエンドテスト**: http://localhost:8080/tests/e2e_frontend_test.html

### 停止とクリーンアップ
```bash
# アプリケーション停止
docker-compose down

# 完全クリーンアップ
make drop

## 使い方

1. **バージョン選択**: 画面上部のセレクターで学習したいGoバージョンを選択
2. **レッスン選択**: 左サイドバーから学習したい機能をクリック
3. **コード編集**: 中央のエディターでコードを編集
4. **実行**: 「▶ 実行」ボタンまたは `Ctrl+Enter` で実行
5. **結果確認**: 下部に実行結果が表示（実行環境情報付き）

## 特徴

### マルチバージョン対応
- **真のバージョン実行**: 各Goバージョン環境で実際にコードを実行
- **バージョン間比較**: 同じコードを異なるバージョンで実行して違いを確認
- **自動バージョン検出**: レッスンファイルから適切なGoバージョンを自動検出

### インタラクティブな学習体験
- **自動保存**: 編集したコードは自動的にブラウザに保存
- **キーボードショートカット**:
  - `Ctrl+Enter` / `Cmd+Enter`: コード実行
  - `Ctrl+S` / `Cmd+S`: 保存確認（視覚フィードバック）
- **タブサポート**: エディター内でTabキーによるインデント

### セキュアな実行環境
- **一時ファイル**: 実行時に一時ファイルを作成し、実行後自動削除
- **セキュリティ検証**: 危険なコードパターンを事前に検出
- **サンドボックス化**: Docker環境での隔離実行

### 包括的なテスト体制
- **E2Eテスト**: 各バージョンでのAPI動作確認
- **統合テスト**: 全レッスンファイルの実行検証
- **自動テスト**: `make test`で全テストを自動実行
- **結果レポート**: JSONとテキスト形式での詳細結果

## アーキテクチャ


go-release-tour/
├── Makefile                     # Make コマンド定義
├── docker-compose.yml           # Docker Compose設定
├── docker-compose.dev.yml       # 開発用Docker Compose
├── .air.toml                    # Air設定（ホットリロード）
├── app/                         # バックエンドアプリケーション
│   ├── cmd/server/main.go       # メインサーバー
│   └── internal/                # 内部パッケージ
│       ├── config/              # 設定管理
│       ├── handlers/            # HTTPハンドラー
│       ├── lessons/             # レッスン管理
│       ├── templates/           # HTMLテンプレート
│       ├── types/               # 型定義
│       └── version/             # バージョン管理・実行
├── config/                      # アプリケーション設定
│   └── versions.json           # サポートバージョン定義
├── docker/                      # Docker関連
│   ├── Dockerfile               # 本番用Dockerfile
│   └── Dockerfile.dev           # 開発用Dockerfile
├── static/                      # フロントエンド資材
│   ├── js/
│   │   ├── components/          # JavaScript コンポーネント
│   │   └── modules/             # JavaScript モジュール
│   └── style.css                # スタイルシート
├── releases/v/                  # バージョン別レッスン
│   ├── 1.18/ ~ 1.25/           # Go 1.18〜1.25の全バージョン
│   │   ├── *.go                # レッスンファイル
│   │   └── README.md           # バージョン説明
└── tests/                       # テストスイート
    ├── e2e/                    # E2Eテスト
    ├── integration/            # 統合テスト
    └── results/                # テスト結果（gitignore対象）


### バックエンド（Go）
- **マルチバージョン実行**: Docker内の複数Goバージョンでコード実行
- **API エンドポイント**:
  - `GET /api/versions`: 利用可能バージョン一覧
  - `GET /api/lessons?version=1.24`: バージョン別レッスン一覧取得
  - `POST /api/run`: バージョン指定コード実行
- **セキュリティ**: 危険なコードパターンの事前検証
- **バージョン管理**: 自動バージョン検出とパス管理

### フロントエンド（モジュール構成）
- **Component-based**: 機能別JavaScript コンポーネント
- **Module System**: ES6モジュールでの構成
- **API Client**: 統一されたAPI通信クライアント
- **Version Management**: バージョンセレクターと状態管理

## デザイン

### カラーパレット
- **プライマリ**: `#00ADD8` (Go Blue)
- **アクセント**: `#5EC9D8` (ライトブルー)
- **成功**: `#28a745` (実行ボタン)
- **エラー**: `#fc8181` (エラー表示)

### レイアウト
- **Go Tour**にインスパイアされたデザイン
- **モダンなカードベース**レイアウト
- **シンタックスハイライト**に対応した表示

## テスト

### テスト実行
```bash
# 全テスト実行（推奨）
make test

# E2Eテストのみ
chmod +x tests/e2e/e2e_api_test.sh
./tests/e2e/e2e_api_test.sh

# 統合テストのみ
chmod +x tests/integration/test_all_lessons.sh
./tests/integration/test_all_lessons.sh
```

### テスト結果確認
```bash
# テスト結果サマリー
cat tests/results/integration_test_results.txt | tail -10

# E2E テスト詳細（JSON）
cat tests/results/e2e_test_results.json | jq '.tests[] | {name, version, status}'

# エラー詳細（統合テスト）
cat tests/results/integration_test_errors.txt
```

### テスト体制
- **E2Eテスト**: API経由での実際の実行テスト
- **統合テスト**: 全38レッスンファイル + 8基本APIテスト = 46テスト
- **品質保証**: 100%成功率での各Goバージョン動作確認

## 開発者向け

### 開発環境セットアップ

```bash
# 開発環境起動（ログ付き）
make dev

# コンテナ状態確認
make status

# ログ確認
make logs
```

### カスタマイズポイント

1. **新しいバージョン追加**:
   - `releases/v/1.XX/`ディレクトリ作成
   - `config/versions.json`にバージョン追加
2. **新しいレッスン追加**: バージョンディレクトリに`.go`ファイル追加
3. **UI変更**: `static/`ディレクトリ内のCSS/JS編集
4. **バックエンド変更**: `app/internal/`パッケージ編集
5. **設定変更**: `config/versions.json`でサポートバージョン管理

### デバッグ
```bash
# サーバーログ確認
make logs

# API直接テスト
curl -X POST http://localhost:8080/api/run \
  -H "Content-Type: application/json" \
  -d '{"code":"package main\nimport \"fmt\"\nfunc main(){fmt.Println(\"test\")}", "version":"1.25"}'

# フロントエンドテストページ
open http://localhost:8080/tests/e2e_frontend_test.html
```
