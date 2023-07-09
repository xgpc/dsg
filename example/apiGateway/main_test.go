package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"testing"
)

func Demo() {
	app := iris.New()

	app.HandleFunc("ALL", "/", func(ctx iris.Context) {
		// 从请求体中读取原始请求信息
		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.WriteString(fmt.Sprintf("Failed to read request body: %s  \n err:%v", string(body), err))
			return
		}

		// 设置要转发的目标URL
		targetURL := "http://localhost:8081"

		// 创建一个新的HTTP请求
		req, err := http.NewRequest(ctx.Method(), targetURL, nil)
		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.WriteString(fmt.Sprintf("Failed to create request: %v", err))
			return
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
			return
		}
		defer res.Body.Close()

		// 从转发响应中读取响应体
		respBody, err := io.ReadAll(res.Body)
		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.WriteString(fmt.Sprintf("Failed to read response body: %v", err))
			return
		}
		b := string(respBody)

		// 将转发响应信息返回给客户端
		ctx.StatusCode(res.StatusCode)
		ctx.ContentType(res.Header.Get("Content-Type"))
		ctx.Write([]byte(b))
	})

	app.Run(iris.Addr(":8080"))
}

func TestName(t *testing.T) {
	a := "/apps/server-name/a"
	b := "/api/apps/server-name/a"

	fmt.Println(GetPathServerName(a))
	fmt.Println(GetPathServerName(b))
}
