package handler

import (
	"net/http"

	"github.com/daikiku10/go-test-app-backend/auth"
	"github.com/gin-gonic/gin"
)

// 認証確認のミドルウェア
func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		if err := j.FillContext(ctx); err != nil {
			APIErrorResponse(ctx, http.StatusUnauthorized, "認証エラー", err.Error())
			return
		}
		ctx.Next()
	})
}
