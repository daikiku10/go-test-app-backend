package testing

import (
	"os"
	"testing"
)

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
