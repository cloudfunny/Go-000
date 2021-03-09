package repo

import (
	"context"

	"github.com/mohuishou/new-project/internal/domain"
	"gorm.io/gorm"
)

type article struct {
	db *gorm.DB
}

// NewArticleRepo init
func NewArticleRepo(db *gorm.DB) domain.IArticleRepo {
	return &article{db: db}
}

func (r *article) GetArticle(ctx context.Context, id int) (*domain.Article, error) {
	var a domain.Article
	if err := r.db.WithContext(ctx).Find(&a, id); err != nil {
		// 这里返回业务错误码
	}
	return &a, nil
}

func (r *article) CreateArticle(ctx context.Context, article *domain.Article) error {

	if err := r.db.WithContext(ctx).Create(article); err != nil {
		// 这里返回业务错误码
	}
	return nil
}
