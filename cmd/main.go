package main

import (
	"fmt"
)

func main() {
	app, err := InitApp()
	if err != nil {
		fmt.Println("初始化失败")
		panic(err)
	}
	app.Server.Run(fmt.Sprintf("%s:%d", app.Config.ServerConfig.Host, app.Config.ServerConfig.Port))
}
