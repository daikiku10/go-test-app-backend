package model

import (
	"fmt"
	"regexp"
)

type TemporaryUserString struct {
	value string
}

func NewTemporaryUserString(temporaryUserString string) *TemporaryUserString {
	return &TemporaryUserString{value: temporaryUserString}
}

// ユーザー情報を改行で区切り、1つの文字列に結合する
//
// @param
// ユーザー情報
//
// @returns
// 連結したユーザー情報
func (tus *TemporaryUserString) Join(firstName, firstNameKana, familyName, familyNameKana, email, password string) string {
	tus.value = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", firstName, firstNameKana, familyName, familyNameKana, email, password)
	return tus.value
}

// 改行で区切られた1つの文字列になっているユーザー情報を分解する
//
// @param
// 1つの文字列になったユーザー情報
//
// @returns
// 分解されたユーザー情報
func (tus *TemporaryUserString) Split() (firstName, firstNameKana, familyName, familyNameKana, email, password string) {
	reg := "\r\n|\n"
	// 第二引数は回数指定で負を指定すればすべて確認する。
	arr := regexp.MustCompile(reg).Split(tus.value, -1)
	return arr[0], arr[1], arr[2], arr[3], arr[4], arr[5]

}
