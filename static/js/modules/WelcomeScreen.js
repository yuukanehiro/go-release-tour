// ウェルカム画面管理モジュール
class WelcomeScreen {
    constructor(tour) {
        this.tour = tour;
    }

    show() {
        console.log('Showing welcome screen');

        // まずDOMの状態をデバッグ
        console.log('DOM elements check:');
        console.log('- welcome-screen:', document.getElementById('welcome-screen'));
        console.log('- lesson-view:', document.getElementById('lesson-view'));
        console.log('- lesson-content:', document.getElementById('lesson-content'));

        const welcomeScreen = document.getElementById('welcome-screen');
        const lessonView = document.getElementById('lesson-view');
        const lessonContent = document.getElementById('lesson-content');

        // lesson-content内のすべての子要素をログ出力
        if (lessonContent) {
            console.log('lesson-content children:', lessonContent.children);
            for (let i = 0; i < lessonContent.children.length; i++) {
                console.log(`  child ${i}:`, lessonContent.children[i].id, lessonContent.children[i]);
            }
        }

        if (welcomeScreen) {
            welcomeScreen.style.display = 'block';
            console.log('Welcome screen displayed');
        } else {
            console.warn('Welcome screen element not found');

            // welcome-screenが見つからない場合の対処
            if (lessonContent) {
                // 既存のlesson-viewを非表示にして、welcome-screenを再作成
                console.log('Recreating welcome screen');
                lessonContent.innerHTML = this.getWelcomeHTML();

                // 再作成後に welcome-screen を取得
                const recreatedWelcomeScreen = document.getElementById('welcome-screen');
                if (recreatedWelcomeScreen) {
                    console.log('Welcome screen recreated and displayed');
                }
            }
        }

        if (lessonView) {
            lessonView.style.display = 'none';
            console.log('Lesson view hidden');
        }

        // サイドバーのレッスンを非表示
        const lessonList = document.getElementById('lesson-list');
        if (lessonList) {
            lessonList.innerHTML = '<p style="color: #666; font-style: italic;">バージョンを選択してレッスンを表示</p>';
        }

        // コードエディターと出力をクリア
        if (this.tour.codeEditor) {
            this.tour.codeEditor.setValue('');
        } else {
            const codeEditor = document.getElementById('code-editor');
            if (codeEditor) {
                codeEditor.value = '';
            }
        }

        const output = document.getElementById('output');
        if (output) {
            output.textContent = '';
            output.className = '';
        }

        // バージョンセレクターをデフォルトに戻す
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            versionSelect.value = '1.25';
        }

        this.tour.currentLesson = null;
        this.tour.currentVersion = '1.25';

        // ウェルカム画面表示時はコードエディターと実行結果を非表示
        const codeSection = document.querySelector('.code-section');
        const outputSection = document.querySelector('.output-section');
        if (codeSection) {
            codeSection.style.display = 'none';
        }
        if (outputSection) {
            outputSection.style.display = 'none';
        }

        console.log('Welcome screen setup complete');
    }

    getWelcomeHTML() {
        return `
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
                                <span class="stars">⭐⭐⭐⭐</span>
                                <span>testing.B Loop</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐</span>
                                <span>os.Root</span>
                            </div>
                            <div class="feature-more">+ 4つの機能</div>
                        </div>
                        <button class="start-learning-btn" data-version="1.24">学習を開始</button>
                    </div>

                    <div class="version-card clickable" data-version="1.23">
                        <div class="version-header">
                            <h3>Go 1.23</h3>
                            <span class="version-badge">安定版</span>
                        </div>
                        <div class="version-features">
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐⭐</span>
                                <span>Structured Logging (slog)</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐⭐</span>
                                <span>Iterators</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐</span>
                                <span>Timer Reset</span>
                            </div>
                            <div class="feature-more">+ 3つの機能</div>
                        </div>
                        <button class="start-learning-btn" data-version="1.23">学習を開始</button>
                    </div>

                    <div class="version-card clickable" data-version="1.22">
                        <div class="version-header">
                            <h3>Go 1.22</h3>
                            <span class="version-badge">基盤版</span>
                        </div>
                        <div class="version-features">
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐⭐</span>
                                <span>For Range over Integers</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐⭐</span>
                                <span>Enhanced Loop Variables</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐</span>
                                <span>math/rand/v2</span>
                            </div>
                            <div class="feature-more">+ 2つの機能</div>
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
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐⭐</span>
                                <span>Slice to Array Conversion</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐</span>
                                <span>errors.Join</span>
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
                                <span>Memory Arenas</span>
                            </div>
                            <div class="feature-item">
                                <span class="stars">⭐⭐⭐</span>
                                <span>Atomic Types</span>
                            </div>
                        </div>
                        <button class="start-learning-btn" data-version="1.19">学習を開始</button>
                    </div>

                    <div class="version-card clickable" data-version="1.18">
                        <div class="version-header">
                            <h3>Go 1.18</h3>
                            <span class="version-badge breakthrough">革新版</span>
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
            </div>
        `;
    }

    async startLearning(version) {
        console.log(`Starting learning for Go ${version}`);

        this.tour.currentVersion = version;

        // バージョンセレクターを更新
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            versionSelect.value = version;
        }

        // ローディング状態を表示
        const lessonList = document.getElementById('lesson-list');
        lessonList.innerHTML = '<p style="color: #666; font-style: italic;">レッスンを読み込み中...</p>';

        // 選択されたバージョンのレッスンを読み込み
        await this.tour.loadLessons(version);
        this.tour.renderLessonList();

        // ウェルカム画面を非表示、レッスンビューを表示
        const welcomeScreen = document.getElementById('welcome-screen');
        const lessonView = document.getElementById('lesson-view');

        if (welcomeScreen) {
            welcomeScreen.style.display = 'none';
        }
        if (lessonView) {
            lessonView.style.display = 'block';
        }

        // コードエディターと実行結果セクションを表示
        const codeSection = document.querySelector('.code-section');
        const outputSection = document.querySelector('.output-section');
        if (codeSection) {
            codeSection.style.display = 'block';
        }
        if (outputSection) {
            outputSection.style.display = 'block';
        }

        // 最初のレッスンを自動選択
        const lessons = this.tour.lessons[version] || [];
        if (lessons.length > 0) {
            this.tour.selectLesson(lessons[0].id, version);
        }
    }
}

// WelcomeScreenをGoReleaseTourに統合
GoReleaseTour.prototype.showWelcomeScreen = function() {
    if (!this.welcomeScreen) {
        this.welcomeScreen = new WelcomeScreen(this);
    }
    this.welcomeScreen.show();
};

GoReleaseTour.prototype.startLearning = function(version) {
    if (!this.welcomeScreen) {
        this.welcomeScreen = new WelcomeScreen(this);
    }
    return this.welcomeScreen.startLearning(version);
};