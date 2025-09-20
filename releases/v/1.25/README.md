# Go 1.25 新機能検証プロジェクト

Go 1.25のリリースノートから主要な新機能を抽出し、サンプルプログラムで検証するプロジェクトです。

## 📋 検証対象の新機能

### 1. Container-aware GOMAXPROCS
**ファイル**: `01_container_aware_gomaxprocs.go`
- **原文**: "The runtime now considers CPU bandwidth limits when setting the default value for GOMAXPROCS in Linux containers"
- **説明**: LinuxコンテナのCPU制限を考慮したGOMAXPROCSの自動調整
- **実用性**: ⭐⭐⭐⭐⭐ (コンテナ環境で重要)

### 2. Trace Flight Recorder
**ファイル**: `02_trace_flight_recorder.go`
- **原文**: "The runtime now supports a trace flight recorder, which allows capturing traces in memory and writing out only the most significant segments"
- **説明**: メモリ内でのトレース収集と重要セグメントのみの書き出し
- **実用性**: ⭐⭐⭐⭐ (本番環境でのデバッグに有用)

### 3. testing/synctest パッケージ
**ファイル**: `03_testing_synctest.go`
- **原文**: "New testing/synctest package supports testing concurrent code"
- **説明**: 並行処理コードのテスト支援
- **実用性**: ⭐⭐⭐⭐⭐ (並行処理テストの品質向上)

### 4. go.mod ignore ディレクティブ
**ファイル**: `04_go_mod_ignore.go`
- **原文**: "The go.mod file format now supports an ignore directive"
- **説明**: 特定モジュールの明示的な除外
- **実用性**: ⭐⭐⭐ (セキュリティとポリシー管理)

### 5. 実験的ガベージコレクター (Green Tea GC)
**ファイル**: `05_experimental_gc.go`
- **原文**: "An experimental garbage collector is available, with potential 10-40% reduction in garbage collection overhead"
- **説明**: 10-40%のGCオーバーヘッド削減を目指した実験的GC
- **実用性**: ⭐⭐⭐ (実験的機能、本番使用は非推奨)

### 6. encoding/json/v2 パッケージ (実験的)
**ファイル**: `06_json_v2.go`
- **原文**: "Experimental encoding/json/v2: Improved JSON implementation"
- **説明**: 改良されたJSON実装
- **実用性**: ⭐⭐⭐⭐ (JSON処理の性能向上)

### 7. go doc -http オプション
**ファイル**: `07_go_doc_http.go`
- **原文**: "New go doc -http option to start documentation server"
- **説明**: ローカルドキュメントサーバーの起動
- **実用性**: ⭐⭐⭐ (開発効率の向上)

## 🚀 実行方法

### 個別実行
```bash
# 各サンプルを個別に実行（main関数の競合を避けるため、個別実行のみ対応）
go run 01_container_aware_gomaxprocs.go
go run 02_trace_flight_recorder.go
go run 03_testing_synctest.go
go run 04_go_mod_ignore.go
go run 05_experimental_gc.go
go run 06_json_v2.go
go run 07_go_doc_http.go
```

### 📁 プロジェクト構成について
- `01_*.go` がメインファイル（通常のmain関数）
- `02_*.go`以降は `// +build ignore` でビルド対象から除外
- 各ファイルは個別に `go run` で実行する設計

### 実験的機能の有効化
```bash
# Green Tea GC を有効にして実行
GOEXPERIMENT=greenteagc go run 05_experimental_gc.go

# 複数の実験的機能を有効化
GOEXPERIMENT=greenteagc,jsonv2 go run your_program.go
```

### コンテナでのテスト
```bash
# Container-aware GOMAXPROCS をテスト
docker run --cpus=2.5 golang:1.25 go run 01_container_aware_gomaxprocs.go
```

## 📊 LT発表用まとめ

### 🎯 開発者にとって重要度の高い機能

1. **Container-aware GOMAXPROCS** (⭐⭐⭐⭐⭐)
   - コンテナ時代必須の機能
   - Kubernetes環境で効果を発揮

2. **testing/synctest** (⭐⭐⭐⭐⭐)
   - 並行処理テストの信頼性向上
   - フレキー テストの削減

3. **Trace Flight Recorder** (⭐⭐⭐⭐)
   - 本番環境でのパフォーマンス分析
   - 低オーバーヘッドでの問題検出

### 🔬 実験的機能（将来有望）

1. **Green Tea GC** (⭐⭐⭐)
   - 大規模アプリケーションでの効果期待
   - 将来的にデフォルトGCになる可能性

2. **encoding/json/v2** (⭐⭐⭐⭐)
   - API開発での性能向上
   - 後方互換性を保った改良

### 🛠️ 開発効率向上

1. **go doc -http** (⭐⭐⭐)
   - ローカル開発での利便性
   - チーム開発でのドキュメント共有

2. **go.mod ignore** (⭐⭐⭐)
   - セキュリティポリシーの強制
   - 依存関係管理の向上

## 📈 パフォーマンス改善

- **GC最適化**: Green Tea GCで10-40%の改善
- **JSON処理**: v2パッケージで20-50%の高速化
- **コンテナ効率**: CPU制限に応じた最適化
- **メモリ効率**: トレースオーバーヘッドの削減

## 🎪 LT スライド構成案

1. **タイトル**: "Go 1.25の新機能を全部試してみた"
2. **概要**: 7つの主要機能の紹介
3. **実演**: Container-aware GOMAXPROCSのデモ
4. **注目**: testing/synctestの重要性
5. **将来性**: 実験的機能の可能性
6. **まとめ**: アップグレードの価値

## ⚠️ 注意事項

- 実験的機能は本番環境での使用を避ける
- Go 1.25が正式リリースされてから使用する
- パフォーマンステストで効果を確認する
- 段階的な機能採用を推奨

## 🔗 参考資料

- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [Go Blog](https://go.dev/blog/)
- [Go GitHub Repository](https://github.com/golang/go)