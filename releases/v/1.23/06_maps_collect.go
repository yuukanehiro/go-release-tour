// Go 1.23 新機能: maps.Collect Function
// 原文: "New maps.Collect function creates maps from iterator sequences"
//
// 説明: Go 1.23では、maps.Collect関数が追加され、
// イテレーター関数から効率的にマップを作成できるようになりました。

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// maps.Collectのシミュレーション関数
func collectStringInt(iterator func(yield func(string, int) bool)) map[string]int {
	result := make(map[string]int)
	iterator(func(key string, value int) bool {
		result[key] = value
		return true
	})
	return result
}

func collectStringString(iterator func(yield func(string, string) bool)) map[string]string {
	result := make(map[string]string)
	iterator(func(key string, value string) bool {
		result[key] = value
		return true
	})
	return result
}

func collectStringMapString(iterator func(yield func(string, map[string]string) bool)) map[string]map[string]string {
	result := make(map[string]map[string]string)
	iterator(func(key string, value map[string]string) bool {
		result[key] = value
		return true
	})
	return result
}

func collectStringFloat64(iterator func(yield func(string, float64) bool)) map[string]float64 {
	result := make(map[string]float64)
	iterator(func(key string, value float64) bool {
		result[key] = value
		return true
	})
	return result
}

func collectIntInt(iterator func(yield func(int, int) bool)) map[int]int {
	result := make(map[int]int)
	iterator(func(key int, value int) bool {
		result[key] = value
		return true
	})
	return result
}

func collectIntSliceString(iterator func(yield func(int, []string) bool)) map[int][]string {
	result := make(map[int][]string)
	iterator(func(key int, value []string) bool {
		result[key] = value
		return true
	})
	return result
}

func main() {
	fmt.Println("=== maps.Collect Function Demo ===")
	fmt.Println("注意: この例はmaps.Collectの概念を示すもので、実際の環境では互換性のある実装を使用しています。")

	// 基本的な使用例
	demonstrateBasicCollect()

	// 実用的なデータ変換
	demonstrateDataTransformation()

	// フィルタリングとマッピング
	demonstrateFilteringAndMapping()

	// 複雑なデータ処理
	demonstrateComplexProcessing()
}

func demonstrateBasicCollect() {
	fmt.Println("\n--- 基本的な使用例 ---")

	// キー・バリューペアのイテレーター
	keyValueIter := func(yield func(string, int) bool) {
		pairs := [][2]interface{}{
			{"apple", 100},
			{"banana", 200},
			{"orange", 150},
		}

		for _, pair := range pairs {
			key := pair[0].(string)
			value := pair[1].(int)
			if !yield(key, value) {
				return
			}
		}
	}

	// イテレーターからマップを作成
	fruitPrices := collectStringInt(keyValueIter)
	fmt.Printf("果物の価格: %v\n", fruitPrices)

	// 数値範囲からマップを作成
	numberSquares := collectIntInt(func(yield func(int, int) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, i*i) {
				return
			}
		}
	})
	fmt.Printf("数値の平方: %v\n", numberSquares)
}

func demonstrateDataTransformation() {
	fmt.Println("\n--- 実用的なデータ変換 ---")

	// CSVライクなデータからマップを作成
	csvData := []string{
		"id:1,name:田中,age:25",
		"id:2,name:佐藤,age:30",
		"id:3,name:鈴木,age:28",
	}

	fmt.Println("元データ:")
	for _, line := range csvData {
		fmt.Printf("  %s\n", line)
	}

	// CSVデータをパースしてマップに変換
	userMap := collectStringMapString(func(yield func(string, map[string]string) bool) {
		for _, line := range csvData {
			parts := strings.Split(line, ",")
			userData := make(map[string]string)
			var id string

			for _, part := range parts {
				kv := strings.Split(part, ":")
				if len(kv) == 2 {
					key, value := kv[0], kv[1]
					if key == "id" {
						id = value
					}
					userData[key] = value
				}
			}

			if id != "" {
				if !yield(id, userData) {
					return
				}
			}
		}
	})

	fmt.Println("\n変換後のマップ:")
	for id, user := range userMap {
		fmt.Printf("  ID %s: %v\n", id, user)
	}

	// 設定ファイルのパース例
	fmt.Println("\n--- 設定ファイルパース例 ---")
	configLines := []string{
		"server.host=localhost",
		"server.port=8080",
		"database.host=db.example.com",
		"database.port=5432",
		"debug.enabled=true",
	}

	configMap := collectStringString(func(yield func(string, string) bool) {
		for _, line := range configLines {
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				key, value := parts[0], parts[1]
				if !yield(key, value) {
					return
				}
			}
		}
	})

	fmt.Println("設定マップ:")
	for key, value := range configMap {
		fmt.Printf("  %s = %s\n", key, value)
	}
}

func demonstrateFilteringAndMapping() {
	fmt.Println("\n--- フィルタリングとマッピング ---")

	// 元データ
	students := []struct {
		Name  string
		Score int
		Grade string
	}{
		{"田中", 85, "A"},
		{"佐藤", 92, "A"},
		{"鈴木", 78, "B"},
		{"高橋", 95, "A"},
		{"渡辺", 82, "B"},
	}

	fmt.Println("学生データ:")
	for _, s := range students {
		fmt.Printf("  %s: %d点 (%s)\n", s.Name, s.Score, s.Grade)
	}

	// A評価の学生のみをマップに
	topStudents := collectStringInt(func(yield func(string, int) bool) {
		for _, student := range students {
			if student.Grade == "A" {
				if !yield(student.Name, student.Score) {
					return
				}
			}
		}
	})

	fmt.Println("\nA評価の学生:")
	for name, score := range topStudents {
		fmt.Printf("  %s: %d点\n", name, score)
	}

	// スコアを正規化（0-1の範囲に変換）
	normalizedScores := collectStringFloat64(func(yield func(string, float64) bool) {
		maxScore := 100.0
		for _, student := range students {
			normalized := float64(student.Score) / maxScore
			if !yield(student.Name, normalized) {
				return
			}
		}
	})

	fmt.Println("\n正規化されたスコア:")
	for name, score := range normalizedScores {
		fmt.Printf("  %s: %.2f\n", name, score)
	}
}

func demonstrateComplexProcessing() {
	fmt.Println("\n--- 複雑なデータ処理 ---")

	// ログデータの分析例
	logEntries := []string{
		"2024-01-15 09:00:01 INFO User login: user123",
		"2024-01-15 09:05:15 ERROR Database connection failed",
		"2024-01-15 09:10:30 INFO User logout: user123",
		"2024-01-15 09:15:45 WARN High memory usage: 85%",
		"2024-01-15 09:20:12 ERROR API timeout: /api/users",
		"2024-01-15 09:25:33 INFO User login: user456",
	}

	fmt.Println("ログエントリ:")
	for _, entry := range logEntries {
		fmt.Printf("  %s\n", entry)
	}

	// ログレベル別の件数を集計
	logCounts := collectStringInt(func(yield func(string, int) bool) {
		counts := make(map[string]int)

		for _, entry := range logEntries {
			parts := strings.Fields(entry)
			if len(parts) >= 3 {
				level := parts[2]
				counts[level]++
			}
		}

		for level, count := range counts {
			if !yield(level, count) {
				return
			}
		}
	})

	fmt.Println("\nログレベル別件数:")
	for level, count := range logCounts {
		fmt.Printf("  %s: %d件\n", level, count)
	}

	// 時間帯別の活動量を分析
	hourlyActivity := collectStringInt(func(yield func(string, int) bool) {
		activity := make(map[string]int)

		for _, entry := range logEntries {
			parts := strings.Fields(entry)
			if len(parts) >= 2 {
				timePart := parts[1]
				if len(timePart) >= 2 {
					hour := timePart[:2]
					activity[hour]++
				}
			}
		}

		for hour, count := range activity {
			if !yield(hour, count) {
				return
			}
		}
	})

	fmt.Println("\n時間帯別活動量:")
	// 時間順にソートして表示
	hours := make([]string, 0, len(hourlyActivity))
	for hour := range hourlyActivity {
		hours = append(hours, hour)
	}
	sort.Strings(hours)

	for _, hour := range hours {
		count := hourlyActivity[hour]
		fmt.Printf("  %s時台: %d件\n", hour, count)
	}
}

// 従来の方法との比較
func demonstrateComparison() {
	fmt.Println("\n--- 従来の方法との比較 ---")

	data := []struct {
		Key   string
		Value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	// 従来の方法
	traditional := make(map[string]int)
	for _, item := range data {
		traditional[item.Key] = item.Value
	}

	// collectStringIntを使用
	modern := collectStringInt(func(yield func(string, int) bool) {
		for _, item := range data {
			if !yield(item.Key, item.Value) {
				return
			}
		}
	})

	fmt.Printf("従来の方法: %v\n", traditional)
	fmt.Printf("collectStringInt: %v\n", modern)

	// より複雑な変換での比較
	fmt.Println("\n複雑な変換での比較:")

	// 従来の方法（文字列の長さをキーとする）
	traditionalByLength := make(map[int][]string)
	words := []string{"Go", "言語", "プログラミング", "開発"}

	for _, word := range words {
		length := len(word)
		traditionalByLength[length] = append(traditionalByLength[length], word)
	}

	// collectIntSliceStringを使った方法
	modernByLength := collectIntSliceString(func(yield func(int, []string) bool) {
		lengthMap := make(map[int][]string)
		for _, word := range words {
			length := len(word)
			lengthMap[length] = append(lengthMap[length], word)
		}

		for length, wordList := range lengthMap {
			if !yield(length, wordList) {
				return
			}
		}
	})

	fmt.Printf("従来の方法: %v\n", traditionalByLength)
	fmt.Printf("collectIntSliceString: %v\n", modernByLength)
}

// パフォーマンステスト
func demonstratePerformance() {
	fmt.Println("\n--- パフォーマンステスト ---")

	size := 1000
	data := make([]struct{ Key, Value string }, size)

	for i := 0; i < size; i++ {
		data[i] = struct{ Key, Value string }{
			Key:   "key" + strconv.Itoa(i),
			Value: "value" + strconv.Itoa(i*2),
		}
	}

	// collectStringStringでマップを作成
	result := collectStringString(func(yield func(string, string) bool) {
		for _, item := range data {
			if !yield(item.Key, item.Value) {
				return
			}
		}
	})

	fmt.Printf("作成されたマップのサイズ: %d\n", len(result))
	fmt.Printf("最初の5つのエントリ:\n")

	count := 0
	for key, value := range result {
		if count >= 5 {
			break
		}
		fmt.Printf("  %s: %s\n", key, value)
		count++
	}
}

// % go run 06_maps_collect.go
// === maps.Collect Function Demo ===
//
// --- 基本的な使用例 ---
// 果物の価格: map[apple:100 banana:200 orange:150]
// 数値の平方: map[1:1 2:4 3:9 4:16 5:25]
//
// --- 実用的なデータ変換 ---
// 元データ:
//   id:1,name:田中,age:25
//   id:2,name:佐藤,age:30
//   id:3,name:鈴木,age:28
//
// 変換後のマップ:
//   ID 1: map[age:25 id:1 name:田中]
//   ID 2: map[age:30 id:2 name:佐藤]
//   ID 3: map[age:28 id:3 name:鈴木]