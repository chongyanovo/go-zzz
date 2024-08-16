package wire

import (
	"github.com/chongyanovo/go-zzz/core"
	"github.com/chongyanovo/go-zzz/core/bootstrap"
	"github.com/chongyanovo/go-zzz/core/logger"
	"github.com/google/wire"
)

var BaseProvider = wire.NewSet(
	bootstrap.NewViper,
	bootstrap.NewConfig,
	bootstrap.NewMysql,
	bootstrap.NewRedis,
	bootstrap.NewZap,
	logger.NewZapLogger,
	core.NewApplication,
)

var WebProvider = wire.NewSet(
	bootstrap.NewServer,
)
