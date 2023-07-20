package service

import "fmt"

type RegisterTemporaryUser struct{}

func NewRegisterTemporaryUser() *RegisterTemporaryUser {
	return &RegisterTemporaryUser{}
}

func (rtu *RegisterTemporaryUser) RegisterTemporaryUser() {
	fmt.Println("サービス層：仮ユーザー登録API")
}
