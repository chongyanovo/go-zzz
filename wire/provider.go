package wire

import (
	"github.com/chongyanovo/go-zzz/core"
	"github.com/chongyanovo/go-zzz/core/bootstrap"
	"github.com/google/wire"
)

var BaseProvider = wire.NewSet(
	bootstrap.NewViper,
	bootstrap.NewConfig,
	bootstrap.NewMysql,
	bootstrap.NewRedis,
	bootstrap.NewZap,
	core.NewApplication,
)

var WebProvider = wire.NewSet(
	bootstrap.NewServer,
)
