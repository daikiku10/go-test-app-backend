package handler

import (
	"context"

	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/service"
)

type RegisterTemporaryUserService interface {
	RegisterTemporaryUser(ctx context.Context, input service.ServiceRegisterTemporaryUserInput) (string, error)
}

type PostRegisterUserService interface {
	PostRegisterUser(ctx context.Context, input service.ServicePostRegisterUserInput) (*model.User, string, error)
}

// チャットアプリ
type RegisterChatUserService interface {
	RegisterChatUser(ctx context.Context, name string) (string, error)
}
