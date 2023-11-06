package service

import (
	"fmt"
	"net/http"

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

// ユーザー削除
func (d *DeleteUser) DeleteUser(ctx *gin.Context) {
	// クエリパラメータの取得
	uID := ctx.Query("userId")
	if uID == "" {
		ctx.JSON(400, "不正なパラメーターです。")
		return
	}

	err := d.Repo.DeleteUserByID(ctx, d.DB, uID)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	// トランザクション開始
	tx, err := d.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// レスポンス作成
	res := struct {
		ID string `json:"userID"`
	}{
		ID: uID,
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, res)
	fmt.Printf("ユーザー削除成功です！")
}
