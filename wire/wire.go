//go:build wireinject

package wire

import (
	"github.com/chongyanovo/go-zzz/core"
	"github.com/google/wire"
)

func InitApp() (core.Application, error) {
	wire.Build(
		BaseProvider,
		WebProvider,
	)
	return core.Application{}, nil
}
