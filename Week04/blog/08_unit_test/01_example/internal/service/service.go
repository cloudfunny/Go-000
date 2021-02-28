package service

import (
	"context"

	"github.com/jinzhu/copier"
	v1 "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/api/product/app/v1"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/domain"
)

var _ v1.BlogServiceHTTPServer = &PostService{}

// PostService PostService
type PostService struct {
	Usecase domain.IPostUsecase
}

// CreateArticle 创建文章
func (p *PostService) CreateArticle(ctx context.Context, req *v1.Article) (*v1.Article, error) {
	article, err := p.Usecase.CreateArticle(ctx, domain.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorId,
	})

	if err != nil {
		return nil, err
	}

	var resp v1.Article
	err = copier.Copy(&resp, &article)
	return &resp, err
}

// GetArticles 获取文章列表
func (p *PostService) GetArticles(ctx context.Context, req *v1.GetArticlesReq) (*v1.GetArticlesResp, error) {
	panic("implement me")
}
