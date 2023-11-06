package service

import (
	"fmt"
	"net/http"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/domain/model"
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
func (u *UpdateUser) UpdateUser(ctx *gin.Context) {
	// リクエスト情報
	var input struct {
		UserID    int    `json:"userId"`
		FirstName string `json:"firstName"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, "不正なパラメーターです。")
		return
	}

	// トランザクション開始
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := u.Repo.UpdateUserByID(ctx, u.DB, model.InputUpdateUserByID(input)); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	// レスポンス作成
	res := struct {
		UserID int `json:"userId"`
	}{
		UserID: input.UserID,
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, res)
	fmt.Printf("%+v", "ユーザー情報更新成功です！")
}
