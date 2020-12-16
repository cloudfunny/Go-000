package usecase

import (
	"context"

	"github.com/mohuishou/go-training/Week04/homework/internal/domain"
)

type user struct {
	repo domain.IUserRepo
}

// NewUserUsecase NewUserUsecase
func NewUserUsecase(repo domain.IUserRepo) domain.IUserUsecase {
	return &user{repo: repo}
}

// GetUserInfo GetUserInfo
func (u *user) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetUserInfo(ctx, id)
}
