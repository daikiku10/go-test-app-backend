package route

import (
	"github.com/daikiku10/go-test-app-backend/context"
	"github.com/daikiku10/go-test-app-backend/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init はルーティングを設定したEchoオブジェクトを作成する
func Init() *echo.Echo {
	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェア（共通処理）定義
	e.Use(CustomContext)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// グループ分け
	// ログイン状態でなければいけないAPIグループ
	api := e.Group("/api")

	// ログイン不要
	e.GET("/test", context.Convert(handler.GetHandler))
	e.POST("/test", context.Convert(handler.PostHandler))

	// ログイン必須
	api.GET("/test", context.Convert(handler.APIGetHandler))

	return e
}

func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ミドルウェアの前処理
		println("前処理です")

		// ミドルウェア内で次のハンドラまたはミドルウェアを呼び出す
		cc := context.NewContext(c)
		err := next(cc)

		// ミドルウェアの後処理
		println("後処理です")

		return err
	}
}
