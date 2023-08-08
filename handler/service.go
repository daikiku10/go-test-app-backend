package handler

import (
	"context"

	"github.com/daikiku10/go-test-app-backend/service"
)

type RegisterTemporaryUserService interface {
	RegisterTemporaryUser(ctx context.Context, input service.ServiceRegisterTemporaryUserInput) (string, error)
}

type PostRegisterUserService interface {
	PostRegisterUser()
}
