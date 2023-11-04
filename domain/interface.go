package domain

import (
	"context"
	"time"

	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/models"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/jmoiron/sqlx"
)

// Userに対するインターフェース
type UserRepo interface {
	// SQLBoiler
	RegisterUserBoiler(ctx context.Context, u *models.User, db *sqlx.DB) error
	GetAllUsers(ctx context.Context, db *sqlx.DB) ([]*models.User, error)
	GetUserByID(ctx context.Context, db *sqlx.DB, tuID string) (*models.User, error)
	UpdateUserByID(ctx context.Context, db *sqlx.DB, input model.InputUpdateUserByID) error
	// SQLBoilerではない
	RegisterUser(ctx context.Context, db repository.Execer, u *model.User) error
	FindUserByEmail(ctx context.Context, db repository.Queryer, email string) (model.User, error)
}

// Cache(redis)に対するインターフェース
type Cache interface {
	Save(ctx context.Context, key, value string, minute time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

// JWTトークンに対するインターフェース
type TokenGenerator interface {
	GenerateToken(ctx context.Context, u model.User) ([]byte, error)
}
