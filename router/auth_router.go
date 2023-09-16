package router

import (
	"context"

	"github.com/daikiku10/go-test-app-backend/auth"
	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/daikiku10/go-test-app-backend/handler"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/daikiku10/go-test-app-backend/utils/clock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 認証があるルーティングの設定を行う
//
// @param
// コンテキスト
// router ルーター
func SetAuthRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {
	// 時刻取得
	clocker := clock.RealClocker{}
	// レポジトリ作成
	// rep := repository.NewRepository(clocker)
	// 一時保存するキャッシュ(redis)の作成
	cache, err := repository.NewKVS(ctx, cfg)
	if err != nil {
		return err
	}
	// アクセストークンの生成
	jwt, err := auth.NewJWTer(cache, clocker)
	if err != nil {
		return err
	}

	// ルーティング設定
	groupRoute := router.Group("api/v1").Use(handler.AuthMiddleware(jwt))

	groupRoute.GET("/test/d")

	return nil
}
