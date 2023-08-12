package repository

import (
	"errors"
)

var (
	ErrNotExistEmail     = errors.New("メールアドレスが存在しません。")
	ErrAlreadyEntry      = errors.New("登録済みのメールアドレスは登録できません。")
	ErrNotFoundSession   = errors.New("確認コードまたは、セッションキーが無効です。")
	ErrNotMatchLogInfo   = errors.New("メールアドレスまたは、パスワードが異なります。")
	ErrNotUser           = errors.New("ユーザが存在しません。")
	ErrDifferentPassword = errors.New("パスワードが異なります。")
	ErrNotFound          = errors.New("データが存在しません。")
	ErrDBException       = errors.New("データベースで予期せぬエラーが起きました。")
	ErrCacheException    = errors.New("キャッシュで予期せぬエラーが起きました。")
)
