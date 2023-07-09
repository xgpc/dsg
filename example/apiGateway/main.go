package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/xgpc/dsg/v2"
	"github.com/xgpc/dsg/v2/pkg/etcd"
	"io"
	"net/http"
	"strings"
)

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	conf := dsg.Config{}
	dsg.Load()

	conf := etcd.Config{
		Name:                 "/apps/{server-name}",
		Address:              "127.0.0.1",
		Port:                 8081,
		Endpoints:            []string{"http://127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          10,
		DefLeaseSecond:       10,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	}
	client := etcd.New(conf)
	dsg.Default(dsg.OptionEtcd(conf))
	err := client.RegisterServiceDefault()
	if err != nil {
		panic(err)
	}

	// 启动iris

	api := iris.Default()

	// 中间件

	//
	api.Handle("all", "/", Handle)

	api.Run(iris.Addr(":8080"))
}

func Handle(ctx *context.Context) {

	// 获取想访问的服务
	path := ctx.Path()
	serverName := GetPathServerName(path)

	// 找到服务器
	list := dsg.GetServiceList(serverName)

	// 转发
	if len(list) == 0 {
		panic(path + "可用服务未上线")
	}

	// TODO: 添加负载均衡

	serverNode := list[0]

	if path[0] != '/' {
		path = "/" + path
	}

	// 服务转发
	res, err := Request(serverNode.GetUrl()+path, ctx)
	// 数据返回
	// 从转发响应中读取响应体
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	// 将转发响应信息返回给客户端
	ctx.StatusCode(res.StatusCode)
	ctx.ContentType(res.Header.Get("Content-Type"))
	ctx.Write(respBody)
}

func Request(url string, ctx *context.Context) (*http.Response, error) {
	// 从请求体中读取原始请求信息
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Failed to read request body: %s  \n err:%v", string(body), err))
		return nil, err
	}

	// 设置要转发的目标URL
	targetURL := url

	// 创建一个新的HTTP请求
	req, err := http.NewRequest(ctx.Method(), targetURL, nil)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Failed to create request: %v", err))
		return nil, err
	}

	// 将原始请求信息复制到新的请求中
	req.Header = ctx.Request().Header.Clone()
	req.Body = io.NopCloser(ctx.Request().Body)
	req.ContentLength = ctx.Request().ContentLength

	// 发送转发请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Failed to forward request: %v", err))
		return nil, err
	}
	defer res.Body.Close()

	return res, nil

}

func GetPathServerName(path string) string {

	split := strings.Split(path, "/")
	if len(split) < 3 {
		panic(path + "未能找到服务")
	}

	return "" + split[1] + "/" + split[2]

}
