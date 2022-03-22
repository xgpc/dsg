package main

import (
	"example/middleware"
	"example/router"
	"github.com/xgpc/dsg"
)

// @title dsg框架测试样例
// @version 2.0
// @description dsg框架2.0版本测试样例
// @termsOfService http://127.0.0.1:8080

// @host 127.0.0.1:8080
// @BasePath /api
func main() {
	server := dsg.New()

	//server.App
	// 中间件
	middleware.Load(server.App)

	// 路由
	router.LoadRouter(server.App)

	// 监听
	server.Start()
}
