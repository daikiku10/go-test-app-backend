package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type DeleteUser struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

func NewDeleteUser(db *sqlx.DB, repo domain.UserRepo) *DeleteUser {
	return &DeleteUser{DB: db, Repo: repo}
}

// ユーザー情報更新
func (c *DeleteUser) DeleteUser(ctx *gin.Context) {
	fmt.Printf("ユーザー削除成功です！")
}
