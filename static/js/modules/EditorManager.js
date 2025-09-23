// CodeMirrorエディター管理モジュール
class EditorManager {
    constructor(tour) {
        this.tour = tour;
    }

    initCodeEditor() {
        const textArea = document.getElementById('code-editor');
        console.log('Initializing CodeMirror...', { textArea, CodeMirror: typeof CodeMirror });

        if (!textArea) {
            console.warn('Code editor textarea not found');
            return;
        }

        if (typeof CodeMirror === 'undefined') {
            console.warn('CodeMirror not loaded, falling back to textarea');
            return;
        }

        try {
            this.tour.codeEditor = CodeMirror.fromTextArea(textArea, {
                mode: 'text/x-go',
                theme: 'monokai',
                lineNumbers: true,
                indentUnit: 4,
                indentWithTabs: true,
                autoCloseBrackets: true,
                matchBrackets: true,
                viewportMargin: Infinity,
                scrollbarStyle: "null",
                lineWrapping: true,
                extraKeys: {
                    'Ctrl-Enter': () => this.tour.runCode(),
                    'Cmd-Enter': () => this.tour.runCode(),
                    'Tab': (cm) => {
                        if (cm.somethingSelected()) {
                            cm.indentSelection('add');
                        } else {
                            cm.replaceSelection('\t');
                        }
                    }
                }
            });


            // テーマセレクターの初期化
            this.setupThemeSelector();

            // エディターをリフレッシュして全コンテンツを表示
            setTimeout(() => {
                this.tour.codeEditor.refresh();
                this.tour.codeEditor.setSize(null, "auto");
            }, 100);

            console.log('CodeMirror editor initialized successfully');
        } catch (error) {
            console.error('Failed to initialize CodeMirror:', error);
        }
    }

    setupThemeSelector() {
        const themeSelect = document.getElementById('theme-selector');
        if (!themeSelect || !this.tour.codeEditor) return;

        themeSelect.addEventListener('change', (e) => {
            this.tour.codeEditor.setOption('theme', e.target.value);
            localStorage.setItem('code-editor-theme', e.target.value);
        });

        // 保存されたテーマを読み込み
        const savedTheme = localStorage.getItem('code-editor-theme');
        if (savedTheme) {
            themeSelect.value = savedTheme;
            this.tour.codeEditor.setOption('theme', savedTheme);
        }
    }

    loadCodeIntoEditor(lesson) {
        // ビルド無視ディレクティブを削除、型定義を含む全コードは保持
        let code = lesson.code;

        // ビルド無視ディレクティブを削除
        code = code.replace(/\/\/go:build ignore\n/g, '');
        code = code.replace(/\/\/ \+build ignore\n/g, '');

        // CodeMirrorエディターまたは通常のtextareaを使用
        if (this.tour.codeEditor) {
            this.tour.codeEditor.setValue(code);
        } else {
            // フォールバック: 通常のtextarea
            const editor = document.getElementById('code-editor');
            if (editor) {
                editor.value = code;
            }
        }
    }

    setupTextareaFallback() {
        // コードエディターの改善（CodeMirrorが使用できない場合のフォールバック）
        if (!this.tour.codeEditor) {
            const editor = document.getElementById('code-editor');
            if (editor) {
                // タブサポートを追加
                editor.addEventListener('keydown', (e) => {
                    if (e.key === 'Tab') {
                        e.preventDefault();
                        const start = editor.selectionStart;
                        const end = editor.selectionEnd;
                        editor.value = editor.value.substring(0, start) + '\t' + editor.value.substring(end);
                        editor.selectionStart = editor.selectionEnd = start + 1;
                    }
                });

            }
        }
    }
}

// EditorManagerをGoReleaseTourに統合
GoReleaseTour.prototype.initCodeEditor = function() {
    if (!this.editorManager) {
        this.editorManager = new EditorManager(this);
    }
    this.editorManager.initCodeEditor();
};

GoReleaseTour.prototype.loadCodeIntoEditor = function(lesson) {
    if (!this.editorManager) {
        this.editorManager = new EditorManager(this);
    }
    this.editorManager.loadCodeIntoEditor(lesson);
};