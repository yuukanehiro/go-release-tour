// Go Release Tour - メインエントリーポイント
// このファイルはモジュールを初期化し、アプリケーションを開始します

// グローバル変数
let tour;

// DOMが読み込まれた後に初期化
document.addEventListener('DOMContentLoaded', function() {
    console.log('DOM loaded, initializing Go Release Tour...');

    // Go Release Tour アプリケーションを初期化
    tour = new GoReleaseTour();

    // グローバルでアクセス可能にする（後方互換性のため）
    window.tour = tour;

    console.log('Go Release Tour initialized successfully');
});

// モジュール読み込み完了のチェック
function checkModulesLoaded() {
    const requiredClasses = [
        'GoReleaseTour',
        'ApiClient',
        'EditorManager',
        'NavigationManager',
        'WelcomeScreen',
        'LessonDisplay'
    ];

    const missingClasses = requiredClasses.filter(className => typeof window[className] === 'undefined');

    if (missingClasses.length > 0) {
        console.error('Missing required classes:', missingClasses);
        return false;
    }

    return true;
}

// エラーハンドリング
window.addEventListener('error', function(event) {
    console.error('JavaScript Error:', event.error);

    // ユーザーにエラーを表示
    const output = document.getElementById('output');
    if (output) {
        output.textContent = `アプリケーションエラーが発生しました: ${event.error.message}`;
        output.className = 'error';
    }
});

// 未処理の Promise エラーをキャッチ
window.addEventListener('unhandledrejection', function(event) {
    console.error('Unhandled Promise Rejection:', event.reason);

    // ユーザーにエラーを表示
    const output = document.getElementById('output');
    if (output) {
        output.textContent = `ネットワークエラーが発生しました: ${event.reason}`;
        output.className = 'error';
    }
});