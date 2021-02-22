package internal

import (
	"github.com/google/wire"
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/repo"
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/service"
	"github.com/mohuishou/go-training/Week04/blog/03_wire/02_project/internal/usecase"
)

// Set Set
var Set = wire.NewSet(
	wire.Struct(new(service.PostService), "*"),
	wire.Struct(new(usecase.PostUsecaseOption), "*"),
	usecase.NewPostUsecase,
	repo.NewPostRepo,
)
