package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func NewRedis(config *Config) redis.Cmdable {
	r := config.RedisConfig
	return redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Username: r.Username,
		Password: r.Password,
	})
}
