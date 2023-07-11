package main

import (
	"bytes"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
	"github.com/xgpc/dsg/v2"
	"github.com/xgpc/dsg/v2/exce"
	"github.com/xgpc/dsg/v2/middleware"
	"io"
	"net/http"
)

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {

	dsg.Load("config.yml")

	dsg.Default(dsg.OptionEtcd(dsg.Conf.Etcd))
	//api-gateway网关不进行注册, 防止重复跳转
	//err := dsg.Etcd.RegisterServiceDefault()
	//if err != nil {
	//   panic(err)
	//}

	// 启动iris

	api := iris.Default()

	c := cors.AllowAll()
	api.WrapRouter(c.ServeHTTP)
	api.Use(middleware.ExceptionLog)

	api.Any("/", func(context iris.Context) {
		p := dsg.NewBase(context)

		p.SuccessWithData(p.Ctx().Path())
	})

	api.Any("/{appName}/{p:path}", func(ctx iris.Context) {
		p := dsg.NewBase(ctx)

		// 获取想访问的服务
		path := ctx.Path()
		if path[0] != '/' {
			path = "/" + path
		}

		serverName := ctx.Params().Get("appName") + "/"

		// 找到服务器
		list := dsg.GetServiceList(serverName)

		// 转发
		if len(list) == 0 {
			exce.ThrowSys(exce.CodeRequestError, "未找到符合条件的服务")
		}

		// TODO: 添加负载均衡
		serverNode := list[0]

		// TODO: 不允许自己跳转自己,  但是又怕跳转别人, 先禁止跳转8082 同时api服务器不注册

		url := serverNode.GetUrl()

		// 服务转发
		res, err := Request(url+path, ctx)
		if err != nil {
			exce.ThrowSys(exce.CodeSysBusy, err.Error())
		}

		p.Ctx().WriteString(string(res))

	})

	api.Handle("ANY", "/member", func(c iris.Context) {
		p := dsg.NewBase(c)
		list, err2 := dsg.Etcd.Client.MemberList(context.Background())
		if err2 != nil {
			panic(err2)
		}

		p.SuccessWithData(list)
	})

	//api.Handle("ANY", "/list", func(c *context.Context) {
	//    p := dsg.NewBase(c)
	//    ////list, err2 := dsg.Etcd.Client.(context2.Background())
	//    //if err2 != nil {
	//    //    panic(err2)
	//    //}
	//    //
	//    p.SuccessWithData("暂未开启该功能")
	//})}

	dsg.Listening(dsg.Conf.Etcd.Port, dsg.Conf.TLS, api)

}

func Request(url string, ctx iris.Context) ([]byte, error) {
	// 从请求体中读取原始请求信息

	body, err := ctx.GetBody()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 设置要转发的目标URL
	targetURL := url

	// 创建一个新的HTTP请求
	req, err := http.NewRequest(ctx.Method(), targetURL, bytes.NewBuffer(body))
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, "http.NewRequest error: "+err.Error())
	}

	// 将原始请求信息复制到新的请求中
	req.Header = ctx.Request().Header.Clone()

	// 发送转发请求
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	defer rsp.Body.Close()

	rspHead := ctx.ResponseWriter()

	for k, v := range rsp.Header.Clone() {
		var value string
		if len(v) > 0 {
			value = v[0]
		}
		rspHead.Header().Set(k, value)
	}

	rspBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	if rsp.StatusCode != http.StatusOK {
		exce.ThrowSys(exce.CodeRequestError, rsp.Status)
	}

	return rspBody, nil
}
