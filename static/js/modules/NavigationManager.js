// ナビゲーション管理モジュール
class NavigationManager {
    constructor(tour) {
        this.tour = tour;
    }

    setupEventListeners() {
        // イベントリスナーを設定
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            versionSelect.addEventListener('change', (e) => {
                this.tour.currentVersion = e.target.value;
                this.tour.loadLessons(this.tour.currentVersion).then(() => {
                    this.tour.renderLessonList();

                    // 最初のレッスンを自動選択
                    const lessons = this.tour.lessons[this.tour.currentVersion] || [];
                    if (lessons.length > 0) {
                        this.tour.selectLesson(lessons[0].id, this.tour.currentVersion);
                    }
                });
            });
        }

        // Home button クリックイベント
        const homeBtn = document.getElementById('home-btn');
        if (homeBtn) {
            homeBtn.addEventListener('click', () => {
                this.tour.showWelcomeScreen();
            });
        }

        // Theme selector setup is handled by EditorManager

        // Run buttons
        const runBtn = document.getElementById('run-btn');
        if (runBtn) {
            runBtn.addEventListener('click', () => {
                this.tour.runCode();
            });
        }

        const runBtnBottom = document.getElementById('run-btn-bottom');
        if (runBtnBottom) {
            runBtnBottom.addEventListener('click', () => {
                this.tour.runCode();
            });
        }

        // 環境変数プリセットボタン
        const presetJsonV2Btn = document.getElementById('preset-jsonv2');
        if (presetJsonV2Btn) {
            presetJsonV2Btn.addEventListener('click', () => {
                const envVarsInput = document.getElementById('env-vars');
                if (envVarsInput) {
                    envVarsInput.value = 'GOEXPERIMENT=jsonv2';
                    // 視覚的フィードバック
                    presetJsonV2Btn.style.background = '#28a745';
                    presetJsonV2Btn.textContent = '✓ 設定済み';
                    setTimeout(() => {
                        presetJsonV2Btn.style.background = '#00ADD8';
                        presetJsonV2Btn.textContent = 'JSON v2';
                    }, 1500);
                }
            });
        }

        const presetClearBtn = document.getElementById('preset-clear');
        if (presetClearBtn) {
            presetClearBtn.addEventListener('click', () => {
                const envVarsInput = document.getElementById('env-vars');
                if (envVarsInput) {
                    envVarsInput.value = '';
                    // 視覚的フィードバック
                    presetClearBtn.style.background = '#28a745';
                    presetClearBtn.textContent = '✓ クリア済み';
                    setTimeout(() => {
                        presetClearBtn.style.background = '#6c757d';
                        presetClearBtn.textContent = 'クリア';
                    }, 1500);
                }
            });
        }


        // Fullscreen toggle
        const fullscreenBtn = document.getElementById('fullscreen-btn');
        if (fullscreenBtn) {
            fullscreenBtn.addEventListener('click', () => {
                this.toggleFullscreen();
            });
        }

        // Back to versions button
        const backBtn = document.getElementById('back-to-versions');
        if (backBtn) {
            backBtn.addEventListener('click', () => {
                this.backToVersions();
            });
        }

        // Welcome screen version selection
        this.setupWelcomeVersionSelection();

        // Global keyboard shortcuts
        this.setupKeyboardShortcuts();
    }

    setupWelcomeVersionSelection() {
        // バージョンカードのクリックイベントを設定（イベント委譲）
        document.addEventListener('click', (e) => {
            if (e.target.classList.contains('start-learning-btn') || e.target.closest('.version-card.clickable')) {
                e.preventDefault();

                let version;
                if (e.target.classList.contains('start-learning-btn')) {
                    version = e.target.dataset.version;
                } else {
                    const versionCard = e.target.closest('.version-card.clickable');
                    version = versionCard.dataset.version;
                }

                if (version) {
                    console.log(`Starting learning for version ${version}`);
                    this.tour.currentVersion = version;

                    // バージョンセレクターを更新
                    const versionSelect = document.getElementById('version-select');
                    if (versionSelect) {
                        versionSelect.value = version;
                    }

                    // レッスンを読み込んで最初のレッスンを表示
                    this.tour.loadLessons(version).then(() => {
                        this.tour.renderLessonList();
                        const lessons = this.tour.lessons[version] || [];
                        if (lessons.length > 0) {
                            this.tour.selectLesson(lessons[0].id, version);
                        }
                    });
                }
            }
        });
    }

    setupKeyboardShortcuts() {
        document.addEventListener('keydown', (e) => {
            // Ctrl/Cmd + Enter でコード実行
            if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
                e.preventDefault();
                this.tour.runCode();
            }

            // Ctrl/Cmd + S で保存（ブラウザのデフォルト保存を防止）
            if ((e.ctrlKey || e.metaKey) && e.key === 's') {
                e.preventDefault();
                // コードは自動保存される、短時間のインジケーターを表示
                const editor = document.getElementById('code-editor');
                if (editor) {
                    const originalBorder = editor.style.border;
                    editor.style.border = '2px solid #28a745';
                    setTimeout(() => {
                        editor.style.border = originalBorder;
                    }, 200);
                }
            }

            // ESCキーでエディター展開を解除
            if (e.key === 'Escape' && this.tour.isEditorExpanded) {
                e.preventDefault();
                this.tour.toggleEditorExpand();
            }

            // F11キーでエディター展開切り替え
            if (e.key === 'F11') {
                e.preventDefault();
                this.tour.toggleEditorExpand();
            }
        });
    }

    toggleFullscreen() {
        if (!document.fullscreenElement) {
            document.documentElement.requestFullscreen().catch(err => {
                console.error('Failed to enter fullscreen:', err);
            });
        } else {
            if (document.exitFullscreen) {
                document.exitFullscreen();
            }
        }
    }

    backToVersions() {
        // ウェルカム画面を表示してレッスン画面を非表示
        const welcomeScreen = document.getElementById('welcome-screen');
        const lessonView = document.getElementById('lesson-view');

        if (welcomeScreen && lessonView) {
            welcomeScreen.style.display = 'block';
            lessonView.style.display = 'none';
        }

        // 現在のレッスンをクリア
        this.tour.currentLesson = null;
        this.tour.currentVersion = null;

        // バージョンセレクターをリセット
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            versionSelect.value = '';
        }
    }
}

// NavigationManagerをGoReleaseTourに統合
GoReleaseTour.prototype.setupEventListeners = function() {
    if (!this.navigationManager) {
        this.navigationManager = new NavigationManager(this);
    }
    this.navigationManager.setupEventListeners();
};