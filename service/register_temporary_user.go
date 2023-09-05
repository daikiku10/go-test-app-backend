package service

import (
	"context"
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/domain/service"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/google/uuid"
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
	userService := service.NewUserService(rtu.Repo)

	// 登録可能なメールか確認する
	existMail, err := userService.ExistByEmail(ctx, &rtu.DB, input.Email)
	if err != nil {
		return "", err
	}
	if existMail {
		return "", fmt.Errorf("failed to register: %q", repository.ErrAlreadyEntry)
	}

	// パスワードのハッシュ化
	password, err := model.NewPassword(input.Password)
	if err != nil {
		return "", fmt.Errorf("cannot create password object: %w", err)
	}
	hashPassword, err := password.CreateHash()
	if err != nil {
		return "", fmt.Errorf("cannot create hash password: %w", err)
	}
	fmt.Println(hashPassword)

	// ユーザー情報をキャッシュに保存
	tempUserInfo := model.NewTemporaryUserString("")
	// キャッシュサーバーに保存するキーの作成
	uid := uuid.New().String()
	// 確認コードの作成
	confirmCode := model.NewConfirmCode().String()
	key := fmt.Sprintf("user:%s:%s", confirmCode, uid)
	// キャッシュサーバーに保存するvalueを作成
	userString := tempUserInfo.Join(input.FirstName, input.FirstNameKana, input.FamilyName, input.FamilyNameKana, input.Email, hashPassword)

	// 成功時
	return "sessionIDTest", nil
}
