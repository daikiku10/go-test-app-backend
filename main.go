package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type res struct {
	TestInt    int    `json:"testInt"`
	TestString string `json:"testString"`
}

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェア
	e.Use(Middleware)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルート指定
	e.GET("/test", getHandler)
	e.POST("/test", postHandler)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
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

	name := c.QueryParam("name")
	println(name)

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
