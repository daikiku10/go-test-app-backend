package service

import (
	"context"
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/repository"
)

type ServicePostRegisterUserInput struct {
	TemporaryUserID string
	ConfirmCode     string
}
type PostRegisterUser struct {
	DB    repository.Execer
	Cache domain.Cache
	Repo  domain.UserRepo
}

func NewPostRegisterUser(db repository.Execer, cache domain.Cache, repo domain.UserRepo) *PostRegisterUser {
	return &PostRegisterUser{DB: db, Cache: cache, Repo: repo}
}

// ユーザー登録サービス
//
// @return ユーザー情報
func (pru *PostRegisterUser) PostRegisterUser(ctx context.Context, input ServicePostRegisterUserInput) (*model.User, string, error) {
	fmt.Println("サービス層：ユーザー登録API")
	// キャッシュからユーザー情報を復元する
	key := fmt.Sprintf("user:%s:%s", input.ConfirmCode, input.TemporaryUserID)
	u, err := pru.Cache.Get(ctx, key)
	if err != nil {
		return nil, "", fmt.Errorf("cannot get user in cache: %w", err)
	}

	// 復元後にキャッシュからユーザー情報を削除する
	if err = pru.Cache.Delete(ctx, key); err != nil {
		return nil, "", fmt.Errorf("cannot delete user in cache: %w", err)
	}

	// 復元したユーザー情報を解析
	temporaryUser := model.NewTemporaryUserString(u)
	firstName, firstNameKana, familyName, familyNameKana, email, password := temporaryUser.Split()

	// DBへ保存する
	user := &model.User{
		FirstName:      firstName,
		FirstNameKana:  firstNameKana,
		FamilyName:     familyName,
		FamilyNameKana: familyNameKana,
		Email:          email,
		Password:       password,
	}
	fmt.Println(user)
	if err := pru.Repo.RegisterUser(ctx, pru.DB, user); err != nil {
		return nil, "", fmt.Errorf("failed to register: %w", err)
	}

	// JWTを作成する

	return user, "", nil
}
