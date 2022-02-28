package router

import (
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/middleware/recover"
)

// Load 注册路由
func Load(app *iris.Application) {
	//跨域

	LoadRouter(app)
}

func LoadRouter(app *iris.Application) {
	api := app.Party("/api")
	Sys(api)
}
