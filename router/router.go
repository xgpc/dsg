package router

import (
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/middleware/recover"
	"github.com/rs/cors"
	"github.com/xgpc/dsg/middleware"
)

// Load 注册路由
func Load(app *iris.Application) {
	//跨域
	c := cors.AllowAll()
	app.WrapRouter(c.ServeHTTP)
	app.Use(middleware.ExceptionLog)
	LoadRouter(app)
}

func LoadRouter(app *iris.Application) {
	api := app.Party("/api")
	Sys(api)
	MsgCode(api)
	User(api)
	Json(api)
	Wechat(api)
}
