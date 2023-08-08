package service

import (
	"context"
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/repository"
)

type ServiceRegisterTemporaryUserInput struct {
	FirstName      string
	FamilyName     string
	FirstNameKana  string
	FamilyNameKana string
	Email          string
	Password       string
}

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
func (rtu *RegisterTemporaryUser) RegisterTemporaryUser(ctx context.Context, input ServiceRegisterTemporaryUserInput) (string, error) {
	fmt.Println("サービス層：仮ユーザー登録API")
	// ユーザードメインサービス
	// userService := service.NewUserService(rtu.Repo)
	// 登録可能なメールか確認する
	u, err := rtu.Repo.FindUserByEmail(ctx, rtu.DB, input.Email)
	fmt.Println(u)
	fmt.Println(err)

	// 成功時
	return "sessionIDTest", nil
}
