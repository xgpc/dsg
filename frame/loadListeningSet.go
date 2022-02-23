package frame

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
)

func loadListeningSet(app *iris.Application) {
	// 请求日志（关闭后并发可突破1万，且不产生错误）
	loadSetCloseRequestLog()

	// 500响应处理
	loadSet500(app)

	// 404响应处理
	loadSet404(app)
}

func loadSetCloseRequestLog() {
	//if env.LogDriver() == log.LogDriverConsole {
	//	env.App().Use(logger.New())
	//}
	//	todo
}

func loadSet500(app *iris.Application) {
	app.OnErrorCode(500, onApp500)
}

func loadSet404(app *iris.Application) {
	app.OnErrorCode(404, onApp404)
}

func onApp500(ctx iris.Context) {
	panic(exce.CodeSysBusy)
}

func onApp404(ctx iris.Context) {
	ctx.WriteString(`
<h2>接口不存在！</h2>
`)
	ctx.StopExecution()
}
