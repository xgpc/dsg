package dsg

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/env"
	"github.com/xgpc/dsg/pkg/util"
	"log"
)

func Load(conf *env.Conf, configPath string) {
	// load config
	util.LoadYmlConf(conf, configPath)
}

// Listening 开始监听端口
func Listening(app *iris.Application) {
	// 开始监听Http(s)
	log.Println("服务启动成功")
	host := ":" + env.Config.App.Port
	if env.Config.App.TLS != "" {
		tlsPrefix := env.Config.App.TLS
		_ = app.Run(iris.TLS(host, tlsPrefix+".pem", tlsPrefix+".key"))
	} else {
		_ = app.Run(iris.Addr(host))
	}
}
