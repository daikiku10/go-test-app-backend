package service

import (
	"fmt"
	"net/http"

	"github.com/daikiku10/go-test-app-backend/domain"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GetUsers struct {
	DB   *sqlx.DB
	Repo domain.UserRepo
}

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	FamilyName  string `json:"FamilyName"`
	MailAddress string `json:"mailAddress"`
}

type GetUsersResponse struct {
	UserList []User `json:"userList"`
}

func NewGetUsers(db *sqlx.DB, repo domain.UserRepo) *GetUsers {
	return &GetUsers{DB: db, Repo: repo}
}

// ユーザー一覧取得
func (g *GetUsers) GetUsers(ctx *gin.Context) {
	users, err := g.Repo.GetAllUsers(ctx, g.DB)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ul := []User{}
	for _, u := range users {
		ul = append(ul, User{
			ID:          int(u.ID),
			FirstName:   u.FirstName,
			FamilyName:  u.FamilyName,
			MailAddress: u.Email,
		})
	}
	res := GetUsersResponse{
		UserList: ul,
	}

	ctx.JSON(http.StatusOK, res)
	fmt.Println("ユーザー一覧取得成功です！")
}
