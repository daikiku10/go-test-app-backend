package model

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

// パスワードオブジェクト作成
// ハッシュ化されてない値を扱う
// コンストラクタ
//
// @params pwd パスワード
//
// @return パスワードオブジェクト
func NewPassword(pwd string) (*Password, error) {
	// バリデーションチェック
	if 50 < utf8.RuneCountInString(pwd) {
		return nil, fmt.Errorf("cannot use password over 51 char")
	}
	return &Password{value: pwd}, nil
}

// ハッシュ化パスワードの作成
func (pwd *Password) CreateHash() (string, error) {
	// パスワードをハッシュ化
	// 第2引数はコストを指定する。値は4 ~ 31の範囲である必要がある。
	password, err := bcrypt.GenerateFromPassword([]byte(pwd.value), bcrypt.DefaultCost)
	return string(password), err
}
