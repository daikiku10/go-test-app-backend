package handler

import (
	"errors"
	"net/http"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/daikiku10/go-test-app-backend/service"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

	// クライアントから正しいパラメータでデータが送られていない
	if err := ctx.ShouldBindJSON(&input); err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errorTitle, err.Error())
		return
	}
	// バリデーションチェック
	err := validation.ValidateStruct(&input,
		validation.Field(
			&input.FirstName,
			validation.Required,
			validation.Length(1, constant.UserNameMaxLength),
		),
		validation.Field(
			&input.FamilyNameKana,
			validation.Required,
			validation.Length(1, constant.UserNameMaxLength),
		),
		validation.Field(
			&input.FamilyName,
			validation.Required,
			validation.Length(1, constant.UserNameMaxLength),
		),
		validation.Field(
			&input.FirstNameKana,
			validation.Required,
			validation.Length(1, constant.UserNameMaxLength),
		),
		validation.Field(
			&input.Password,
			validation.Required,
			validation.Length(1, constant.PasswordMaxLength),
		),
		validation.Field(
			&input.Email,
			validation.Required,
			validation.Length(1, constant.MailAddressMaxLength),
		))
	if err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errorTitle, err.Error())
		return
	}

	// サービス層のInputを作成する
	sInput := service.ServiceRegisterTemporaryUserInput{
		FirstName:      input.FirstName,
		FamilyName:     input.FamilyName,
		FirstNameKana:  input.FirstNameKana,
		FamilyNameKana: input.FamilyNameKana,
		Email:          input.Email,
		Password:       input.Password,
	}

	// サービス層へ依頼
	sessionID, err := rtu.Service.RegisterTemporaryUser(ctx, sInput)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyEntry) {
			APIErrorResponse(ctx, http.StatusConflict, errorTitle, repository.ErrAlreadyEntry.Error())
			return
		}
		APIErrorResponse(ctx, http.StatusInternalServerError, errorTitle, err.Error())
		return
	}

	// 成功
	rsp := struct {
		ID string `json:"temporaryUserId"`
	}{ID: sessionID}
	APIResponse(ctx, http.StatusCreated, "本登録メールを送信しました。", rsp)
}
