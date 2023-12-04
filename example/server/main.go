package main

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2"
)

func main() {
	dsg.Load("")
	dsg.Default()

	api := iris.New()

	if dsg.Conf.TLS != "" {
		api.Run(iris.TLS(":8080", "server.crt", "server.key"))
	} else {
		api.Run(iris.Addr(":8080"))
	}
}
