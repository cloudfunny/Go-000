package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/domain"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	mock_domain "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/mock/domain"

	v1 "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/api/product/app/v1"
)

type testPostService struct {
	post    *PostService
	usecase *mock_domain.MockIPostUsecase
	handler *gin.Engine
}

func initPostService(t *testing.T) *testPostService {
	ctrl := gomock.NewController(t)
	usecase := mock_domain.NewMockIPostUsecase(ctrl)
	service := &PostService{Usecase: usecase}

	handler := gin.New()
	v1.RegisterBlogServiceHTTPServer(handler, service)

	return &testPostService{
		post:    service,
		usecase: usecase,
		handler: handler,
	}
}

func TestPostService_CreateArticle(t *testing.T) {
	s := initPostService(t)
	s.usecase.EXPECT().
		CreateArticle(gomock.Any(), gomock.Eq(domain.Article{Title: "err", AuthorID: 1})).
		Return(domain.Article{}, fmt.Errorf("err"))
	s.usecase.EXPECT().
		CreateArticle(gomock.Any(), gomock.Eq(domain.Article{Title: "success", AuthorID: 2})).
		Return(domain.Article{Title: "success"}, nil)

	tests := []struct {
		name       string
		params     *v1.Article
		want       *v1.Article
		wantStatus int
		wantCode   int
		wantErr    string
	}{
		{
			name: "参数错误 author_id 必须",
			params: &v1.Article{
				Title:    "1",
				Content:  "2",
				AuthorId: 0,
			},
			want:       nil,
			wantStatus: 400,
			wantCode:   400,
		},
		{
			name: "失败",
			params: &v1.Article{
				Title:    "err",
				AuthorId: 1,
			},
			want:       nil,
			wantStatus: 500,
			wantCode:   -1,
		},
		{
			name: "成功",
			params: &v1.Article{
				Title:    "success",
				AuthorId: 2,
			},
			want: &v1.Article{
				Title: "success",
			},
			wantStatus: 200,
			wantCode:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 下面这些一般都会封装在一起，这里是为了演示

			// 初始化请求
			b, err := json.Marshal(tt.params)
			require.NoError(t, err)
			uri := fmt.Sprintf("/v1/author/%d/articles", tt.params.AuthorId)
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewReader(b))

			// 初始化响应
			w := httptest.NewRecorder()

			// 调用相应的handler接口
			s.handler.ServeHTTP(w, req)

			// 提取响应
			resp := w.Result()
			defer resp.Body.Close()
			require.Equal(t, tt.wantStatus, resp.StatusCode)

			// 读取响应body
			respBody, _ := ioutil.ReadAll(resp.Body)
			r := struct {
				Code int         `json:"code"`
				Msg  string      `json:"msg"`
				Data *v1.Article `json:"data"`
			}{}
			require.NoError(t, json.Unmarshal(respBody, &r))

			assert.Equal(t, tt.wantCode, r.Code)
			assert.Equal(t, tt.want, r.Data)
		})
	}
}
