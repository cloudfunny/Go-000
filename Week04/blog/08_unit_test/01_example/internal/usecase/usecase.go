package usecase

import "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/domain"

// PostUsecase PostUsecase
type PostUsecase struct {
	repo domain.IPostRepo
}

// PostUsecaseOption PostUsecaseOption
type PostUsecaseOption struct {
	Repo domain.IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(opt *PostUsecaseOption) (domain.IPostUsecase, error) {
	return &PostUsecase{repo: opt.Repo}, nil
}
