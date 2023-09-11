package handler

import (
	"errors"
	"net/http"

	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/repository"
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
	u, jwt, err := pru.Service.PostRegisterUser(ctx, sInput)
	if err != nil {
		// キャッシュ有効期限切れ
		if errors.Is(err, repository.ErrNotFoundSession) {
			APIErrorResponse(ctx, http.StatusUnauthorized, errorTitle, repository.ErrNotFoundSession.Error())
			return
		}
		// 同じメールアドレスが存在する
		if errors.Is(err, repository.ErrAlreadyEntry) {
			APIErrorResponse(ctx, http.StatusConflict, errorTitle, repository.ErrAlreadyEntry.Error())
			return
		}
		APIErrorResponse(ctx, http.StatusInternalServerError, errorTitle, err.Error())
		return
	}

	rsp := struct {
		ID    model.UserID `json:"userId"`
		Token string       `json:"accessToken"`
	}{ID: u.ID, Token: jwt}
	APIResponse(ctx, http.StatusCreated, "本登録が完了しました。", rsp)
}
