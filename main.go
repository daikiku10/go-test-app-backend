package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Go Echo")
}

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
