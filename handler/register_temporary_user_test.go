package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daikiku10/go-test-app-backend/repository"
	"github.com/daikiku10/go-test-app-backend/service"
	"github.com/daikiku10/go-test-app-backend/utilstest"
	"github.com/gin-gonic/gin"
)

func TestRegisterTemporaryUser(t *testing.T) {
	type want struct {
		status  int
		resFile string
	}

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"正しいリクエストの時は201となる": {
			reqFile: "reqTest1",
			want: want{
				status:  http.StatusCreated,
				resFile: "resTest1",
			},
		},
		"正しくないパラメータの時は400となる": {
			reqFile: "reqTest2",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "resTest2",
			},
		},
		"バリデーションエラーの時は400となる": {
			reqFile: "reqTest3",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "resTest3",
			},
		},
		"既にユーザーが存在する時は409となる": {
			reqFile: "reqTest4",
			want: want{
				status:  http.StatusConflict,
				resFile: "resTest4",
			},
		},
		"予期せぬエラー時は500となる": {
			reqFile: "reqTest5",
			want: want{
				status:  http.StatusInternalServerError,
				resFile: "resTest5",
			},
		},
	}

	for k, test := range tests {

		t.Run(k, func(t *testing.T) {
			// サービス層のモック定義
			moq := &RegisterTemporaryUserServiceMock{}
			moq.RegisterTemporaryUserFunc = func(ctx context.Context, input service.ServiceRegisterTemporaryUserInput) (string, error) {
				// status によってレスポンスを変更する
				if test.want.status == http.StatusCreated {
					return "sessionID", nil
				}
				if test.want.status == http.StatusConflict {
					return "", repository.ErrAlreadyEntry
				}
				return "", errors.New("error from mock")
			}

			// テストデータの挿入する
			// gin Context の生成
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			// リクエストの生成
			req, _ := http.NewRequest("POST", "/temporary_user", bytes.NewReader(utilstest.LoadFile(t, test.reqFile)))
			// リクエスト情報をコンテキストに入れる
			c.Request = req

			// リクエストの送信

			// レスポンス

			// 検証
		})
		fmt.Println(k)
		fmt.Println(test)
	}
}
