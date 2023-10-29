package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/daikiku10/go-test-app-backend/auth"
	"github.com/daikiku10/go-test-app-backend/config"
	"github.com/daikiku10/go-test-app-backend/handler"
	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/daikiku10/go-test-app-backend/service"
	"github.com/daikiku10/go-test-app-backend/utils/clock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 認証がないルーティングの設定を行う
//
// @param
// ctx コンテキスト
// router ルーター
func SetRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {
	// 時刻取得
	clocker := clock.RealClocker{}
	// レポジトリ作成
	rep := repository.NewRepository(clocker)
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
	// ルーティングの設定
	groupRoute := router.Group("/api/v1")

	// 仮ユーザー登録
	registerTemporaryUserService := service.NewRegisterTemporaryUser(rep, db, cache)
	registerTemporaryUserHandler := handler.NewRegisterTemporaryUser(registerTemporaryUserService)
	groupRoute.POST("temporary_user", registerTemporaryUserHandler.ServerHTTP)
	// ユーザー登録
	postRegisterUserService := service.NewPostRegisterUser(db, cache, rep, jwt)
	postRegisterUserHandler := handler.NewPostRegisterUser(postRegisterUserService)
	groupRoute.POST("/user", postRegisterUserHandler.ServerHTTP)
	// ユーザー登録(SQLBoilerを使用して)
	createUserService := service.NewCreateUser(db, rep)
	groupRoute.GET("/user/d", createUserService.CreateUser)
	// ユーザー一覧取得
	getUsersService := service.NewGetUsers(db, rep)
	groupRoute.GET("/users/d", getUsersService.GetUsers)

	router.GET("/test1", func(ctx *gin.Context) {
		fmt.Printf("aaa")
		rsp := struct {
			ID string `json:"temporaryUserId"`
		}{ID: "テストIDだよ！！！！！"}

		handler.APIResponse(ctx, http.StatusCreated, "テストAPI成功しました。", rsp)

	})
	return nil
}

// echoの練習------------------------------------------

// // Init はルーティングを設定したEchoオブジェクトを作成する
// func Init() *echo.Echo {
// 	// Echoインスタンスの作成
// 	e := echo.New()

// 	// ミドルウェア（共通処理）定義
// 	e.Use(CustomContext)
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// グループ分け
// 	// ログイン状態でなければいけないAPIグループ
// 	api := e.Group("/api")

// 	// ログイン不要
// 	e.GET("/test", context.Convert(handler.GetHandler))
// 	e.POST("/test", context.Convert(handler.PostHandler))

// 	// ログイン必須
// 	api.GET("/test", context.Convert(handler.APIGetHandler))

// 	return e
// }

// // CustomContext はハンドラに渡すContext型をカスタムコンテキストに変換するミドルウェア
// func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// ミドルウェアの前処理
// 		println("前処理です")

// 		// ミドルウェア内で次のハンドラまたはミドルウェアを呼び出す
// 		// 型変換する
// 		cc := context.NewContext(c)
// 		err := next(cc)

// 		// ミドルウェアの後処理
// 		println("後処理です")

// 		return err
// 	}
// }
