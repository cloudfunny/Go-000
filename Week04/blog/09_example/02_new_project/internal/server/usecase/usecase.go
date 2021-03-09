package usecase

import (
	"context"

	"github.com/mohuishou/new-project/internal/domain"
)

type article struct {
	repo domain.IArticleRepo
}

// NewArticleUsecase init
func NewArticleUsecase(repo domain.IArticleRepo) domain.IArticleUsecase {
	return &article{repo: repo}
}

func (u *article) GetArticle(ctx context.Context, id int) (*domain.Article, error) {
	// 这里可能有其他业务逻辑...
	return u.repo.GetArticle(ctx, id)
}

func (u *article) CreateArticle(ctx context.Context, article *domain.Article) error {
	return u.repo.CreateArticle(ctx, article)
}
