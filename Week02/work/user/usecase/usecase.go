package usecase

import "github.com/mohuishou/go-training/Week02/work/domain"

type usecase struct {
	repo domain.IUserRepository
}

// NewUserUsecase NewUserUsecase
func NewUserUsecase(repo domain.IUserRepository) domain.IUserUsecase {
	return &usecase{repo: repo}
}

// Login 用户登录
func (u *usecase) Login(username, password string) (*domain.User, error) {
	return u.repo.Login(username, password)
}
