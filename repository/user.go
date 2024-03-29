package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/models"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ユーザー登録
//
// @param
// ユーザー情報
func (r *Repository) RegisterUser(ctx context.Context, db Execer, u *model.User) error {
	// 作成時間と更新時間を現在の時間にする。
	u.CreatedAt = r.Clocker.Now()
	u.UpdateAt = r.Clocker.Now()

	// sql作成
	sql := `INSERT INTO users(
		first_name, first_name_kana, family_name, family_name_kana, email, password, created_at, update_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// sql実行
	result, err := db.ExecContext(ctx, sql, u.FirstName, u.FirstNameKana, u.FamilyName, u.FamilyNameKana, u.Email, u.Password, u.CreatedAt, u.UpdateAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot create same email user: %w", ErrAlreadyEntry)
		}
		return err
	}
	// SQLがDBに新しい行を挿入した場合に、その行のIDを取得する。
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = model.UserID(id)

	return nil
}

// ユーザー登録(SQLBoiler)
func (r *Repository) RegisterUserBoiler(ctx context.Context, u *models.User, db *sqlx.DB) error {
	// 作成時間と更新時間を現在の時間にする。
	u.CreatedAt = r.Clocker.Now()
	u.UpdateAt = r.Clocker.Now()

	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot create same email user: %w", ErrAlreadyEntry)
		}
		return err
	}

	return nil
}

// ユーザー一覧取得(SQLBoiler)
func (r *Repository) GetAllUsers(ctx context.Context, db *sqlx.DB) ([]*models.User, error) {
	users, err := models.Users().All(ctx, db)
	if err != nil {
		return []*models.User{}, err
	}

	return users, nil
}

// ユーザー情報取得(SQLBoiler)
func (r *Repository) GetUserByID(ctx context.Context, db *sqlx.DB, tuID string) (*models.User, error) {
	user, err := models.Users(
		qm.Where("id=?", tuID),
	).One(ctx, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.User{}, errors.New("データが見つかりませんでした。")
		}
		return &models.User{}, err
	}

	return user, nil
}

// ユーザー情報更新(SQLBoiler)
func (r *Repository) UpdateUserByID(ctx context.Context, db *sqlx.DB, input model.InputUpdateUserByID) error {
	// 更新対象ユーザーの取得
	user, err := models.Users(
		qm.Where("id=?", input.UserID),
	).One(ctx, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("データが見つかりませんでした。")
		}
		return err
	}
	// 値の更新
	user.FirstName = input.FirstName
	user.UpdateAt = r.Clocker.Now()
	// DBへの更新
	rowsAft, err := user.Update(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println(rowsAft)
	return nil
}

// ユーザー削除(SQLBoiler)
func (r *Repository) DeleteUserByID(ctx context.Context, db *sqlx.DB, tuID string) error {
	// 更新対象ユーザーの取得
	user, err := models.Users(
		qm.Where("id=?", tuID),
	).One(ctx, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("データが見つかりませんでした。")
		}
		return err
	}
	// ユーザーの削除
	rowsAft, err := user.Delete(ctx, db)
	if err != nil {
		return err
	}
	fmt.Println(rowsAft)
	return nil
}

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
