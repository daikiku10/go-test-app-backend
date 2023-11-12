package repository

import (
	"context"

	"github.com/daikiku10/go-test-app-backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// グループ登録
func (r *Repository) RegisterGroup(ctx context.Context, db *sqlx.DB, g *models.Group) error {
	// 作成時間と更新時間を現在の時間にする。
	g.CreatedAt = r.Clocker.Now()
	g.UpdateAt = r.Clocker.Now()

	if err := g.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	return nil
}
