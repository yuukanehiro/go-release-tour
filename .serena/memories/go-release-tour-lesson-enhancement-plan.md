# Go Release Tour レッスン拡張・改善計画

## 現在のレッスン構造分析

### データ構造
```go
type Lesson struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Code        string `json:"code"`
    Filename    string `json:"filename"`
    Stars       int    `json:"stars"`
    Version     string `json:"version"`
}
```

### 現在のレッスン配置
- **Go 1.18**: generics.go, workspace_mode.go, type_constraints.go, generic_data_structures.go, type_inference.go
- **Go 1.19**: memory_arenas.go, atomic_types.go
- **Go 1.20**: comparable_types.go, slice_to_array_conversion.go, errors_join.go
- **Go 1.21**: built_in_functions.go, slices_package.go, maps_package.go
- **Go 1.22**: for_range_integers.go, loop_variables.go, math_rand_v2.go, slices_concat.go, enhanced_http_routing.go
- **Go 1.23**: structured_logging.go, iterators.go, timer_reset.go, slices_concat.go, cmp_or.go, maps_collect.go
- **Go 1.24**: generic_type_aliases.go, tool_dependencies.go, swiss_tables_maps.go, testing_loop.go, os_root.go, crypto_mlkem.go, weak_pointers.go
- **Go 1.25**: container_aware_gomaxprocs.go, trace_flight_recorder.go, testing_synctest.go, go_mod_ignore.go, experimental_gc.go, json_v2.go, go_doc_http.go

## 改善計画

### 1. 公式リリースノートへのリンク追加
各レッスンに以下のメタデータを追加する必要：
- `ReleaseNotesURL`: 公式リリースノートの該当セクションへの直接リンク
- `GoDocURL`: 該当パッケージのGo公式ドキュメントリンク（該当する場合）
- `ProposalURL`: Go Proposalへのリンク（該当する場合）

### 2. 不足している重要機能の特定

#### Go 1.18
- **追加候補**: Type Parameters in Function Signatures, Generic Interfaces, Type Switches with Generics

#### Go 1.19
- **追加候補**: 
  - `sync/atomic`の新しい型
  - `runtime/debug`の改善
  - `doc comment`の形式改善

#### Go 1.20
- **追加候補**:
  - `context.WithCancelCause`
  - `crypto/rand`の改善
  - `time.Time`のLayoutの新しいConstant

#### Go 1.21
- **追加候補**:
  - `log/slog`パッケージ（構造化ログ）
  - `slices.SortFunc`などの高度な機能
  - `clear()`組み込み関数

#### Go 1.22
- **追加候補**:
  - `crypto/tls`の改善
  - `database/sql`の新機能
  - `net/http`のルーティング機能詳細

#### Go 1.23
- **追加候補**:
  - 独自イテレーター作成
  - `unique`パッケージ
  - `time`パッケージの改善

#### Go 1.24
- **追加候補**:
  - `math/rand/v2`の活用
  - `go.mod`のtoolchain指定
  - プロファイリング改善

#### Go 1.25
- **追加候補**:
  - WebAssembly System Interface (WASI)
  - `encoding/gob`の改善
  - `go vet`の新しいチェック

### 3. エビデンス性向上のための具体的URLs

#### Go 1.18
- Release Notes: https://go.dev/doc/go1.18
- Generics Proposal: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md

#### Go 1.19  
- Release Notes: https://go.dev/doc/go1.19

#### Go 1.20
- Release Notes: https://go.dev/doc/go1.20

#### Go 1.21
- Release Notes: https://go.dev/doc/go1.21

#### Go 1.22
- Release Notes: https://go.dev/doc/go1.22

#### Go 1.23
- Release Notes: https://go.dev/doc/go1.23

#### Go 1.24
- Release Notes: https://go.dev/doc/go1.24

#### Go 1.25
- Release Notes: https://go.dev/doc/go1.25

### 4. 実装アプローチ
1. **Lesson構造体の拡張**: URL関連フィールドの追加
2. **既存レッスンファイルの更新**: メタデータコメントの追加
3. **新規レッスンファイルの作成**: 重要機能のカバー
4. **フロントエンドの更新**: リンク表示機能の追加

### 5. 優先順位
1. **高**: 公式リリースノートリンクの追加
2. **高**: Go 1.21のslog、Go 1.22のHTTPルーティング
3. **中**: 各バージョンの主要機能の追加レッスン
4. **低**: フロントエンドでのリンク表示機能

## 次のステップ
1. Lesson構造体にURL関連フィールドを追加
2. 重要度の高いレッスンから順次作成・更新
3. 公式ドキュメントとの整合性確保
4. ユーザーエクスペリエンスの向上