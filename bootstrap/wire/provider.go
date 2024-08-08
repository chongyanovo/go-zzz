package wire

import (
	"github.com/chongyanovo/go-zzz/bootstrap"
	"github.com/chongyanovo/go-zzz/bootstrap/internal"
	"github.com/google/wire"
)

var BaseProvider = wire.NewSet(
	internal.NewViper,
	internal.NewConfig,
	internal.NewMysql,
	internal.NewRedis,
	internal.NewZap,
	bootstrap.NewApplication,
)

var WebProvider = wire.NewSet(
	internal.NewServer,
)
