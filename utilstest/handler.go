package utilstest

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// テスト用
// アサーション関数(JSON)
// @params
// t テストユーティリティ
// want 検証のレスポンス
// got 実際のレスポンス
func AssertJSON(t *testing.T, want, got []byte) {
	// テスト関数がエラーの発生源であることを示すために指定する
	t.Helper()

	// interface{}
	var jw, jg any
	// JSONデータを構造体に変換する
	// wantを変換
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}
	// gotを変換
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarshal got %q: %v", got, err)
	}

	// wantとgotの差分
	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

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
	// JSONでの検証
	AssertJSON(t, body, gb)
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
