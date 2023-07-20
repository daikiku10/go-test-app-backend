package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterTemporaryUser struct {
	Service RegisterTemporaryUserService
}

func NewRegisterTemporaryUser(rtu RegisterTemporaryUserService) *RegisterTemporaryUser {
	return &RegisterTemporaryUser{Service: rtu}
}

func (rtu *RegisterTemporaryUser) ServerHTTP(ctx *gin.Context) {
	fmt.Println("ハンドラー層：仮ユーザー登録API")
	rtu.Service.RegisterTemporaryUser()
}
