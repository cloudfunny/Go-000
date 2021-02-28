package domain

import "context"

// Article 文章
type Article struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorID int32 `json:"author_id,omitempty" form:"author_id" uri:"author_id"`
}

// IPostRepo IPostRepo
type IPostRepo interface{}

// IPostUsecase IPostUsecase
type IPostUsecase interface {
	CreateArticle(ctx context.Context, article Article) (Article, error)
}
