package repo

import "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/domain"

// NewPostRepo NewPostRepo
func NewPostRepo() (domain.IPostRepo, error) {
	return new(domain.IPostRepo), nil
}
