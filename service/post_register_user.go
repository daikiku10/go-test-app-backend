package service

import "fmt"

type PostRegisterUser struct{}

func NewPostRegisterUser() *PostRegisterUser {
	return &PostRegisterUser{}
}

func (r *PostRegisterUser) PostRegisterUser() {
	fmt.Println("サービス層：ユーザー登録API")
}
