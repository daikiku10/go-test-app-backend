package service

import "fmt"

type ServicePostRegisterUserInput struct {
	TemporaryUserID string
	ConfirmCode     string
}
type PostRegisterUser struct{}

func NewPostRegisterUser() *PostRegisterUser {
	return &PostRegisterUser{}
}

func (r *PostRegisterUser) PostRegisterUser(input ServicePostRegisterUserInput) {
	fmt.Println("サービス層：ユーザー登録API")
}
