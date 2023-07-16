package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/daikiku10/go-test-app-backend/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := route.Init()

	// サーバ起動
	// e.Start(":8080")
	e.Logger.Fatal(e.Start(":8080"))

	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	fmt.Println("aaa")
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
	_, cleanup, err := repository.NewDB(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	return nil
}
