// Go Release Tour - Interactive Web Tutorial for Go 1.25 Features
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Lesson struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Filename    string `json:"filename"`
	Stars       int    `json:"stars"`
	Version     string `json:"version"`
}

type Server struct {
	lessons map[string][]Lesson // version -> lessons
}

func main() {
	server := &Server{
		lessons: make(map[string][]Lesson),
	}
	server.loadLessons()

	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// API endpoints
	http.HandleFunc("/api/versions", server.handleVersions)
	http.HandleFunc("/api/lessons", server.handleLessons)
	http.HandleFunc("/api/run", server.handleRun)

	// Debug page
	http.HandleFunc("/debug.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "debug.html")
	})

	// Main page
	http.HandleFunc("/", server.handleIndex)

	// Get port from environment or default to 8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	goVersion := os.Getenv("GO_VERSION")
	if goVersion != "" {
		fmt.Printf("Go Release Tour server (Go %s) starting on :%s\n", goVersion, port)
	} else {
		fmt.Printf("Go Release Tour server starting on :%s\n", port)
	}
	fmt.Printf("Visit http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *Server) loadLessons() {
	// Load lessons based on GO_VERSION environment variable
	goVersion := os.Getenv("GO_VERSION")
	if goVersion != "" {
		s.loadVersionLessons(goVersion)
	} else {
		// Load all versions by default
		s.loadVersionLessons("1.25")
		s.loadVersionLessons("1.24")
	}
}

func (s *Server) loadVersionLessons(version string) {
	lessonsDir := fmt.Sprintf("releases/v/%s", version)
	files, err := filepath.Glob(filepath.Join(lessonsDir, "*.go"))
	if err != nil {
		log.Printf("Error loading lessons for version %s: %v", version, err)
		return
	}

	var lessonData map[string]struct {
		title string
		stars int
	}

	if version == "1.25" {
		lessonData = map[string]struct {
			title string
			stars int
		}{
			"01_container_aware_gomaxprocs.go": {"Container-aware GOMAXPROCS", 5},
			"02_trace_flight_recorder.go":      {"Trace Flight Recorder", 4},
			"03_testing_synctest.go":           {"testing/synctest Package", 5},
			"04_go_mod_ignore.go":              {"go.mod ignore Directive", 3},
			"05_experimental_gc.go":            {"Experimental Green Tea GC", 3},
			"06_json_v2.go":                    {"encoding/json/v2 Package", 4},
			"07_go_doc_http.go":                {"go doc -http Option", 3},
		}
	} else if version == "1.24" {
		lessonData = map[string]struct {
			title string
			stars int
		}{
			"01_generic_type_aliases.go": {"Generic Type Aliases", 5},
			"02_tool_dependencies.go":    {"Tool Dependencies in go.mod", 5},
			"03_swiss_tables_maps.go":    {"Swiss Tables Map Implementation", 5},
			"04_testing_loop.go":         {"testing.B.Loop() Method", 4},
			"05_os_root.go":              {"os.Root Type", 4},
			"06_crypto_mlkem.go":         {"crypto/mlkem Package", 5},
			"07_weak_pointers.go":        {"weak Package", 4},
		}
	} else {
		return
	}

	var lessons []Lesson
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

		// Extract description from comments
		lines := strings.Split(string(content), "\n")
		var description string
		for _, line := range lines {
			if strings.HasPrefix(line, "// 説明:") {
				description = strings.TrimPrefix(line, "// 説明: ")
				break
			}
		}

		lesson := Lesson{
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
	s.lessons[version] = lessons
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Release Tour - Go新機能学習</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <div id="app">
        <header>
            <h1>🚀 Go Release Tour</h1>
            <p>Goの新機能をインタラクティブに学習しよう</p>
        </header>

        <div class="container">
            <aside class="sidebar">
                <div class="version-selector">
                    <h3>バージョン選択</h3>
                    <select id="version-select">
                        <option value="1.25">Go 1.25</option>
                        <option value="1.24">Go 1.24</option>
                    </select>
                </div>
                <h3>レッスン一覧</h3>
                <div id="lesson-list"></div>
            </aside>

            <main class="content">
                <div id="lesson-content">
                    <div id="welcome-screen">
                        <h2>Welcome to Go Release Tour!</h2>
                        <p>Go の最新機能をインタラクティブに学習しましょう。学習したいバージョンを選択してください。</p>

                        <div class="version-selection">
                            <div class="version-card clickable" data-version="1.25">
                                <div class="version-header">
                                    <h3>Go 1.25</h3>
                                    <span class="version-badge latest">最新</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Container-aware GOMAXPROCS</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>testing/synctest Package</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐</span>
                                        <span>Trace Flight Recorder</span>
                                    </div>
                                    <div class="feature-more">+ 4つの機能</div>
                                </div>
                                <button class="start-learning-btn" data-version="1.25">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.24">
                                <div class="version-header">
                                    <h3>Go 1.24</h3>
                                    <span class="version-badge stable">安定版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Generic Type Aliases</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Swiss Tables Maps</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>crypto/mlkem Package</span>
                                    </div>
                                    <div class="feature-more">+ 4つの機能</div>
                                </div>
                                <button class="start-learning-btn" data-version="1.24">学習を開始</button>
                            </div>
                        </div>

                        <div class="tour-info">
                            <h3>Go Tour について</h3>
                            <p>このツアーでは、Go の最新機能を実際にコードを実行しながら学習できます。</p>
                            <ul>
                                <li>ブラウザ上でコードを編集・実行</li>
                                <li>各機能の実用性を5段階で評価</li>
                                <li>段階的に学習できる構成</li>
                            </ul>
                        </div>
                    </div>

                    <div id="lesson-view" style="display: none;">
                        <div class="lesson-header">
                            <button id="back-to-versions" class="back-btn">← バージョン選択に戻る</button>
                            <div class="lesson-title">
                                <h2 id="current-lesson-title"></h2>
                                <div id="current-lesson-stars"></div>
                            </div>
                        </div>
                        <div id="lesson-description"></div>
                    </div>
                </div>

                <div class="code-section">
                    <div class="code-header">
                        <h4>コードエディター</h4>
                        <button id="run-btn" onclick="runCode()">▶ 実行</button>
                    </div>
                    <textarea id="code-editor" placeholder="ここにGoコードを入力してください..."></textarea>
                </div>

                <div class="output-section">
                    <h4>実行結果</h4>
                    <pre id="output"></pre>
                </div>
            </main>
        </div>
    </div>

    <script src="/static/app.js"></script>
</body>
</html>`

	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, nil)
}

func (s *Server) handleVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	versions := make([]string, 0, len(s.lessons))
	for version := range s.lessons {
		versions = append(versions, version)
	}
	json.NewEncoder(w).Encode(versions)
}

func (s *Server) handleLessons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	version := r.URL.Query().Get("version")
	if version == "" {
		version = "1.25" // default version
	}
	if lessons, exists := s.lessons[version]; exists {
		json.NewEncoder(w).Encode(lessons)
	} else {
		http.Error(w, "Version not found", http.StatusNotFound)
	}
}

func (s *Server) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create temporary file
	tmpFile := filepath.Join(os.TempDir(), "go-tour-"+strconv.FormatInt(time.Now().UnixNano(), 10)+".go")
	defer os.Remove(tmpFile)

	if err := os.WriteFile(tmpFile, []byte(req.Code), 0644); err != nil {
		http.Error(w, "Failed to write temporary file", http.StatusInternalServerError)
		return
	}

	// Run the code
	cmd := exec.Command("go", "run", tmpFile)
	output, err := cmd.CombinedOutput()

	result := struct {
		Output string `json:"output"`
		Error  string `json:"error"`
	}{
		Output: string(output),
	}

	if err != nil {
		result.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}