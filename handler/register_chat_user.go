package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegisterChatUser struct {
	Service RegisterChatUserService
}

func NewRegisterChatUser(rcu RegisterChatUserService) *RegisterChatUser {
	return &RegisterChatUser{Service: rcu}
}

// チャットユーザー登録
func (rcu *RegisterChatUser) ServerHTTP(ctx *gin.Context) {
	// エラータイトル
	const errTitle = "ユーザー登録エラー"
	// クライアント リクエスト情報
	var input struct {
		Name string `json:"name"`
	}

	// クライアントから正しいパラメータでデータが送られていない
	if err := ctx.ShouldBindJSON(&input); err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errTitle, err.Error())
		return
	}
	// バリデーションチェック
	err := validation.ValidateStruct(&input,
		validation.Field(
			&input.Name,
			validation.Required,
			validation.Length(1, constant.ChatUserNameMaxLength),
		),
	)
	if err != nil {
		APIErrorResponse(ctx, http.StatusBadRequest, errTitle, err.Error())
		return
	}

	// サービス層へ依頼する
	userID, err := rcu.Service.RegisterChatUser(ctx, input.Name)
	if err != nil {
		// TODO: サービス層が返すエラーによってはエラーハンドリングを行う
		APIErrorResponse(ctx, http.StatusInternalServerError, errTitle, err.Error())
		return
	}

	// 成功
	res := struct {
		ID string `json:"userId"`
	}{ID: userID}
	APIResponse(ctx, http.StatusCreated, "ユーザー登録を行いました。", res)
}
