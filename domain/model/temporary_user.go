package model

import "fmt"

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
