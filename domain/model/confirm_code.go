package model

import (
	"math/rand"
	"time"

	"github.com/daikiku10/go-test-app-backend/constant"
)

type ConfirmCode struct {
	value []byte
}

// ランダム文字列の確認コードを作成する
func NewConfirmCode() *ConfirmCode {
	// どの文字列からランダム文字列を生成するか
	const letters = "0123456789"

	// 異なる乱数を出すためにシードを与える
	rand.Seed(time.Now().UnixNano())
	value := make([]byte, constant.ConfirmCodeLength)
	for i := range value {
		r := rand.Int63() % int64(len(letters))
		value[i] = letters[int(r)]
	}

	return &ConfirmCode{value: value}
}

// 文字列の確認コードを返す
// @return
// 確認コード
func (cc *ConfirmCode) String() string {
	return string(cc.value)
}
