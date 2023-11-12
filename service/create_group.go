package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/daikiku10/go-test-app-backend/models"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jmoiron/sqlx"
)

type CreateGroup struct {
	DB   *sqlx.DB
	Repo domain.GroupRepo
}

func NewCreateGroup(db *sqlx.DB, repo domain.GroupRepo) *CreateGroup {
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

	// トランザクション開始
	tx, err := cg.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	g := &models.Group{
		Name: input.Name,
	}

	if err := cg.Repo.RegisterGroup(ctx, cg.DB, g); err != nil {
		tx.Rollback()
		ctx.JSON(400, err.Error())
		return
	}
	tx.Commit()
	fmt.Println("グループ登録成功!")
}
