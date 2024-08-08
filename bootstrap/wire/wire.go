//go:build wireinject

package wire

import (
	"github.com/chongyanovo/go-zzz/bootstrap"
	"github.com/google/wire"
)

func InitApp() (bootstrap.Application, error) {
	wire.Build(
		BaseProvider,
		WebProvider,
	)
	return bootstrap.Application{}, nil
}
