package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PostRegisterUser struct {
	Service PostRegisterUserService
}

func NewPostRegisterUser(pru PostRegisterUserService) *PostRegisterUser {
	return &PostRegisterUser{Service: pru}
}

func (pru *PostRegisterUser) ServerHTTP(ctx *gin.Context) {
	fmt.Println("ハンドラー層：ユーザー登録API")
	pru.Service.PostRegisterUser()

}
