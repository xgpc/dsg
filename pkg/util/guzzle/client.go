// Package guzzle
// @Author:        asus
// @Description:   $
// @File:          client
// @Data:          2022/1/2016:41
package guzzle

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	clone    int
	Host     string
	Error    error
	param    *Request
	response *Response
	ctx      context.Context
	request  *http.Request
}

func NewClient(options ...Options) *Client {
	client := &Client{
		ctx:   context.Background(),
		param: NewRequest(),
	}

	for _, option := range options {
		option.apply(client)
	}
	return client
}

// 获取到一个新的client进行请求
func (c *Client) getInstance() *Client {
	if c.clone == 0 {
		nc := *c
		nc.clone = 1

		return &nc
	}
	return c
}

// RequestHeader 设置请求头
func (c *Client) RequestHeader(key, value string) *Client {
	nc := c.getInstance()
	nc.param.Header(key, value)
	return nc
}

// RequestHeaders 获取请求Header头信息
func (c *Client) RequestHeaders() map[string]string {
	return getHeaders(c.param.header)
}

// RequestQuery 设置请求参数
func (c *Client) RequestQuery(query map[string][]string) *Client {
	nc := c.getInstance()
	nc.param.Query(query)
	return nc
}

// RequestSetBody 设置请求体
func (c *Client) RequestSetBody(body []byte) *Client {
	nc := c.getInstance()
	nc.param.Body(body)
	return nc
}

// RequestJSON 设置Json请求体
func (c *Client) RequestJSON(v interface{}) *Client {
	nc := c.getInstance()
	err := nc.param.JSON(v)
	if err != nil {
		nc.Error = err
	}
	return nc
}

// RequestForm 设置Form请求体
func (c *Client) RequestForm(form map[string][]string) *Client {
	nc := c.getInstance()
	nc.param.Form(form)
	return nc
}

func (c *Client) Get(path string) (*Response, error) {
	return c.Request(http.MethodGet, path)
}

func (c *Client) Post(path string) (*Response, error) {
	return c.Request(http.MethodPost, path)
}

func (c *Client) Put(path string) (*Response, error) {
	return c.Request(http.MethodPut, path)
}

func (c *Client) Delete(path string) (*Response, error) {
	return c.Request(http.MethodDelete, path)
}

// Request 发送请求
func (c *Client) Request(method, path string) (*Response, error) {
	nc := c.getInstance()

	// 请求前查看是否出现错误，出现错误就先抛出错误。等待用户解决完成后才重新执行
	if nc.Error != nil {
		return nil, nc.Error
	}

	// 设置路径,去除请求路径末尾的/,防止请求路径出现错误
	if c.Host != "" {
		path = c.Host + "/" + strings.TrimLeft(path, "/")
	}
	path = strings.TrimRight(path, "/")

	// body请求参数
	var buf io.Reader
	if len(nc.param.body) > 0 {
		buf = bytes.NewBuffer(nc.param.body)
	}

	// 获取request
	request, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	// 设置header头
	for k, val := range nc.RequestHeaders() {
		request.Header.Add(k, val)
	}

	// 设置query参数
	if len(nc.param.query) > 0 {
		request.URL.RawQuery = nc.param.query.Encode()
	}

	// 发送请求
	client := nc.client()
	response, err := client.Do(request.WithContext(nc.ctx))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	nc.request = request
	return NewResponse(response)
}

// 获取http请求客户端
func (c *Client) client() *http.Client {
	return &http.Client{}
}
