package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/go-redis/redis/v9"
)

type KVS struct {
	Cli *redis.Client
}

func NewKVS(ctx context.Context, cfg *config.Config) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		DB:   0,
	})
	// Redisサーバーに "PING" コマンドを送信し、サーバーの状態を確認する
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &KVS{Cli: cli}, nil
}

// 値の保存
func (kvs *KVS) Save(ctx context.Context, key, value string, minute time.Duration) error {
	return kvs.Cli.Set(ctx, key, value, minute*time.Minute).Err()
}

// 値の取得
func (kvs *KVS) Get(ctx context.Context, key string) (string, error) {
	result, err := kvs.Cli.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get by %q: %w", key, ErrNotFoundSession)
	}
	return result, err
}
