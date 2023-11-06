package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jmoiron/sqlx"
)

type CreateGroup struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

func NewCreateGroup(db *sqlx.DB, repo domain.UserRepo) *CreateGroup {
	return &CreateGroup{DB: db, Repo: repo}
}

// グループ登録
func (cg *CreateGroup) CreateGroup(ctx *gin.Context) {
	// リクエスト情報
	var input struct {
		Name string `json:"groupName"`
	}
	// クライアントからのパラメータの確認
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	fmt.Println(input.Name)
	// バリデーションチェック
	err := validation.ValidateStruct(&input,
		validation.Field(
			&input.Name,
			validation.Required,
			validation.RuneLength(1, constant.GroupNameMaxLength),
		),
	)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	fmt.Println("バリデーションOK")
}
