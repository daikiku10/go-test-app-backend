package model

import (
	"fmt"
	"unicode/utf8"

	"github.com/daikiku10/go-test-app-backend/constant"
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
	if constant.PasswordMaxLength < utf8.RuneCountInString(pwd) {
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
