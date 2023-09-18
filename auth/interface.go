package auth

import (
	"context"
	"time"
)

type Store interface {
	Save(ctx context.Context, key, value string, minute time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Expired(ctx context.Context, key string, minute time.Duration) error
}
