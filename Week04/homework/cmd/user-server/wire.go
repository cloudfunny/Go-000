//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mohuishou/go-training/Week04/homework/internal/service"
)

func InitUserService() (*service.UserService, error) {
	wire.Build(UserSet, NewDB, InitConfig)
	return new(service.UserService), nil
}
