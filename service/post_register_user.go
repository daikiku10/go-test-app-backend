package service

import (
	"context"
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/domain/model"
)

type ServicePostRegisterUserInput struct {
	TemporaryUserID string
	ConfirmCode     string
}
type PostRegisterUser struct {
	Cache domain.Cache
}

func NewPostRegisterUser(cache domain.Cache) *PostRegisterUser {
	return &PostRegisterUser{Cache: cache}
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
	fmt.Println(u)

	// 復元後にキャッシュからユーザー情報を削除する
	if err = pru.Cache.Delete(ctx, key); err != nil {
		return nil, "", fmt.Errorf("cannot delete user in cache: %w", err)
	}

	// 復元したユーザーぞ情報を解析

	// DBへ保存する

	// JWTを作成する

	return nil, "", nil
}
