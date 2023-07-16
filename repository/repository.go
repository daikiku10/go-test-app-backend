package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/jmoiron/sqlx"
)

// DBへの接続
// @params
// ctx コンテキスト
// cfg 接続設定の環境変数
//
// @return
// DBインスタンス
// DBのクローズ関数(呼び先でdeferで呼ぶ必要あり)
func NewDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	// sqlx.Connectを使用すると内部でpingする
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}
	// Openは実際に接続テストが行われない。
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil

}