// Package guzzle
// @Author:        asus
// @Description:   $
// @File:          options
// @Data:          2022/1/2016:41
package guzzle

import (
	"context"
	"strings"
	"time"
)

type Options interface {
	apply(*Client)
}

type OptionsFunc func(*Client)

func (f OptionsFunc) apply(client *Client) {
	f(client)
}

// 设置请求头
func WithHost(host string) OptionsFunc {
	return func(client *Client) {
		client.Host = strings.TrimRight(host, "/")
	}
}

func WithHead(key, value string) OptionsFunc {
	return func(client *Client) {
		client.param.Header(key, value)
	}
}

// 设置超时时间
func WithTimeOut(timeout int64) OptionsFunc {
	return func(client *Client) {
		client.ctx, _ = context.WithTimeout(client.ctx, time.Duration(timeout)*time.Second)
	}
}
