// Package dsg
// @Author:        asus
// @Description:   $
// @File:          New
// @Data:          2022/2/2118:09
//
package dsg

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	//"github.com/xgpc/dsg/middleware"
	"github.com/xgpc/dsg/router"
	service2 "github.com/xgpc/dsg/service"
)

type service struct {
	App *iris.Application
}

func New() *service {
	app := iris.New()

	// 加载配置
	frame.Load(app, "config.yaml")

	// 中间件
	//middleware.Load(app)

	//路由
	router.Load(app)

	// 加载服务
	service2.LoadService()

	return &service{
		App: app,
	}
}

// Start 启动监听
func (app *service) Start() {
	frame.Listening(app.App)
}
