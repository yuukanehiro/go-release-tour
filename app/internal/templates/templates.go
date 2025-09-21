package templates

import (
	"html/template"
	"net/http"
)

// HandleIndex serves the main application page
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Release Tour - Go新機能学習</title>
    <link rel="icon" type="image/svg+xml" href="/static/favicon.svg">
    <link rel="stylesheet" href="/static/style.css">
    <!-- CodeMirror CSS -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/codemirror.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/theme/monokai.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/theme/material.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/theme/dracula.min.css">
    <!-- CodeMirror JavaScript -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/codemirror.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/mode/go/go.min.js"></script>
</head>
<body>
    <div id="app">
        <header>
            <div class="header-logo clickable-title" id="home-btn">
                <img src="/static/header-logo.png" alt="Go Release Tour" class="logo-image">
            </div>
            <p>Goの新機能をインタラクティブに学習しよう</p>
        </header>

        <div class="container">
            <aside class="sidebar">
                <div class="version-selector">
                    <h3>バージョン選択</h3>
                    <select id="version-select">
                        <option value="1.25" selected>Go 1.25 (最新)</option>
                        <option value="1.24">Go 1.24</option>
                        <option value="1.23">Go 1.23</option>
                        <option value="1.22">Go 1.22</option>
                        <option value="1.21">Go 1.21</option>
                        <option value="1.20">Go 1.20</option>
                        <option value="1.19">Go 1.19</option>
                        <option value="1.18">Go 1.18 (Generics)</option>
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

                            <div class="version-card clickable" data-version="1.23">
                                <div class="version-header">
                                    <h3>Go 1.23</h3>
                                    <span class="version-badge">人気版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Structured Logging</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Range over Function Types</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐</span>
                                        <span>Timer.Reset改善</span>
                                    </div>
                                    <div class="feature-more">+ 3つの機能</div>
                                </div>
                                <button class="start-learning-btn" data-version="1.23">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.22">
                                <div class="version-header">
                                    <h3>Go 1.22</h3>
                                    <span class="version-badge">改良版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>For-Range over Integers</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Enhanced Loop Variables</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐</span>
                                        <span>math/rand/v2</span>
                                    </div>
                                    <div class="feature-more">+ 1つの機能</div>
                                </div>
                                <button class="start-learning-btn" data-version="1.22">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.21">
                                <div class="version-header">
                                    <h3>Go 1.21</h3>
                                    <span class="version-badge">基盤版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Built-in Functions (min/max/clear)</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>slices Package</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>maps Package</span>
                                    </div>
                                </div>
                                <button class="start-learning-btn" data-version="1.21">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.20">
                                <div class="version-header">
                                    <h3>Go 1.20</h3>
                                    <span class="version-badge">改善版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐</span>
                                        <span>Comparable Types Enhancement</span>
                                    </div>
                                </div>
                                <button class="start-learning-btn" data-version="1.20">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.19">
                                <div class="version-header">
                                    <h3>Go 1.19</h3>
                                    <span class="version-badge">実験版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐</span>
                                        <span>Memory Arenas (experimental)</span>
                                    </div>
                                </div>
                                <button class="start-learning-btn" data-version="1.19">学習を開始</button>
                            </div>

                            <div class="version-card clickable" data-version="1.18">
                                <div class="version-header">
                                    <h3>Go 1.18</h3>
                                    <span class="version-badge revolutionary">革命版</span>
                                </div>
                                <div class="version-features">
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Generics (Type Parameters)</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐⭐</span>
                                        <span>Type Constraints</span>
                                    </div>
                                    <div class="feature-item">
                                        <span class="stars">⭐⭐⭐⭐</span>
                                        <span>Workspace Mode</span>
                                    </div>
                                    <div class="feature-more">+ 2つの機能</div>
                                </div>
                                <button class="start-learning-btn" data-version="1.18">学習を開始</button>
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
                        <div id="lesson-links" style="display: none;">
                            <h4>参考リンク</h4>
                            <div class="links-container">
                                <a id="release-notes-link" href="#" target="_blank" rel="noopener noreferrer">
                                    Go Release Notes
                                </a>
                                <a id="go-doc-link" href="#" target="_blank" rel="noopener noreferrer" style="display: none;">
                                    Go Documentation
                                </a>
                                <a id="proposal-link" href="#" target="_blank" rel="noopener noreferrer" style="display: none;">
                                    Go Proposal
                                </a>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="code-section">
                    <div class="code-header">
                        <h4>コードエディター</h4>
                        <div class="editor-controls">
                            <select id="theme-selector">
                                <option value="default">デフォルト</option>
                                <option value="monokai" selected>Monokai (ダーク)</option>
                                <option value="material">Material</option>
                                <option value="dracula">Dracula</option>
                            </select>
                            <button id="run-btn">▶ 実行</button>
                        </div>
                    </div>
                    <div id="code-editor-container">
                        <textarea id="code-editor" placeholder="ここにGoコードを入力してください..."></textarea>
                    </div>
                    <div class="code-editor-footer">
                        <button id="run-btn-bottom" class="run-btn-bottom">▶ 実行</button>
                    </div>
                </div>

                <div class="output-section">
                    <h4>実行結果</h4>
                    <pre id="output"></pre>
                </div>
            </main>
        </div>
    </div>

    <!-- JavaScript modules -->
    <script src="/static/js/components/GoReleaseTour.js"></script>
    <script src="/static/js/modules/ApiClient.js"></script>
    <script src="/static/js/modules/EditorManager.js"></script>
    <script src="/static/js/modules/NavigationManager.js"></script>
    <script src="/static/js/modules/WelcomeScreen.js"></script>
    <script src="/static/js/modules/LessonDisplay.js"></script>
    <script src="/static/js/app.js"></script>
</body>
</html>`

	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, nil)
}