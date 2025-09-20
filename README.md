# Go Release Tour

**Go 1.25の新機能をインタラクティブに学習できるWebアプリケーション**

[Go Tour](https://go-tour-jp.appspot.com/)風のインターフェースで、Go 1.25の新機能を実際にコードを実行しながら学べます。

## 機能

- **インタラクティブなコードエディター**: ブラウザ上でGoコードを編集・実行
- **7つの新機能を網羅**: Go 1.25の主要な新機能をカバー
- **リアルタイム実行**: コードをその場で実行して結果を確認
- **実用性評価**: 各機能の実用性を5段階で表示
- **レスポンシブデザイン**: デスクトップ・モバイル対応

## 学習できる新機能

### 超重要
1. **Container-aware GOMAXPROCS** - コンテナ環境でのCPU最適化
2. **testing/synctest Package** - 並行処理テストの品質向上

### 重要
3. **Trace Flight Recorder** - 本番環境でのパフォーマンス分析
4. **encoding/json/v2 Package** - 改良されたJSON処理

### 有用
5. **go.mod ignore Directive** - 依存関係管理の向上
6. **Experimental Green Tea GC** - ガベージコレクション最適化
7. **go doc -http Option** - ローカル開発効率化

## 使用方法

### 前提条件

#### Docker Compose使用時
- Docker & Docker Compose
- Webブラウザ（Chrome, Firefox, Safari等）

#### ローカル開発時
- Go 1.25以上がインストールされている
- Webブラウザ（Chrome, Firefox, Safari等）

### インストール・実行

#### Docker Compose（推奨）

```bash
# リポジトリをクローン
git clone <repository-url>
cd go-release-tour

# 本番環境での起動
docker-compose up -d

# ログ確認
docker-compose logs -f
```

#### 開発環境（ホットリロード対応）

```bash
# リポジトリをクローン
git clone <repository-url>
cd go-release-tour

# 開発環境での起動（Air使用）
docker-compose -f docker-compose.dev.yml up

# または、ローカルでAirを使用
go install github.com/air-verse/air@latest
air
```

#### ローカル開発（手動）

```bash
# リポジトリをクローン
git clone <repository-url>
cd go-release-tour

# 依存関係をダウンロード
go mod download

# サーバーを起動
go run main.go
```

### アクセス
ブラウザで `http://localhost:8080` にアクセス

### 停止

```bash
# 本番環境停止
docker-compose down

# 開発環境停止
docker-compose -f docker-compose.dev.yml down

# ローカルAir停止
Ctrl+C

# コンテナとイメージも削除
docker-compose down --rmi all

## 使い方

1. **レッスン選択**: 左サイドバーから学習したい機能をクリック
2. **コード編集**: 中央のエディターでコードを編集
3. **実行**: 「▶ 実行」ボタンまたは `Ctrl+Enter` で実行
4. **結果確認**: 下部に実行結果が表示

## 特徴

### インタラクティブな学習体験
- **自動保存**: 編集したコードは自動的にブラウザに保存
- **キーボードショートカット**:
  - `Ctrl+Enter` / `Cmd+Enter`: コード実行
  - `Ctrl+S` / `Cmd+S`: 保存確認（視覚フィードバック）
- **タブサポート**: エディター内でTabキーによるインデント

### セキュアな実行環境
- **一時ファイル**: 実行時に一時ファイルを作成し、実行後自動削除
- **サンドボックス化**: サーバーサイドでコード実行（ローカル環境）

### レスポンシブUI
- **デスクトップ**: サイドバー + メインコンテンツのレイアウト
- **モバイル**: 縦積みレイアウトで最適化

## アーキテクチャ

```
go-release-tour/
├── main.go                      # Webサーバー（Go）
├── go.mod                       # Go モジュール定義
├── .air.toml                    # Air設定（ホットリロード）
├── docker/                      # Docker関連ファイル
│   ├── Dockerfile               # 本番用Dockerファイル
│   └── Dockerfile.dev           # 開発用Dockerファイル
├── docker-compose.yml           # 本番用Docker Compose
├── docker-compose.dev.yml       # 開発用Docker Compose
├── .gitignore                   # Git除外設定
├── static/                      # フロントエンド資材
│   ├── style.css                # スタイルシート
│   └── app.js                   # JavaScript
├── releases/v/                  # バージョン別サンプルコード
│   ├── 1.24/                    # Go 1.24新機能
│   │   ├── 01_generic_type_aliases.go
│   │   ├── 02_tool_dependencies.go
│   │   ├── ... (他5ファイル)
│   │   └── README.md
│   └── 1.25/                    # Go 1.25新機能
│       ├── 01_container_aware_gomaxprocs.go
│       ├── 02_trace_flight_recorder.go
│       ├── ... (他5ファイル)
│       └── README.md
└── README.md
```

### バックエンド（Go）
- **HTTP サーバー**: `net/http`を使用
- **API エンドポイント**:
  - `GET /api/versions`: 利用可能バージョン一覧
  - `GET /api/lessons?version=1.24`: バージョン別レッスン一覧取得
  - `POST /api/run`: コード実行
- **コード実行**: `os/exec`でGoコンパイル・実行
- **ホットリロード**: Air による自動リビルド対応

### フロントエンド（Vanilla JS + CSS）
- **ES6クラス**: モダンなJavaScript構文
- **Fetch API**: サーバーとの通信
- **CSS Grid/Flexbox**: レスポンシブレイアウト

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

## 対応ブラウザ

- Chrome 70+
- Firefox 65+
- Safari 12+
- Edge 79+

## 開発者向け

### 開発環境セットアップ

```bash
# Air（ホットリロード）をインストール
go install github.com/air-verse/air@latest

# 開発サーバー起動（ホットリロード有効）
air

# または Docker開発環境
docker-compose -f docker-compose.dev.yml up
```

### カスタマイズポイント

1. **新しいバージョン追加**: `releases/v/1.XX/`ディレクトリを作成
2. **新しいレッスン追加**: バージョンディレクトリに`.go`ファイルを追加
3. **スタイル変更**: `static/style.css`を編集（Airで自動リロード）
4. **JavaScript機能拡張**: `static/app.js`を編集（Airで自動リロード）
5. **サーバーロジック**: `main.go`を編集（Airで自動リビルド）

### Air設定

`.air.toml`で以下をカスタマイズ可能:
- **監視ファイル**: `include_ext = ["go", "html", "css", "js"]`
- **除外ディレクトリ**: `exclude_dir = ["tmp", "vendor"]`
- **ビルドコマンド**: `cmd = "go build -o ./tmp/main ."`
- **プロキシポート**: `proxy_port = 3000`

### 本番環境デプロイ
```bash
# Docker本番環境
docker-compose up -d

# または手動ビルド
go build -o go-release-tour main.go
./go-release-tour
```

## 今後の拡張予定

- [ ] **シンタックスハイライト**: Monaco EditorやCodeMirror統合
- [ ] **共有機能**: コードスニペットのURL共有
- [ ] **プリセット**: よくあるパターンのコードテンプレート
- [ ] **多言語対応**: 英語版インターフェース
- [ ] **Docker化**: コンテナでの簡単実行

## コントリビューション

1. Forkしてブランチを作成
2. 変更を実装
3. プルリクエストを作成

## ライセンス

MIT License

## 謝辞

- [Go Team](https://golang.org/team) - 素晴らしい言語とツールチェーン
- [Go Tour](https://go-tour-jp.appspot.com/) - インスピレーション源