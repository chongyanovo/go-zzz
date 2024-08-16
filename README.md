<div align="center">
  <h3 align="center">go-zzz</h3>

  <p align="center">
    基于Gin开发,搭配 <a href="https://github.com/google/wire">wire</a> 依赖注入模式可写出更简洁和安全的代码。
    <br />
    <a href="https://github.com/chongyanovo/go-zzz"><strong>Explore »</strong></a>
    <br />
  </p>
</div>

```text
.
├── bootstrap                     启动程序所需要的初始化逻辑
│   ├── app.go
│   ├── internal                  内部包
│   │   ├── config.go
│   │   ├── file_rotatelogs.go
│   │   ├── mysql.go
│   │   ├── redis.go
│   │   ├── server.go
│   │   ├── viper.go
│   │   └── zap.go
│   └── wire                      Wire库相关代码
│       ├── provider.go
│       ├── wire.go
│       └── wire_gen.go
├── cmd                           服务启动入口
│   └── main.go
├── config                        配置文件
│   └── config.toml
├── go.mod
├── go.sum
└─── internal                     项目内部包
    ├── domain                    领域模型
    ├── repository                数据访问层
    ├── service                   业务逻辑层
    └── web                       Http相关代码
```


## Dependencies
- [Golang1.22](https://github.com/golang/go)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/jinzhu/gorm)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [Redis](https://github.com/go-redis/redis)

## Getting Started
使用 go-zzz 快速构建一个新项目:
```bash
# 下载 go-zzz
git clone https://github.com/chongyanovo/go-zzz

# 进入项目目录
cd go-zzz

# 安装 go mod 依赖
go mod tidy
```

## Todo
- [x] 集成 gorm
- [x] 日志接口并集成 gorm
- [ ] 集成 JWT
- [ ] 基础用户模块实现
- [ ] 短信验证服务模块开发
- [ ] 集成 Kafka
- [ ] 集成 swagger


