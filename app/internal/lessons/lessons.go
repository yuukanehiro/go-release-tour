package lessons

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-release-tour/app/internal/types"
)

// LoadLessons loads all lessons or specific version based on GO_VERSION environment variable
func LoadLessons(s *types.Server) {
	// GO_VERSION環境変数に基づいてレッスンを読み込み
	goVersion := os.Getenv("GO_VERSION")
	if goVersion != "" {
		loadVersionLessons(s, goVersion)
	} else {
		// デフォルトで全バージョンを読み込み
		loadVersionLessons(s, "1.25")
		loadVersionLessons(s, "1.24")
		loadVersionLessons(s, "1.23")
		loadVersionLessons(s, "1.22")
		loadVersionLessons(s, "1.21")
		loadVersionLessons(s, "1.20")
		loadVersionLessons(s, "1.19")
		loadVersionLessons(s, "1.18")
	}
}

// loadVersionLessons loads lessons for a specific Go version
func loadVersionLessons(s *types.Server, version string) {
	lessonsDir := fmt.Sprintf("releases/v/%s", version)
	files, err := filepath.Glob(filepath.Join(lessonsDir, "*.go"))
	if err != nil {
		log.Printf("Error loading lessons for version %s: %v", version, err)
		return
	}

	lessonData := getLessonData(version)
	if lessonData == nil {
		return
	}

	var lessons []types.Lesson
	for i, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}

		filename := filepath.Base(file)
		data, exists := lessonData[filename]
		if !exists {
			continue
		}

		// コメントから説明を抽出
		lines := strings.Split(string(content), "\n")
		var description string
		for _, line := range lines {
			if strings.HasPrefix(line, "// 説明:") {
				description = strings.TrimPrefix(line, "// 説明: ")
				break
			}
		}

		lesson := types.Lesson{
			ID:          i + 1,
			Title:       data.title,
			Description: description,
			Code:        string(content),
			Filename:    filename,
			Stars:       data.stars,
			Version:     version,
		}
		lessons = append(lessons, lesson)
	}
	s.Lessons[version] = lessons
}

// lessonInfo holds lesson metadata
type lessonInfo struct {
	title string
	stars int
}

// getLessonData returns lesson metadata for a specific version
func getLessonData(version string) map[string]lessonInfo {
	switch version {
	case "1.25":
		return map[string]lessonInfo{
			"01_container_aware_gomaxprocs.go": {"Container-aware GOMAXPROCS", 5},
			"02_trace_flight_recorder.go":      {"Trace Flight Recorder", 4},
			"03_testing_synctest.go":           {"testing/synctest Package", 5},
			"04_go_mod_ignore.go":              {"go.mod ignore Directive", 3},
			"05_experimental_gc.go":            {"Experimental Green Tea GC", 3},
			"06_json_v2.go":                    {"encoding/json/v2 Package", 4},
			"07_go_doc_http.go":                {"go doc -http Option", 3},
		}
	case "1.24":
		return map[string]lessonInfo{
			"01_generic_type_aliases.go": {"Generic Type Aliases", 5},
			"02_tool_dependencies.go":    {"Tool Dependencies in go.mod", 5},
			"03_swiss_tables_maps.go":    {"Swiss Tables Map Implementation", 5},
			"04_testing_loop.go":         {"testing.B.Loop() Method", 4},
			"05_os_root.go":              {"os.Root Type", 4},
			"06_crypto_mlkem.go":         {"crypto/mlkem Package", 5},
			"07_weak_pointers.go":        {"weak Package", 4},
		}
	case "1.23":
		return map[string]lessonInfo{
			"01_structured_logging.go": {"Structured Logging (log/slog)", 5},
			"02_iterators.go":          {"Range over Function Types", 5},
			"03_timer_reset.go":        {"Timer.Reset Behavior Change", 4},
			"04_slices_concat.go":      {"slices.Concat Function", 4},
			"05_cmp_or.go":             {"cmp.Or Function", 3},
			"06_maps_collect.go":       {"maps.Collect Function", 4},
		}
	case "1.22":
		return map[string]lessonInfo{
			"01_for_range_integers.go":    {"For-Range over Integers", 5},
			"02_loop_variables.go":        {"Enhanced Loop Variables", 5},
			"03_math_rand_v2.go":          {"math/rand/v2 Package", 4},
			"04_slices_concat.go":         {"slices Package Enhancements", 4},
			"05_enhanced_http_routing.go": {"Enhanced HTTP Routing", 5},
		}
	case "1.21":
		return map[string]lessonInfo{
			"01_built_in_functions.go": {"Built-in Functions (min, max, clear)", 5},
			"02_slices_package.go":     {"slices Package", 5},
			"03_maps_package.go":       {"maps Package", 5},
		}
	case "1.20":
		return map[string]lessonInfo{
			"01_comparable_types.go":         {"Comparable Types Enhancement", 4},
			"02_slice_to_array_conversion.go": {"Slice to Array Conversion", 4},
			"03_errors_join.go":              {"errors.Join Function", 4},
		}
	case "1.19":
		return map[string]lessonInfo{
			"01_memory_arenas.go": {"Memory Arenas (experimental)", 3},
			"02_atomic_types.go":  {"New Atomic Types", 4},
		}
	case "1.18":
		return map[string]lessonInfo{
			"01_generics.go":              {"Generics (Type Parameters)", 5},
			"02_workspace_mode.go":        {"Workspace Mode", 4},
			"03_type_constraints.go":      {"Type Constraints", 5},
			"04_generic_data_structures.go": {"Generic Data Structures", 5},
			"05_type_inference.go":        {"Type Inference", 4},
		}
	default:
		return nil
	}
}