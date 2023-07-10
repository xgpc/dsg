package dsg

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"github.com/xgpc/dsg/v2/pkg/util"
	"log"
)

func LoadYml(out interface{}, configPath string) {
	util.LoadYmlConf(out, configPath)
}

func Load(configPath string) {
	// load config
	util.LoadYmlConf(Conf, configPath)
}

// Listening 开始监听端口
func Listening(Port int, Tls string, app *iris.Application) {
	// 开始监听Http(s)
	log.Println("服务启动成功")
	host := ":" + strconv.Itoa(Port)
	if Tls != "" {
		tlsPrefix := Tls
		_ = app.Run(iris.TLS(host, tlsPrefix+".pem", tlsPrefix+".key"))
	} else {
		_ = app.Run(iris.Addr(host))
	}
}
