package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/utils/clock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	UserID = "user_id"
	Email  = "email"
)

//go:embed certificate/secret.pem
var rawPrivKey []byte

//go:embed certificate/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey jwk.Key
	PublicKey  jwk.Key
	Store      Store
	Clocker    clock.Clocker
}

// JWTのインスタンス作成
func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {
	privKey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubKey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWaTer: public key: %w", err)
	}

	j := &JWTer{Store: s, PrivateKey: privKey, PublicKey: pubKey, Clocker: c}
	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

// アクセストークンの作成
//
// @return
// token アクセストークン
func (j *JWTer) GenerateToken(ctx context.Context, u model.User) ([]byte, error) {
	token, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/daikiku10/go-test-app-backend`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(time.Duration(constant.MaxTokenExpiration_m)*time.Minute)).
		Claim(Email, u.Email).
		Claim(UserID, u.ID).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}
	// redisにアクセストークンを保存する
	if err := j.Store.Save(ctx, fmt.Sprint(u.ID), token.JwtID(), time.Duration(constant.TokenExpiration_m)); err != nil {
		return nil, err
	}

	// JWTに署名をつける
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

// トークンを取得し、解析する
func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	// リクエストからアクセストークンを取得して解析する
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false),
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// アクセストークンを解析し、contextにUserIDとEmailをセットする
func (j *JWTer) FillContext(ctx *gin.Context) error {
	// トークンを解析する
	token, err := j.GetToken(ctx.Request.Context(), ctx.Request)
	if err != nil {
		return err
	}

	// 有効期限が切れていないか確認する
	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return fmt.Errorf("FillContext: failed to validate token: %w", err)
	}

	// トークンからUserIDを取得する
	id, ok := token.Get(UserID)
	if !ok {
		return fmt.Errorf("FillContext: not found %s", UserID)
	}
	uid := fmt.Sprintf("%v", id)
	// キャッシュ(redis)からトークンを取得する
	jwi, err := j.Store.Get(ctx, uid)
	if err != nil {
		// エラーの場合はキャッシュ側が有効期限切れ
		return fmt.Errorf("FillContext: %v expired %w", id, err)
	}

	// 他のログインを検査
	if jwi != token.JwtID() {
		return fmt.Errorf("FillContext: expired token %s because login another", jwi)
	}

	// 有効なアクセストークンを確認できたため、有効期限を延長する
	if err := j.Store.Expired(ctx, uid, time.Duration(constant.TokenExpiration_m)); err != nil {
		return fmt.Errorf("FillContext: can not be extended: %w", err)
	}

	return nil
}
