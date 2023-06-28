package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// e := route.Init()

	// // サーバ起動
	// // e.Start(":8080")
	// e.Logger.Fatal(e.Start(":8080"))

	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	fmt.Println("aaa")

	return nil
}
