package repository

import (
	"github.com/mohuishou/go-training/Week02/work/domain"
	"github.com/mohuishou/go-training/Week02/work/pkg/errcode"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewUserRepository NewUserRepository
func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &repository{db: db}
}

// Login 用户登录
func (r *repository) Login(username, password string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where(domain.User{Name: username, Password: password}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(errcode.UserLogin, "ErrRecordNotFound")
	}

	if err != nil {
		return nil, errors.Wrapf(errcode.DBQuery, "db errors is: %+v", err)
	}

	return &user, nil
}
