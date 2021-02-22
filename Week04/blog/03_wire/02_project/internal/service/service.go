package service

import (
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/domain"
)

// PostService PostService
type PostService struct {
	Usecase domain.IPostUsecase
}
