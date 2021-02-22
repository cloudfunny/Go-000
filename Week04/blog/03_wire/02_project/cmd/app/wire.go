//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal"
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/service"
)

func NewPostService() (*service.PostService, error) {
	panic(wire.Build(internal.Set))
}
