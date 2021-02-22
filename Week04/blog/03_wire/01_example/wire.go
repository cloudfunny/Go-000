//+build wireinject

package example

import (
	"github.com/google/wire"
)

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		PostServiceSet,
	))
}
