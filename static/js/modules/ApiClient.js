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

        if (!code.trim()) {
            this.tour.showError('コードを入力してください');
            return;
        }

        // ローディング状態を表示
        runBtn.disabled = true;
        runBtn.textContent = '▶ 実行中...';
        output.textContent = '実行中...';
        output.className = '';

        try {
            // 標準のAPIエンドポイントを使用
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
            this.tour.showError(`コードの実行に失敗しました: ${error.message}`);
        } finally {
            runBtn.disabled = false;
            runBtn.textContent = '▶ 実行';
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