package main

import (
	"context"

	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mohuishou/go-training/Week04/homework/internal/repository"
	"github.com/mohuishou/go-training/Week04/homework/internal/repository/ent"
	"github.com/mohuishou/go-training/Week04/homework/internal/service"
	"github.com/mohuishou/go-training/Week04/homework/internal/usecase"
	"github.com/spf13/viper"
)

// 注意这个文件下一般情况下不会放在 main 当中，我们一般会放在一个统一的初始化包
// 或者是分布在各个包的初始化文件当中，这里由于时间原因暂时先放在这里了

// UserSet 这里因为只有这一个二进制就暂时先在这里面定义了
// 一般情况下 internal 如果分为多个 app 会在每个 app 下定义一个
var UserSet = wire.NewSet(
	service.NewUserService,
	repository.NewRepository,
	usecase.NewUserUsecase,
)

// NewDB 初始化 db 连接
func NewDB(v *viper.Viper) (*ent.Client, error) {
	client, err := ent.Open(
		v.Sub("db").GetString("type"),
		v.Sub("db").GetString("dsn"),
	)
	if err != nil {
		return nil, err
	}

	// 数据迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

// InitConfig 初始化配置文件
func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
