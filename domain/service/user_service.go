package service

import (
	"github.com/daikiku10/go-test-app-backend/domain"
)

type UserService struct {
	repository domain.UserRepo
}

func NewUserService(rep domain.UserRepo) *UserService {
	return &UserService{repository: rep}
}
