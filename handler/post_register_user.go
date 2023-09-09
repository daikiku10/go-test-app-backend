package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/service"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PostRegisterUser struct {
	Service PostRegisterUserService
}

func NewPostRegisterUser(pru PostRegisterUserService) *PostRegisterUser {
	return &PostRegisterUser{Service: pru}
}

// ユーザー登録
//
// @param ctx ginContext
func (pru *PostRegisterUser) ServerHTTP(ctx *gin.Context) {
	// エラータイトル
	const errorTitle = "ユーザー登録エラー"
	// クライアント リクエスト情報
	var input struct {
		TemporaryUserID string `json:"temporaryUserId"`
		ConfirmCode     string `json:"confirmCode"`
	}

	// クライアントから正しいパラメータでデータが送られていない
	if err := ctx.ShouldBindJSON(&input); err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errorTitle, err.Error())
		return
	}
	// バリデーションチェック
	err := validation.ValidateStruct(&input,
		validation.Field(
			&input.TemporaryUserID,
			validation.Required,
		),
		validation.Field(
			&input.ConfirmCode,
			validation.Required,
		),
	)
	if err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errorTitle, err.Error())
		return
	}

	// サービス層のInputを作成する
	sInput := service.ServicePostRegisterUserInput{
		TemporaryUserID: input.TemporaryUserID,
		ConfirmCode:     input.ConfirmCode,
	}

	// サービス層へ依頼
	pru.Service.PostRegisterUser(ctx, sInput)

	// TODO: 成功
}
