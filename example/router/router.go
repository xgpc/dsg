package router

import (
	"example/api/category"
	_ "example/docs"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/middleware/recover"
)

// LoadRouter 注册路由
func LoadRouter(app *iris.Application) {

	swaggerUI := swagger.WrapHandler(swaggerFiles.Handler,
		swagger.URL("127.0.0.1:19610/swagger/doc.json"),
		swagger.DeepLinking(true),
	)
	app.Get("/swagger/{any:path}", swaggerUI)
	app.Get("/swagger", swaggerUI)

	api := app.Party("/api")
	Sys(api)
	api.Post("/category", category.CreateCategory)
	api.Get("/", category.GetList)
}
