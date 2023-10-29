package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GetUsers struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

func NewGetUsers(db *sqlx.DB, repo domain.UserRepo) *GetUsers {
	return &GetUsers{DB: db, Repo: repo}
}

// ユーザー一覧取得
func (c *GetUsers) GetUsers(ctx *gin.Context) {
	fmt.Println("ユーザー一覧取得API")
}
