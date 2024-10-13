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
├─cmd/
│ ├─main.go------------- # 服务启动入口
│ ├─wire_gen.go--------- # Wire生成代码
│ └─wire.go------------- # Wire相关代码
├─config/
│ └─config.toml--------- # 配置文件
├─core/
│ ├─bootstrap/---------- # 启动程序所需要的初始化逻辑
│ │ ├─config.go--------- # 配置
│ │ ├─middleware.go----- # 中间件
│ │ ├─mysql.go---------- # mysql
│ │ ├─redis.go---------- # redis
│ │ ├─server.go--------- # gin服务端
│ │ ├─viper.go---------- # viper
│ │ ├─websocket.go------ # websocket
│ │ └─zap.go------------ # zap日志
│ └─app.go
├─internal/------------- # 内部包
│ ├─domain/------------- # 领域模型
│ │ └─user.go
│ ├─handler/
│ │ ├─middleware/
│ │ │ ├─ratelimit/------------ # 限流策略
│ │ │ │ ├─builder.go
│ │ │ │ └─slide_window.lua---- # 滑动窗口策略
│ │ │ └─login.go
│ │ ├─user.go
│ │ └─websocket.go
│ ├─repository/--------- # 数据访问层
│ │ ├─dao/
│ │ │ └─user.go
│ │ └─user.go
│ └─service/------------ # 业务逻辑层
│   └─user.go
├─pkg/------------------ # 外部包
│ └─ginx/
│   ├─jwt/
│   │ └─jwt.go
│   ├─middleware/
│   │ └─logger/
│   │   └─builder.go
│   └─websocket/
│     ├─client.go
│     └─manager.go
├─.gitignore
├─go.mod
├─go.sum
├─LICENSE
└─README.md
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
- [x] 集成 WebSocket
- [x] 集成 JWT
- [x] 登录 Token 校验
- [x] 限流策略
- [ ] 基础用户模块实现
- [ ] 短信验证服务模块开发
- [ ] 集成 Kafka
- [ ] 集成 swagger


