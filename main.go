package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/daikiku10/go-test-app-backend/repository"
	routers "github.com/daikiku10/go-test-app-backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 環境変数の取得
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// gin のモード設定
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if cfg.Env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	// CORSミドルウェアの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// DB関係の初期化
	db, cleanup, err := repository.NewDB(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	// ルーティングの初期化
	if err := routers.SetRouting(ctx, db, router, cfg); err != nil {
		return err
	}

	// サーバー起動
	return router.Run(fmt.Sprintf(":%d", cfg.Port))

	// 課題：グレースフルシャットダウン
	// log.Printf("Listening and serving HTTP on:%v", cfg.Port)
	// server := NewServer(router, fmt.Sprintf(":%d", cfg.Port))
	// return server.Run(ctx)
}
