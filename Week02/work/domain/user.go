package domain

import "gorm.io/gorm"

// User user
type User struct {
	gorm.Model

	Name     string `json:"name"`
	Password string `json:"password"`
}

// IUserRepository IUserRepository
type IUserRepository interface {
	Login(username, password string) (*User, error)
}

// IUserUsecase IUserUsecase
type IUserUsecase interface {
	Login(username, password string) (*User, error)
}
