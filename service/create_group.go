package service

import (
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
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
	fmt.Println("グループ登録開始")
}
