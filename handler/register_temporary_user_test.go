package handler

import (
	"bytes"
	"context"
	"errors"
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
			reqFile: "testdata/register_temporary_user/201_req_json_golden",
			want: want{
				status:  http.StatusCreated,
				resFile: "testdata/register_temporary_user/201_res_json_golden",
			},
		},
		"正しくないパラメータの時は400となる": {
			reqFile: "testdata/register_temporary_user/400_1_req_json_golden",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "testdata/register_temporary_user/400_1_res_json_golden",
			},
		},
		"バリデーションエラーの時は400となる": {
			reqFile: "testdata/register_temporary_user/400_2_req_json_golden",
			want: want{
				status:  http.StatusBadRequest,
				resFile: "testdata/register_temporary_user/400_2_res_json_golden",
			},
		},
		"既にユーザーが存在する時は409となる": {
			reqFile: "testdata/register_temporary_user/409_req_json_golden",
			want: want{
				status:  http.StatusConflict,
				resFile: "testdata/register_temporary_user/409_res_json_golden",
			},
		},
		"予期せぬエラー時は500となる": {
			reqFile: "testdata/register_temporary_user/500_req_json_golden",
			want: want{
				status:  http.StatusInternalServerError,
				resFile: "testdata/register_temporary_user/500_res_json_golden",
			},
		},
	}

	for k, test := range tests {

		t.Run(k, func(t *testing.T) {
			// TODO: パラレル
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
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			// リクエストの生成
			req, _ := http.NewRequest("POST", "/temporary_user", bytes.NewReader(utilstest.LoadFile(t, test.reqFile)))
			// リクエスト情報をコンテキストに入れる
			c.Request = req

			// リクエストの送信
			rtu := RegisterTemporaryUser{
				Service: moq,
			}
			rtu.ServerHTTP(c)

			// レスポンス
			res := w.Result()

			// 検証
			utilstest.AssertResponse(t, res, test.want.status, utilstest.LoadFile(t, test.want.resFile))
		})
	}
}
