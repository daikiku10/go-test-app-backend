package main

import (
	"github.com/daikiku10/go-test-app-backend/route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := route.Init()

	// サーバ起動
	// e.Start(":8080")
	e.Logger.Fatal(e.Start(":8080"))
}
