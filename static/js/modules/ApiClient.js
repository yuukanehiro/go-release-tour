// API通信クライアント
class ApiClient {
    constructor(tour) {
        this.tour = tour;
    }

    async loadLessons(version) {
        try {
            // 標準のAPIエンドポイントを使用（バージョン固有ではない）
            const response = await fetch(`/api/lessons?version=${version}`);
            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }
            const lessons = await response.json();
            this.tour.lessons[version] = lessons;
            console.log(`Loaded ${lessons.length} lessons for version ${version}`);
        } catch (error) {
            console.error('Failed to load lessons:', error);
            this.tour.showError(`レッスンの読み込みに失敗しました: ${error.message}`);
        }
    }

    async runCode() {
        // CodeMirrorまたは通常のtextareaからコードを取得
        const code = this.tour.codeEditor ? this.tour.codeEditor.getValue() : document.getElementById('code-editor').value;
        const output = document.getElementById('output');
        const runBtn = document.getElementById('run-btn');
        const runBtnBottom = document.getElementById('run-btn-bottom');

        if (!code.trim()) {
            this.tour.showError('コードを入力してください');
            return;
        }

        // 現在のレッスン情報を取得
        const currentLesson = this.tour.currentLesson;
        const lessonPath = currentLesson?.file_path || currentLesson?.FilePath;

        // バージョンセレクターからバージョンを取得
        const versionSelect = document.getElementById('version-select');
        console.log('Debug: versionSelect element =', versionSelect);

        let selectedVersion = null;
        if (versionSelect && versionSelect.options && versionSelect.options.length > 0) {
            selectedVersion = versionSelect.value;
            console.log('Debug: versionSelect.value =', selectedVersion);
            console.log('Debug: versionSelect.selectedIndex =', versionSelect.selectedIndex);
            console.log('Debug: versionSelect.options =', Array.from(versionSelect.options).map(o => `${o.value}:${o.selected}`));

            // 値が空文字の場合、選択されたオプションから取得を試行
            if (!selectedVersion && versionSelect.selectedIndex >= 0) {
                selectedVersion = versionSelect.options[versionSelect.selectedIndex].value;
                console.log('Debug: Fallback to selectedIndex value =', selectedVersion);
            }
        } else {
            console.error('Debug: version-select element not found or has no options!');
            console.log('Debug: versionSelect =', versionSelect);
            console.log('Debug: versionSelect.options =', versionSelect ? versionSelect.options : 'N/A');
        }

        console.log('Debug: currentLesson =', currentLesson);
        console.log('Debug: selectedVersion =', selectedVersion);
        console.log('Debug: lessonPath =', lessonPath);

        // ローディング状態を表示
        runBtn.disabled = true;
        runBtn.textContent = '▶ 実行中...';
        if (runBtnBottom) {
            runBtnBottom.disabled = true;
            runBtnBottom.textContent = '▶ 実行中...';
        }
        output.textContent = '実行中...';
        output.className = '';

        try {
            // バージョン検出ロジック - 簡素化
            let detectedVersion = null;

            console.log('=== VERSION DETECTION START ===');
            console.log('selectedVersion:', selectedVersion);
            console.log('lessonPath:', lessonPath);
            console.log('code preview:', code.substring(0, 100));

            // 1. バージョンセレクターから取得
            if (selectedVersion && selectedVersion.trim()) {
                detectedVersion = selectedVersion.trim();
                console.log('✓ Using version from selector:', detectedVersion);
            }

            // 2. コードから検出
            if (!detectedVersion) {
                console.log('Attempting code detection...');
                const codeMatch = code.match(/Go\s+(\d+\.\d+)/);
                if (codeMatch) {
                    detectedVersion = codeMatch[1];
                    console.log('✓ Detected from code:', detectedVersion);
                }
            }

            // 3. レッスンパスから検出
            if (!detectedVersion && lessonPath) {
                console.log('Attempting path detection...');
                const pathMatch = lessonPath.match(/\/v\/(\d+\.\d+)\//);
                if (pathMatch) {
                    detectedVersion = pathMatch[1];
                    console.log('✓ Detected from path:', detectedVersion);
                }
            }

            // 4. デフォルト
            if (!detectedVersion) {
                detectedVersion = '1.25';
                console.log('✓ Using default version:', detectedVersion);
            }

            console.log('=== FINAL VERSION:', detectedVersion, '===');

            // 型安全なペイロード構築
            if (!detectedVersion) {
                throw new Error('バージョンが決定できませんでした');
            }

            const payload = {
                code: code,
                version: detectedVersion
            };

            // ペイロード検証
            if (!payload.code || !payload.version) {
                throw new Error(`無効なペイロード: code=${!!payload.code}, version=${!!payload.version}`);
            }

            console.log('Debug: Final payload =', JSON.stringify(payload, null, 2));

            // バージョン対応のAPIエンドポイントを使用
            const response = await fetch('/api/run', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            });

            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }

            const result = await response.json();

            // バージョン情報を表示
            let versionInfo = '';
            if (result.used_version || result.go_version) {
                versionInfo = `実行環境: Go ${result.go_version || result.used_version}`;
                if (result.detected_version && result.detected_version !== result.used_version) {
                    versionInfo += ` (検出: ${result.detected_version})`;
                }
                if (result.execution_time) {
                    versionInfo += ` | 実行時間: ${result.execution_time}`;
                }
                versionInfo += '\n' + '='.repeat(50) + '\n';
            }

            if (result.error) {
                output.textContent = versionInfo + `エラー: ${result.error}\n\n出力:\n${result.output}`;
                output.className = 'error';
            } else {
                output.textContent = versionInfo + (result.output || '実行完了（出力なし）');
                output.className = '';
            }
        } catch (error) {
            console.error('Execution error:', error);
            this.tour.showError(`コードの実行に失敗しました: ${error.message}`);
        } finally {
            runBtn.disabled = false;
            runBtn.textContent = '▶ 実行';
            if (runBtnBottom) {
                runBtnBottom.disabled = false;
                runBtnBottom.textContent = '▶ 実行';
            }
        }
    }
}

// ApiClientをGoReleaseTourに統合
GoReleaseTour.prototype.loadLessons = function(version) {
    if (!this.apiClient) {
        this.apiClient = new ApiClient(this);
    }
    return this.apiClient.loadLessons(version);
};

GoReleaseTour.prototype.runCode = function() {
    if (!this.apiClient) {
        this.apiClient = new ApiClient(this);
    }
    return this.apiClient.runCode();
};