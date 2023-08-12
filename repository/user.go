package repository

import (
	"context"

	"github.com/daikiku10/go-test-app-backend/domain/model"
)

// メールと一致するユーザーを取得する
// @params
// ctx context
// db dbインスタンス
// email email
//
// @returns
// model.User ユーザ情報
func (r *Repository) FindUserByEmail(ctx context.Context, db Queryer, email string) (model.User, error) {
	sql := `
		SELECT * from users
		WHERE users.email = ?`

	var user model.User

	if err := db.GetContext(ctx, &user, sql, email); err != nil {
		// 見つけられない時(その他のエラーも含む)
		// 見つけられない時のエラーは利用側で
		// errors.Is(err, sql.ErrNoRows) で判断する
		return user, err
	}
	return user, nil
}
