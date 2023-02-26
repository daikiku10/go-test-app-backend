package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db := sqlConnect()
	defer db.Close()
	
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
  USER := "go_test"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	// dbコンテナが起動してからもMySQLが立ち上がるまでに時間がかかるため、
	// もしDBにつながらなかった場合に1秒待ってからリトライする。
	// 3分つながらなかった場合はエラーする。
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Println(".")
			time.Sleep(time.Second)
			count ++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}