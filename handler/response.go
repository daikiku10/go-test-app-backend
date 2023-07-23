package handler

import "github.com/gin-gonic/gin"

// クライアントへの返すレスポンスの型
type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// エラーレスポンスの型
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Message    string `json:"message"`
}

// APIレスポンスの作成（成功時）
//
// @params
// ctx ginのコンテキスト
// StatusCode ステータスコード
// Message メッセージ
// Data 返却するデータ
func APIResponse(ctx *gin.Context, StatusCode int, Message string, Data interface{}) {
	jsonResponse := &Response{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}
	ctx.JSON(StatusCode, jsonResponse)
}

// エラーレスポンス作成
//
// @params
// ctx ginのコンテキスト
// StatusCode ステータスコード
// Title エラータイトル
// Message メッセージ
func APIErrorResponse(ctx *gin.Context, StatusCode int, Title, Message string) {
	jsonResponse := ErrorResponse{
		StatusCode: StatusCode,
		Title:      Title,
		Message:    Message,
	}
	// リクエストの中断
	defer ctx.AbortWithStatus(StatusCode)
	ctx.JSON(StatusCode, jsonResponse)
}
