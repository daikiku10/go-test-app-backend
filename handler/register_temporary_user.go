package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterTemporaryUser struct {
	Service RegisterTemporaryUserService
}

func NewRegisterTemporaryUser(rtu RegisterTemporaryUserService) *RegisterTemporaryUser {
	return &RegisterTemporaryUser{Service: rtu}
}

// ユーザー仮登録ハンドラー
//
// @param ctx ginContext
func (rtu *RegisterTemporaryUser) ServerHTTP(ctx *gin.Context) {
	fmt.Println("ハンドラー層：仮ユーザー登録API")
	// エラータイトル
	const errorTitle = "ユーザー仮登録エラー"
	// クライアント リクエスト情報
	var input struct {
		FirstName      string `json:"firstName"`
		FirstNameKana  string `json:"firstNameKana"`
		FamilyName     string `json:"FamilyName"`
		FamilyNameKana string `json:"FamilyNameKana"`
		Password       string `json:"password"`
		Email          string `json:"email"`
	}

	fmt.Printf("%+v", &input)

	// クライアントから正しいパラメータでデータが送られていない
	if err := ctx.ShouldBindJSON(&input); err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errorTitle, err.Error())
		return
	}

	fmt.Printf("%+v", &input)
	rtu.Service.RegisterTemporaryUser()
}
