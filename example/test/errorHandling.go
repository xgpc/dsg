package main

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2/exce"
	"github.com/xgpc/dsg/v2/middleware"
)

func errorHanding() {

	// 报错
	// {
	//	"code" : exce.CodeSysBusy
	//	"msg" : "错误信息"
	// }
	exce.ThrowSys(exce.CodeSysBusy, "错误信息")

	// 中间件 使用
	// ExceptionLog 会捕获异常并返回
	api := iris.Default()
	api.Use(middleware.ExceptionLog)

}
