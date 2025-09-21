// Go Release Tour - メインクラス
class GoReleaseTour {
    constructor() {
        this.lessons = {}; // バージョン -> レッスン
        this.currentVersion = '1.25';
        this.currentLesson = null;
        this.codeEditor = null; // CodeMirrorエディターインスタンス
        this.isEditorExpanded = false; // エディター展開状態
        this.init();
    }

    async init() {
        console.log('Initializing Go Release Tour...');

        // HTMLに既にバージョンセレクターがあるので、populateVersionSelectorは不要
        // デフォルトバージョンをHTMLの選択値に合わせる
        const versionSelect = document.getElementById('version-select');
        if (versionSelect && versionSelect.value) {
            this.currentVersion = versionSelect.value;
        }

        // 最初にイベントリスナーを設定
        this.setupEventListeners();

        // CodeMirrorエディターを初期化（DOMが準備できてから）
        setTimeout(() => {
            this.initCodeEditor();
        }, 100);

        // デフォルトでウェルカム画面を表示
        this.showWelcomeScreen();

        console.log('Initialization complete');
    }

    populateVersionSelector() {
        const versionSelect = document.getElementById('version-select');
        if (versionSelect) {
            // 全バージョンを表示（新しい順）
            const allVersions = [
                { value: '1.25', label: 'Go 1.25 (最新)' },
                { value: '1.24', label: 'Go 1.24' },
                { value: '1.23', label: 'Go 1.23' },
                { value: '1.22', label: 'Go 1.22' },
                { value: '1.21', label: 'Go 1.21' },
                { value: '1.20', label: 'Go 1.20' },
                { value: '1.19', label: 'Go 1.19' },
                { value: '1.18', label: 'Go 1.18 (Generics)' }
            ];
            versionSelect.innerHTML = allVersions.map(version =>
                `<option value="${version.value}" ${version.value === this.currentVersion ? 'selected' : ''}>
                    ${version.label}
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
        // 全レッスンからactiveクラスを削除
        document.querySelectorAll('.lesson-item').forEach(item => {
            item.classList.remove('active');
        });

        // 選択したレッスンにactiveクラスを追加
        const selectedItem = document.querySelector(`[data-lesson-id="${lessonId}"][data-version="${version}"]`);
        if (selectedItem) {
            selectedItem.classList.add('active');
        }

        const lessons = this.lessons[version] || [];
        const lesson = lessons.find(l => l.id === lessonId);
        if (!lesson) return;

        this.currentLesson = lesson;

        // まだ表示されていなければレッスンビューを表示
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

    showError(message) {
        const output = document.getElementById('output');
        output.textContent = message;
        output.className = 'error';
    }

    // 展開機能は無効化されました
    // toggleEditorExpand() {
    //     // 機能無効化のため削除
    // }
}