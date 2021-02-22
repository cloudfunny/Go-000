package repo

import "github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/domain"

// NewPostRepo NewPostRepo
func NewPostRepo() (domain.IPostRepo, error) {
	return new(domain.IPostRepo), nil
}
