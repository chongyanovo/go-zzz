<div align="center">
  <h3 align="center">go-zzz</h3>

  <p align="center">
    基于Gin开发,搭配 <a href="https://github.com/google/wire">wire</a> 依赖注入模式可写出更简洁和安全的代码。
    <br />
    <a href="https://github.com/chongyanovo/go-zzz"><strong>Explore »</strong></a>
    <br />
  </p>
</div>

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
- [ ] 基础用户模块实现
- [ ] 短信验证服务模块开发
- [ ] 集成 Kafka
- [ ] 集成 swagger