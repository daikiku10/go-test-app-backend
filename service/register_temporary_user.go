package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/domain/service"
	"github.com/daikiku10/go-test-app-backend/repository"
)

type RegisterTemporaryUser struct {
	Repo domain.UserRepo
	DB   repository.Queryer
}

func NewRegisterTemporaryUser(rep domain.UserRepo, db repository.Queryer) *RegisterTemporaryUser {
	return &RegisterTemporaryUser{Repo: rep, DB: db}
}

// ユーザー仮登録サービス
//
// @params
// firstName 名前
// firstNameKana 名前カナ
// familyName 名字
// familyNameKana 名字カナ
// password パスワード
// email メールアドレス
//
// @returns
// temporaryUserId 一時保存したユーザーを識別するID
func (rtu *RegisterTemporaryUser) RegisterTemporaryUser() {
	fmt.Println("サービス層：仮ユーザー登録API")
	// ユーザードメインサービス
	userService := service.NewUserService(rtu.Repo)
}
