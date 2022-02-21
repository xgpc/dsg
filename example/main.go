package main

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/middleware"
	"github.com/xgpc/dsg/router"
	"github.com/xgpc/dsg/service"
)

func main() {
	app := iris.New()

	// 加载配置
	frame.Load(app, "config.yml")

	// 中间件
	middleware.Load(app)

	//路由
	router.Load(app)

	// 加载服务
	service.LoadService()

	// 监听Http(s)
	frame.Listening(app)

}
