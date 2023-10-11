package utilstest

import (
	"io"
	"net/http"
	"os"
	"testing"
)

// テスト用
// アサーション関数
// @params
// t テストユーティリティ
// got 実際のレスポンス
// status  wantレスポンスステータスコード
// body wantレスポンスボディ
func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	// テスト関数がエラーの発生源であることを示すために指定する
	t.Helper()

	t.Cleanup(func() { _ = got.Body.Close() })
	// データを読み取りバイトスライスで返す
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	// 期待としても実体としてもレスポンスボディがない場合
	if len(gb) == 0 && len(body) == 0 {
		// AssertJSONを呼ぶ必要はない。
		return
	}
	// AssertJSON(t, body, gb)
}

// テスト用
// JSONファイルの読み込み
func LoadFile(t *testing.T, path string) []byte {
	// テスト関数がエラーの発生源であることを示すために指定する
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %v", path, err)
	}
	return bt
}
