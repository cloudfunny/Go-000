//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/service"
)

func NewPostService() (*service.PostService, error) {
	panic(wire.Build(internal.Set))
}
