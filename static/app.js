// Go Release Tour - Frontend JavaScript
class GoReleaseTour {
    constructor() {
        this.lessons = {}; // version -> lessons
        this.currentVersion = '1.25';
        this.currentLesson = null;
        this.init();
    }

    async init() {
        console.log('Initializing Go Release Tour...');

        // Always populate version selector
        this.populateVersionSelector();

        // Setup event listeners first
        this.setupEventListeners();

        // Show welcome screen by default
        this.showWelcomeScreen();

        console.log('Initialization complete');
    }


    async loadLessons(version) {
        try {
            // Use standard API endpoint (not version-specific)
            const response = await fetch(`/api/lessons?version=${version}`);
            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }
            const lessons = await response.json();
            this.lessons[version] = lessons;
            console.log(`Loaded ${lessons.length} lessons for version ${version}`);
        } catch (error) {
            console.error('Failed to load lessons:', error);
            this.showError(`レッスンの読み込みに失敗しました: ${error.message}`);
        }
    }

    populateVersionSelector() {
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            // Always show both versions
            const allVersions = ['1.24', '1.25'];
            versionSelect.innerHTML = allVersions.map(version =>
                `<option value="${version}" ${version === this.currentVersion ? 'selected' : ''}>
                    Go ${version}
                </option>`
            ).join('');
        }
    }

    renderLessonList() {
        const lessonList = document.getElementById('lesson-list');
        const currentLessons = this.lessons[this.currentVersion] || [];

        console.log(`Rendering ${currentLessons.length} lessons for version ${this.currentVersion}`);

        if (currentLessons.length === 0) {
            lessonList.innerHTML = '<p>レッスンを読み込み中...</p>';
            return;
        }

        lessonList.innerHTML = currentLessons.map(lesson => `
            <div class="lesson-item" data-lesson-id="${lesson.id}" data-version="${lesson.version}" onclick="tour.selectLesson(${lesson.id}, '${lesson.version}')">
                <div class="lesson-title">${lesson.title}</div>
                <div class="lesson-description">${lesson.description}</div>
                <div class="lesson-stars">${this.renderStars(lesson.stars)}</div>
            </div>
        `).join('');
    }

    renderStars(count) {
        return '★'.repeat(count) + '☆'.repeat(5 - count);
    }

    selectLesson(lessonId, version) {
        // Remove active class from all lessons
        document.querySelectorAll('.lesson-item').forEach(item => {
            item.classList.remove('active');
        });

        // Add active class to selected lesson
        const selectedItem = document.querySelector(`[data-lesson-id="${lessonId}"][data-version="${version}"]`);
        if (selectedItem) {
            selectedItem.classList.add('active');
        }

        const lessons = this.lessons[version] || [];
        const lesson = lessons.find(l => l.id === lessonId);
        if (!lesson) return;

        this.currentLesson = lesson;

        // Show lesson view if not already visible
        const welcomeScreen = document.getElementById('welcome-screen');
        const lessonView = document.getElementById('lesson-view');

        if (welcomeScreen) {
            welcomeScreen.style.display = 'none';
        }
        if (lessonView) {
            lessonView.style.display = 'block';
        }

        this.showLessonContent(lesson);
        this.loadCodeIntoEditor(lesson);
    }

    loadCodeIntoEditor(lesson) {
        const editor = document.getElementById('code-editor');
        if (!editor) return;

        // Remove build ignore directives but keep all code including type definitions
        let code = lesson.code;

        // Remove build ignore directives
        code = code.replace(/\/\/go:build ignore\n/g, '');
        code = code.replace(/\/\/ \+build ignore\n/g, '');

        // Load complete code to preserve type definitions and structure
        editor.value = code;
    }

    async changeVersion(version) {
        this.currentVersion = version;

        // Load lessons for this version if not already loaded
        if (!this.lessons[version]) {
            await this.loadLessons(version);
        }

        this.renderLessonList();

        // Clear current lesson selection
        this.currentLesson = null;
        const content = document.getElementById('lesson-content');
        content.innerHTML = `
            <h2>Go ${version} Features</h2>
            <p>左のレッスンリストからGo ${version}の新機能を選択してください。</p>
        `;

        // Clear editor
        document.getElementById('code-editor').value = '';
        document.getElementById('output').textContent = '';
    }

    setupEventListeners() {
        // Version selector
        const versionSelect = document.getElementById('version-select');
        versionSelect.addEventListener('change', (e) => {
            this.changeVersion(e.target.value);
        });

        // Code editor improvements
        const editor = document.getElementById('code-editor');

        // Add tab support
        editor.addEventListener('keydown', (e) => {
            if (e.key === 'Tab') {
                e.preventDefault();
                const start = editor.selectionStart;
                const end = editor.selectionEnd;
                editor.value = editor.value.substring(0, start) + '\t' + editor.value.substring(end);
                editor.selectionStart = editor.selectionEnd = start + 1;
            }
        });

        // Auto-save to localStorage
        editor.addEventListener('input', () => {
            if (this.currentLesson) {
                localStorage.setItem(`lesson-${this.currentLesson.version}-${this.currentLesson.id}-code`, editor.value);
            }
        });

        // Load saved code when selecting lesson
        const originalSelectLesson = this.selectLesson.bind(this);
        this.selectLesson = (lessonId, version) => {
            console.log(`Selecting lesson ${lessonId} for version ${version}`);
            originalSelectLesson(lessonId, version);
            const savedCode = localStorage.getItem(`lesson-${version}-${lessonId}-code`);
            if (savedCode) {
                document.getElementById('code-editor').value = savedCode;
            }
        };

        // Version selection from welcome screen
        document.addEventListener('click', (e) => {
            if (e.target.classList.contains('start-learning-btn')) {
                const version = e.target.dataset.version;
                this.startLearning(version);
            }
        });

        // Back to versions button
        const backToVersionsBtn = document.getElementById('back-to-versions');
        if (backToVersionsBtn) {
            backToVersionsBtn.addEventListener('click', () => {
                this.showWelcomeScreen();
            });
        }
    }

    async runCode() {
        const code = document.getElementById('code-editor').value;
        const output = document.getElementById('output');
        const runBtn = document.getElementById('run-btn');

        if (!code.trim()) {
            this.showError('コードを入力してください');
            return;
        }

        // Show loading state
        runBtn.disabled = true;
        runBtn.textContent = '▶ 実行中...';
        output.textContent = '実行中...';
        output.className = '';

        try {
            // Use standard API endpoint
            const response = await fetch('/api/run', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ code }),
            });

            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }

            const result = await response.json();

            if (result.error) {
                output.textContent = `エラー: ${result.error}\n\n出力:\n${result.output}`;
                output.className = 'error';
            } else {
                output.textContent = result.output || '実行完了（出力なし）';
                output.className = '';
            }
        } catch (error) {
            console.error('Execution error:', error);
            this.showError(`コードの実行に失敗しました: ${error.message}`);
        } finally {
            runBtn.disabled = false;
            runBtn.textContent = '▶ 実行';
        }
    }

    showError(message) {
        const output = document.getElementById('output');
        output.textContent = message;
        output.className = 'error';
    }
}

// Global function for button onclick
function runCode() {
    if (window.tour) {
        window.tour.runCode();
    }
}

// Initialize the application
document.addEventListener('DOMContentLoaded', () => {
    window.tour = new GoReleaseTour();
});

// Keyboard shortcuts
document.addEventListener('keydown', (e) => {
    // Ctrl/Cmd + Enter to run code
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        e.preventDefault();
        runCode();
    }

    // Ctrl/Cmd + S to save (prevent default browser save)
    if ((e.ctrlKey || e.metaKey) && e.key === 's') {
        e.preventDefault();
        // Code is auto-saved, just show a brief indicator
        const editor = document.getElementById('code-editor');
        const originalBorder = editor.style.border;
        editor.style.border = '2px solid #28a745';
        setTimeout(() => {
            editor.style.border = originalBorder;
        }, 200);
    }
});

// Add navigation methods to GoReleaseTour class prototype
GoReleaseTour.prototype.showWelcomeScreen = function() {
    const welcomeScreen = document.getElementById('welcome-screen');
    const lessonView = document.getElementById('lesson-view');

    if (welcomeScreen) {
        welcomeScreen.style.display = 'block';
    }
    if (lessonView) {
        lessonView.style.display = 'none';
    }

    // Hide sidebar lessons
    const lessonList = document.getElementById('lesson-list');
    if (lessonList) {
        lessonList.innerHTML = '<p style="color: #666; font-style: italic;">バージョンを選択してレッスンを表示</p>';
    }

    // Clear code editor and output
    const codeEditor = document.getElementById('code-editor');
    const output = document.getElementById('output');

    if (codeEditor) {
        codeEditor.value = '';
    }
    if (output) {
        output.textContent = '';
    }

    this.currentLesson = null;
};

GoReleaseTour.prototype.startLearning = async function(version) {
    console.log(`Starting learning for Go ${version}`);

    this.currentVersion = version;

    // Update version selector
    const versionSelect = document.getElementById('version-select');
    if (versionSelect) {
        versionSelect.value = version;
    }

    // Show loading state
    const lessonList = document.getElementById('lesson-list');
    lessonList.innerHTML = '<p style="color: #666; font-style: italic;">レッスンを読み込み中...</p>';

    // Load lessons for selected version
    await this.loadLessons(version);
    this.renderLessonList();

    // Hide welcome screen, show lesson view
    const welcomeScreen = document.getElementById('welcome-screen');
    const lessonView = document.getElementById('lesson-view');

    if (welcomeScreen) {
        welcomeScreen.style.display = 'none';
    }
    if (lessonView) {
        lessonView.style.display = 'block';
    }

    // Auto-select first lesson if available
    if (this.lessons[version] && this.lessons[version].length > 0) {
        const firstLesson = this.lessons[version][0];
        this.selectLesson(firstLesson.id, version);
    }
};

GoReleaseTour.prototype.showLessonContent = function(lesson) {
    const lessonTitle = document.getElementById('current-lesson-title');
    const lessonStars = document.getElementById('current-lesson-stars');
    const lessonDescription = document.getElementById('lesson-description');

    if (lessonTitle) {
        lessonTitle.textContent = lesson.title;
    }

    if (lessonStars) {
        lessonStars.innerHTML = '★'.repeat(lesson.stars) + '☆'.repeat(5 - lesson.stars);
    }

    if (lessonDescription) {
        lessonDescription.innerHTML = `<p>${lesson.description}</p>`;
    }

    // Load lesson code
    const codeEditor = document.getElementById('code-editor');
    if (codeEditor) {
        // Check for saved code first
        const savedCode = localStorage.getItem(`lesson-${lesson.version}-${lesson.id}-code`);
        codeEditor.value = savedCode || lesson.code || '';
    }

    // Clear output
    document.getElementById('output').textContent = '';
};