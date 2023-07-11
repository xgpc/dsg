package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rs/cors"
	"github.com/xgpc/dsg/v2"
	"github.com/xgpc/dsg/v2/middleware"
	"io"
	"net/http"
	"strings"
	"time"
)

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {

	dsg.Load("config.yml")

	dsg.Default(dsg.OptionEtcd(dsg.Conf.Etcd))
	err := dsg.Etcd.RegisterServiceDefault()
	if err != nil {
		panic(err)
	}

	// 启动iris

	api := iris.Default()

	c := cors.AllowAll()
	api.WrapRouter(c.ServeHTTP)
	api.Use(middleware.ExceptionLog)

	// 中间件
	api.WrapRouter(Handle)

	api.Handle("ANY", "/1", func(c *context.Context) {
		time.Sleep(time.Second * 10)
		c.JSON(`{"a":1}`)
	})

	api.Handle("ANY", "/2", func(c *context.Context) {
		c.JSON(`{"a":2}`)
	})

	err = api.Run(iris.Addr(":8082"))
	if err != nil {
		panic(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// 获取想访问的服务
	path := r.RequestURI
	serverName := GetPathServerName(path)

	// 找到服务器
	list := dsg.GetServiceList(serverName)

	// 转发
	if len(list) == 0 {
		next(w, r)
		return
	}

	// TODO: 添加负载均衡

	serverNode := list[0]

	if path[0] != '/' {
		path = "/" + path
	}

	// 服务转发
	res, err := Request(serverNode.GetUrl(dsg.GetEtcdLocalHost())+path, w, r)
	if err != nil {
		w.WriteHeader(res.StatusCode)
		io.WriteString(w, fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	// 数据返回
	// 从转发响应中读取响应体
	respBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		w.WriteHeader(res.StatusCode)
		io.WriteString(w, fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	// 将转发响应信息返回给客户端
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Write(respBody)
}

func Request(url string, w http.ResponseWriter, r *http.Request) (*http.Response, error) {
	// 从请求体中读取原始请求信息
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("Failed to read request body: %s  \n err:%v", string(body), err))
		return nil, err
	}

	// 设置要转发的目标URL
	targetURL := url

	// 创建一个新的HTTP请求
	req, err := http.NewRequest(r.Method, targetURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("Failed to create request: %v", err))
		return nil, err
	}

	// 将原始请求信息复制到新的请求中
	req.Header = r.Header.Clone()
	req.Body = io.NopCloser(r.Body)

	req.ContentLength = r.ContentLength

	// 发送转发请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("Failed to forward request: %v", err))
		return nil, err
	}

	return res, nil

}

func GetPathServerName(path string) string {

	split := strings.Split(path, "/")
	if len(split) > 2 {
		return split[1] + "/"
	}

	return ""

}
