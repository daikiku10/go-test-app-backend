package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type res struct {
	TestInt    int    `json:"testInt"`
	TestString string `json:"testString"`
}

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェア
	e.Use(middleware)

	// ルート指定
	e.GET("/test", getHandler)
	e.POST("/test", postHandler)

	// サーバ起動
	e.Start(":8080")
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ミドルウェアの前処理
		println("前処理です")

		// ミドルウェア内で次のハンドラまたはミドルウェアを呼び出す
		err := next(c)

		// ミドルウェアの後処理
		println("後処理です")

		return err
	}
}

func getHandler(c echo.Context) error {
	response := res{
		TestInt:    21,
		TestString: "getTestです。",
	}
	return c.JSON(http.StatusOK, response)
}

func postHandler(c echo.Context) error {
	response := res{
		TestInt:    21,
		TestString: "postTestです。",
	}
	return c.JSON(http.StatusCreated, response)
}
