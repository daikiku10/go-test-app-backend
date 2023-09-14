package auth

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/daikiku10/go-test-app-backend/constant"
	"github.com/daikiku10/go-test-app-backend/domain/model"
	"github.com/daikiku10/go-test-app-backend/utils/clock"
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
