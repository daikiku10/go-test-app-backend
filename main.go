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

	// ルート指定
	e.GET("/test", getHandler)
	e.POST("/test", postHandler)

	// サーバ起動
	e.Start(":8080")
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
