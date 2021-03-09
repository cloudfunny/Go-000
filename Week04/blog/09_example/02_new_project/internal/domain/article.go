package domain

import "context"

// Article 文章
type Article struct {
	Model // 基础结构体，包含 id, created_at, deleted_at, updated_at

	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags" gorm:"many2many:article_tags"`
}

// IArticleUsecase IArticleUsecase
type IArticleUsecase interface {
	GetArticle(ctx context.Context, id int) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
}

// IArticleRepo IArticleRepo
type IArticleRepo interface {
	GetArticle(ctx context.Context, id int) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
}
