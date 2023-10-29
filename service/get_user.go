package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GetUser struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

type UserByID struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	FamilyName  string `json:"FamilyName"`
	MailAddress string `json:"mailAddress"`
}

type GetUserResponse struct {
	User UserByID `json:"user"`
}

func NewGetUser(db *sqlx.DB, repo domain.UserRepo) *GetUser {
	return &GetUser{DB: db, Repo: repo}
}

// ユーザー情報取得
func (c *GetUser) GetUser(ctx *gin.Context) {
	fmt.Println("ユーザー情報取得成功です！")
}
