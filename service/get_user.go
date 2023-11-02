package service

import (
	"fmt"
	"net/http"

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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NameKana    string `json:"nameKana"`
	MailAddress string `json:"mailAddress"`
}

func NewGetUser(db *sqlx.DB, repo domain.UserRepo) *GetUser {
	return &GetUser{DB: db, Repo: repo}
}

// ユーザー情報取得
func (c *GetUser) GetUser(ctx *gin.Context) {
	// クエリパラメータの取得
	uID := ctx.Query("userId")
	if uID == "" {
		ctx.JSON(400, "不正なパラメーターです。")
		return
	}

	u, err := c.Repo.GetUserByID(ctx, c.DB, uID)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	// 名前の結合, 半角スペースあり
	n := u.FamilyName + " " + u.FirstName
	nk := u.FamilyNameKana + " " + u.FirstNameKana

	// レスポンス作成
	res := GetUserResponse{
		ID:          int(u.ID),
		Name:        n,
		NameKana:    nk,
		MailAddress: u.Email,
	}
	ctx.JSON(http.StatusOK, res)
	fmt.Println("ユーザー情報取得成功です！")
}
