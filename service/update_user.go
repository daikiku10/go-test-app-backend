package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UpdateUser struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

func NewUpdateUser(db *sqlx.DB, repo domain.UserRepo) *UpdateUser {
	return &UpdateUser{DB: db, Repo: repo}
}

// ユーザー情報更新
func (c *UpdateUser) UpdateUser(ctx *gin.Context) {
	fmt.Println("ユーザー更新処理開始")
}
