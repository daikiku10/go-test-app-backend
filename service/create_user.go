package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type CreateUser struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

func NewCreateUser(db *sqlx.DB, repo domain.UserRepo) *CreateUser {
	return &CreateUser{DB: db, Repo: repo}
}

// ユーザー登録
func (c *CreateUser) CreateUser(ctx *gin.Context) {
	// トランザクション開始
	tx, err := c.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	u := &models.User{
		FamilyName:     "test",
		FamilyNameKana: "テスト",
		FirstName:      "taro",
		FirstNameKana:  "太郎",
		Email:          "testMail",
		Password:       "pass",
	}
	if err := c.Repo.RegisterUserBoiler(ctx, u, c.DB); err != nil {
		tx.Rollback()
		ctx.JSON(400, err.Error())
		return
	}
	tx.Commit()
	fmt.Println("登録成功です！")
}
