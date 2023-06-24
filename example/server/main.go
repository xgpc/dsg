package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/xgpc/dsg"
)

func main() {
	dsg.Default()

	api := iris.Default()

	api.Use(func(ctx *context.Context) {
		//p := dsg.New(ctx)
		// TODO: 可以调用dsg框架中的接口
		//p.Ctx
	})

}
