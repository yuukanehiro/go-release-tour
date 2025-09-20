// レッスン表示管理モジュール
class LessonDisplay {
    constructor(tour) {
        this.tour = tour;
    }

    showLessonContent(lesson) {
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

        // 参考リンクを表示
        this.displayLessonLinks(lesson);

        // レッスンコードを読み込み
        const codeEditor = document.getElementById('code-editor');
        if (codeEditor) {
            // まず保存されたコードをチェック
            const savedCode = localStorage.getItem(`lesson-${lesson.version}-${lesson.id}-code`);
            codeEditor.value = savedCode || lesson.code || '';
        }

        // 出力をクリア
        const output = document.getElementById('output');
        if (output) {
            output.textContent = '';
            output.className = '';
        }
    }

    renderLessonList() {
        const lessonList = document.getElementById('lesson-list');
        const currentLessons = this.tour.lessons[this.tour.currentVersion] || [];

        console.log(`Rendering ${currentLessons.length} lessons for version ${this.tour.currentVersion}`);

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

        const lessons = this.tour.lessons[version] || [];
        const lesson = lessons.find(l => l.id === lessonId);
        if (!lesson) return;

        this.tour.currentLesson = lesson;

        // まだ表示されていなければレッスンビューを表示
        const welcomeScreen = document.getElementById('welcome-screen');
        const lessonView = document.getElementById('lesson-view');

        if (welcomeScreen) {
            welcomeScreen.style.display = 'none';
        }
        if (lessonView) {
            lessonView.style.display = 'block';
        }

        // コードエディターと実行結果セクションを表示（学習開始時）
        const codeSection = document.querySelector('.code-section');
        const outputSection = document.querySelector('.output-section');
        if (codeSection) {
            codeSection.style.display = 'block';
        }
        if (outputSection) {
            outputSection.style.display = 'block';
        }

        this.showLessonContent(lesson);
        this.tour.loadCodeIntoEditor(lesson);
    }

    displayLessonLinks(lesson) {
        const linksContainer = document.getElementById('lesson-links');
        const releaseNotesLink = document.getElementById('release-notes-link');
        const goDocLink = document.getElementById('go-doc-link');
        const proposalLink = document.getElementById('proposal-link');

        if (!linksContainer) return;

        // コードからリンク情報を抽出
        const links = this.extractLinksFromCode(lesson);

        if (links.releaseNotes) {
            releaseNotesLink.href = links.releaseNotes;
            releaseNotesLink.style.display = 'inline-flex';
        } else {
            // デフォルトのリリースノートリンク
            releaseNotesLink.href = `https://go.dev/doc/go${lesson.version}`;
            releaseNotesLink.style.display = 'inline-flex';
        }

        if (links.goDoc) {
            goDocLink.href = links.goDoc;
            goDocLink.style.display = 'inline-flex';
        } else {
            goDocLink.style.display = 'none';
        }

        if (links.proposal) {
            proposalLink.href = links.proposal;
            proposalLink.style.display = 'inline-flex';
        } else {
            proposalLink.style.display = 'none';
        }

        linksContainer.style.display = 'block';
    }

    extractLinksFromCode(lesson) {
        const links = {
            releaseNotes: null,
            goDoc: null,
            proposal: null
        };

        if (!lesson.code) return links;

        const lines = lesson.code.split('\n');
        for (const line of lines) {
            if (line.includes('Go 1.') && line.includes('Release Notes:')) {
                const urlMatch = line.match(/https:\/\/[^\s]+/);
                if (urlMatch) {
                    links.releaseNotes = urlMatch[0];
                }
            } else if (line.includes('Package:') || line.includes('Documentation:')) {
                const urlMatch = line.match(/https:\/\/pkg\.go\.dev\/[^\s]+/);
                if (urlMatch) {
                    links.goDoc = urlMatch[0];
                }
            } else if (line.includes('Proposal:')) {
                const urlMatch = line.match(/https:\/\/[^\s]+/);
                if (urlMatch) {
                    links.proposal = urlMatch[0];
                }
            }
        }

        return links;
    }
}

// LessonDisplayをGoReleaseTourに統合
GoReleaseTour.prototype.showLessonContent = function(lesson) {
    if (!this.lessonDisplay) {
        this.lessonDisplay = new LessonDisplay(this);
    }
    this.lessonDisplay.showLessonContent(lesson);
};

GoReleaseTour.prototype.renderLessonList = function() {
    if (!this.lessonDisplay) {
        this.lessonDisplay = new LessonDisplay(this);
    }
    this.lessonDisplay.renderLessonList();
};

GoReleaseTour.prototype.selectLesson = function(lessonId, version) {
    if (!this.lessonDisplay) {
        this.lessonDisplay = new LessonDisplay(this);
    }
    this.lessonDisplay.selectLesson(lessonId, version);
};