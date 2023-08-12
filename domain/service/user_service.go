package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/repository"
)

type UserService struct {
	repository domain.UserRepo
}

func NewUserService(rep domain.UserRepo) *UserService {
	return &UserService{repository: rep}
}

// emailでユーザ検索
//
// @params
// db dbインスタンス
//
// @returns
// isExist true 存在, false 存在しない
func (us *UserService) ExistByEmail(ctx context.Context, db *repository.Queryer, email string) (bool, error) {
	_, err := us.repository.FindUserByEmail(ctx, *db, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
