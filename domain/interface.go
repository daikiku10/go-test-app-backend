package domain

import (
	"context"
	"time"

	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/repository"
)

// Userに対するインターフェース
type UserRepo interface {
	FindUserByEmail(ctx context.Context, db repository.Queryer, email string) (model.User, error)
}

// Cache(redis)に対するインターフェース
type Cache interface {
	Save(ctx context.Context, key, value string, minute time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
