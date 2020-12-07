package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohuishou/go-training/Week02/work/domain"
	"github.com/mohuishou/go-training/Week02/work/user/api"
	"github.com/mohuishou/go-training/Week02/work/user/repository"
	"github.com/mohuishou/go-training/Week02/work/user/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("db is must: %+v", err)
	}

	// 这里应该单独放一个地方
	db.AutoMigrate(&domain.User{})

	e := gin.Default()

	// 后面可以用 wire 进行依赖注入
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	api.NewUserHandler(e, usecase)

	if err := e.Run(":8080"); err != nil {
		log.Fatalf("gin exit: %+v", err)
	}
}
