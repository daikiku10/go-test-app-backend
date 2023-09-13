package constant

const (
	// 確認コードの桁数
	ConfirmCodeLength int = 4
	// 確認コードの有効期限
	ConfirmCodeExpiration_m = 60
	// キャッシュ側のトークンの有効期限
	TokenExpiration_m int = 60
	// JWT側の最大有効期限（連続操作この指定時間経てば期限が切れる）
	MaxTokenExpiration_m int = 3600
	// ランダムパスワードの桁数
	RandomPasswordLength int = 12

	// 名字・名前の最大文字数
	UserNameMaxLength int = 50
	// パスワードの最大文字数
	PasswordMaxLength int = 50
	// メールアドレスの最大文字数
	MailAddressMaxLength int = 256
)
