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

        // 環境変数プリセットを設定
        this.setupEnvPresets(lesson);

        // レッスンコードを読み込み
        const codeEditor = document.getElementById('code-editor');
        if (codeEditor) {
            codeEditor.value = lesson.code || '';
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

    setupEnvPresets(lesson) {
        const envPresetsContainer = document.getElementById('env-presets');
        const envInfoText = document.getElementById('env-info-text');

        if (!envPresetsContainer) return;

        // 既存のプリセット表示をクリア
        envPresetsContainer.innerHTML = '';

        // レッスンに環境変数プリセットがある場合
        if (lesson.env_presets && lesson.env_presets.length > 0) {
            // .env形式のテキストを表示
            const presetDisplay = document.createElement('div');
            presetDisplay.className = 'env-preset-display';

            const title = document.createElement('h5');
            title.textContent = '利用可能な環境変数:';
            title.style.margin = '0 0 0.5rem 0';
            title.style.fontSize = '0.9rem';
            title.style.color = '#495057';
            presetDisplay.appendChild(title);

            lesson.env_presets.forEach(preset => {
                const presetItem = document.createElement('div');
                presetItem.className = 'env-preset-item';
                presetItem.style.marginBottom = '0.75rem';
                presetItem.style.padding = '0.5rem';
                presetItem.style.background = '#f8f9fa';
                presetItem.style.border = '1px solid #e9ecef';
                presetItem.style.borderRadius = '4px';
                presetItem.style.cursor = 'pointer';
                presetItem.style.transition = 'background-color 0.2s';

                // プリセット名と説明
                const presetHeader = document.createElement('div');
                presetHeader.style.display = 'flex';
                presetHeader.style.justifyContent = 'space-between';
                presetHeader.style.alignItems = 'center';
                presetHeader.style.marginBottom = '0.25rem';

                const presetName = document.createElement('strong');
                presetName.textContent = preset.name;
                presetName.style.fontSize = '0.85rem';
                presetName.style.color = '#495057';

                const copyBtn = document.createElement('button');
                copyBtn.textContent = 'コピー';
                copyBtn.style.fontSize = '0.75rem';
                copyBtn.style.padding = '2px 6px';
                copyBtn.style.border = '1px solid #6c757d';
                copyBtn.style.background = 'white';
                copyBtn.style.borderRadius = '3px';
                copyBtn.style.cursor = 'pointer';

                presetHeader.appendChild(presetName);
                presetHeader.appendChild(copyBtn);

                // .env形式の値（複数行対応）
                const envValue = document.createElement('pre');

                // カンマ区切りの環境変数を複数行に変換
                const formatEnvVars = (envString) => {
                    if (envString.includes(',')) {
                        return envString.split(',')
                            .map(env => env.trim())
                            .filter(env => env.length > 0)
                            .join('\n');
                    }
                    return envString;
                };

                envValue.textContent = formatEnvVars(preset.value);
                envValue.style.display = 'block';
                envValue.style.fontSize = '0.8rem';
                envValue.style.background = '#e9ecef';
                envValue.style.padding = '0.5rem';
                envValue.style.borderRadius = '3px';
                envValue.style.marginBottom = '0.25rem';
                envValue.style.fontFamily = 'Monaco, Consolas, "Courier New", monospace';
                envValue.style.whiteSpace = 'pre';
                envValue.style.overflow = 'auto';
                envValue.style.margin = '0';
                envValue.style.lineHeight = '1.4';

                // 説明
                const description = document.createElement('small');
                description.textContent = preset.description;
                description.style.color = '#6c757d';
                description.style.fontSize = '0.75rem';

                presetItem.appendChild(presetHeader);
                presetItem.appendChild(envValue);
                presetItem.appendChild(description);

                // クリックでコピー
                const copyToInput = () => {
                    const envVarsInput = document.getElementById('env-vars');
                    if (envVarsInput) {
                        envVarsInput.value = preset.value;
                        envVarsInput.focus();

                        // 視覚的フィードバック
                        const originalText = copyBtn.textContent;
                        const originalBg = copyBtn.style.background;
                        copyBtn.style.background = '#28a745';
                        copyBtn.style.color = 'white';
                        copyBtn.textContent = '✓';

                        setTimeout(() => {
                            copyBtn.style.background = originalBg;
                            copyBtn.style.color = '';
                            copyBtn.textContent = originalText;
                        }, 1000);
                    }
                };

                copyBtn.addEventListener('click', (e) => {
                    e.stopPropagation();
                    copyToInput();
                });

                presetItem.addEventListener('click', copyToInput);

                // ホバー効果
                presetItem.addEventListener('mouseenter', () => {
                    presetItem.style.background = '#e9ecef';
                });
                presetItem.addEventListener('mouseleave', () => {
                    presetItem.style.background = '#f8f9fa';
                });

                presetDisplay.appendChild(presetItem);
            });

            envPresetsContainer.appendChild(presetDisplay);

            // 情報テキストを更新
            if (envInfoText) {
                envInfoText.textContent = `💡 ${lesson.env_presets.length}個の環境変数設定例が利用可能です（クリックでコピー）`;
            }
        } else {
            // プリセットがない場合
            if (envInfoText) {
                envInfoText.textContent = '💡 このレッスンには環境変数設定例がありません';
            }
        }
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