package frame

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg"
	"log"
)

func Load(app *iris.Application, configPath string) {
	// 设置请求日志和常见HTTP错误码处理
	loadListeningSet(app)

	// load config
	dsg.LoadConf(configPath)

	// 加载MySQL、Redis
	loadMysql()
	loadRedis()

}

// Listening 开始监听端口
func Listening(app *iris.Application) {
	// 开始监听Http(s)
	log.Println("服务启动成功")
	host := ":" + dsg.Config.App.Port
	if dsg.Config.App.TLS != "" {
		tlsPrefix := dsg.Config.App.TLS
		_ = app.Run(iris.TLS(host, tlsPrefix+".crt", tlsPrefix+".key"))
	} else {
		_ = app.Run(iris.Addr(host))
	}
}
