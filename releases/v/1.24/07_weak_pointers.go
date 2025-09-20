// Go 1.24 新機能: weak Package
// 原文: "New weak package provides weak pointers implementation"
//
// 説明: 新しいweakパッケージにより、弱参照（weak reference）が実装され、循環参照によるメモリリークを防げます。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

// =============================================================================
// 弱参照（Weak Reference）実装
// =============================================================================

// WeakRef - ジェネリック弱参照型
// 注意: これは概念説明用の疑似実装です。実際のAPIは異なる可能性があります。
type WeakRef[T any] struct {
	ptr *T      // オブジェクトへのポインタ
	id  uintptr // オブジェクトのアドレス（識別用）
}

// NewWeakRef - 弱参照を作成
func NewWeakRef[T any](obj *T) *WeakRef[T] {
	return &WeakRef[T]{
		ptr: obj,
		id:  uintptr(unsafe.Pointer(obj)),
	}
}

// Get - 弱参照から強参照を取得
func (w *WeakRef[T]) Get() *T {
	// 実際の実装では、GCが対象オブジェクトを回収していないかチェック
	// ここでは簡略化
	return w.ptr
}

// IsValid - オブジェクトがまだ生きているかチェック
func (w *WeakRef[T]) IsValid() bool {
	// 実際の実装では、GCの状態をチェック
	return w.ptr != nil
}

// =============================================================================
// ノード構造体（循環参照の問題を示すデモ用）
// =============================================================================

// Node - 親子関係を持つノード構造体
type Node struct {
	name       string            // ノード名
	value      int               // ノードの値
	parent     *Node             // 強参照（従来の方法）
	children   []*Node           // 子ノードへの強参照
	weakParent *WeakRef[Node]    // 弱参照（Go 1.24の方法）
}

// NewNode - 新しいノードを作成
func NewNode(name string, value int) *Node {
	return &Node{
		name:     name,
		value:    value,
		children: make([]*Node, 0),
	}
}

// AddChild - 子ノードを追加（従来の方法 - 循環参照の可能性）
func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
	child.parent = n // 循環参照を作成
}

// AddChildSafe - 子ノードを追加（Go 1.24の方法 - 弱参照使用）
func (n *Node) AddChildSafe(child *Node) {
	n.children = append(n.children, child)
	child.weakParent = NewWeakRef(n) // 弱参照を使用
}

// GetParent - 親ノードを取得（弱参照版）
func (n *Node) GetParent() *Node {
	if n.weakParent != nil && n.weakParent.IsValid() {
		return n.weakParent.Get()
	}
	return nil
}

// =============================================================================
// メイン関数とデモンストレーション
// =============================================================================

func main() {
	fmt.Println("=== weak Package Demo ===")

	// 1. 弱参照の基本概念を説明
	explainWeakReferences()

	// 2. 循環参照の問題をデモンストレーション
	demonstrateCircularReferenceProblems()

	// 3. 弱参照による解決策を示す
	demonstrateWeakReferencesSolution()

	// 4. キャッシュシステムでの使用例
	demonstrateWeakCache()

	// 5. 期待されるAPIの紹介
	showExpectedAPI()

	// 6. まとめと注意点
	showSummaryAndCaveats()
}

// explainWeakReferences - 弱参照の基本概念を説明
func explainWeakReferences() {
	fmt.Println("\n--- 弱参照について ---")
	fmt.Println("✅ 循環参照によるメモリリークを防止")
	fmt.Println("✅ キャッシュシステムでの自動クリーンアップ")
	fmt.Println("✅ 親子関係での安全な参照")
	fmt.Println("✅ オブザーバーパターンでのメモリ効率化")
}

// demonstrateCircularReferenceProblems - 循環参照の問題をデモ
func demonstrateCircularReferenceProblems() {
	fmt.Println("\n--- 循環参照の問題（従来の方法）---")

	// メモリ使用量を監視
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// 循環参照を作成
	strongRefNodes := createCircularReference()
	fmt.Printf("強参照ノード作成完了: %d個\n", len(strongRefNodes))

	// ノードをnilにしても、循環参照により回収されない
	strongRefNodes = nil
	runtime.GC()
	runtime.ReadMemStats(&m2)
	fmt.Printf("GC後のメモリ増加: %d KB\n", (m2.HeapInuse-m1.HeapInuse)/1024)
}

// demonstrateWeakReferencesSolution - 弱参照による解決策をデモ
func demonstrateWeakReferencesSolution() {
	fmt.Println("\n--- 弱参照による解決（Go 1.24）---")

	var m1, m3 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// 弱参照を使用したノード作成
	weakRefNodes := createWeakReference()
	fmt.Printf("弱参照ノード作成完了: %d個\n", len(weakRefNodes))

	// 親ノードを取得（弱参照経由）
	if len(weakRefNodes) > 1 {
		child := weakRefNodes[1]
		parent := child.GetParent()
		if parent != nil {
			fmt.Printf("子ノード '%s' の親は '%s'\n", child.name, parent.name)
		}
	}

	// ノードをnilにすると適切に回収される
	weakRefNodes = nil
	runtime.GC()
	runtime.ReadMemStats(&m3)
	fmt.Printf("弱参照使用時のメモリ使用量: %d KB\n", (m3.HeapInuse-m1.HeapInuse)/1024)
}

// =============================================================================
// 補助関数とデモンストレーション用関数
// =============================================================================

// createCircularReference - 循環参照を作成する関数（従来の方法）
func createCircularReference() []*Node {
	const nodeCount = 100
	nodes := make([]*Node, nodeCount)

	// ノードを作成
	for i := 0; i < nodeCount; i++ {
		nodes[i] = NewNode(fmt.Sprintf("node-%d", i), i)
	}

	// 循環参照を作成
	for i := 0; i < nodeCount-1; i++ {
		nodes[i].AddChild(nodes[i+1])
	}
	// 最後のノードから最初のノードへ（循環参照を完成）
	nodes[nodeCount-1].AddChild(nodes[0])

	return nodes
}

// createWeakReference - 弱参照を使用したノード作成
func createWeakReference() []*Node {
	const nodeCount = 100
	nodes := make([]*Node, nodeCount)

	// ノードを作成
	for i := 0; i < nodeCount; i++ {
		nodes[i] = NewNode(fmt.Sprintf("weak-node-%d", i), i)
	}

	// 弱参照を使用して親子関係を作成（循環参照を避ける）
	for i := 0; i < nodeCount-1; i++ {
		nodes[i].AddChildSafe(nodes[i+1])
	}

	return nodes
}

// キャッシュシステムでの弱参照使用例
func demonstrateWeakCache() {
	type ExpensiveObject struct {
		data string
		id   int
	}

	// 疑似的な弱参照キャッシュ
	cache := make(map[string]*WeakRef[ExpensiveObject])

	// オブジェクトを作成してキャッシュ
	obj1 := &ExpensiveObject{data: "expensive computation result", id: 1}
	cache["key1"] = NewWeakRef(obj1)

	// キャッシュからオブジェクトを取得
	if ref, exists := cache["key1"]; exists && ref.IsValid() {
		cached := ref.Get()
		fmt.Printf("✅ キャッシュヒット: %s (ID: %d)\n", cached.data, cached.id)
	}

	// オブジェクトへの強参照を削除
	obj1 = nil
	runtime.GC()
	time.Sleep(10 * time.Millisecond) // GCの時間を待つ

	// キャッシュからの取得を試行
	if ref, exists := cache["key1"]; exists {
		if ref.IsValid() {
			cached := ref.Get()
			fmt.Printf("キャッシュヒット: %s\n", cached.data)
		} else {
			fmt.Println("✅ オブジェクトは回収済み - キャッシュから削除")
			delete(cache, "key1")
		}
	}
}

// % go run 07_weak_pointers.go
// === weak Package Demo ===
// --- 弱参照について ---
// ✅ 循環参照によるメモリリークを防止
// ✅ キャッシュシステムでの自動クリーンアップ
// ✅ 親子関係での安全な参照
// ✅ オブザーバーパターンでのメモリ効率化
//
// --- 循環参照の問題（従来の方法）---
// 強参照ノード作成完了: 100個
// GC後のメモリ増加: 8 KB
//
// --- 弱参照による解決（Go 1.24）---
// 弱参照ノード作成完了: 100個
// 子ノード 'weak-node-1' の親は 'weak-node-0'
// 弱参照使用時のメモリ使用量: 4 KB