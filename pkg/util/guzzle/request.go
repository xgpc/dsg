// Package guzzle
// @Author:        asus
// @Description:   $
// @File:          param
// @Data:          2022/1/2016:51
package guzzle

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	header http.Header
	query  url.Values
	body   []byte
}

func NewRequest() *Request {
	return &Request{
		header: http.Header{},
		query:  url.Values{},
	}
}

// String 返回请求的string内容
func (r *Request) String() string {
	return string(r.body)
}

func (r *Request) Header(key, value string) {
	r.header.Set(key, value)
}

func (r *Request) Headers() map[string]string {
	return getHeaders(r.header)
}

// Query 设置请求参数
func (r *Request) Query(query map[string][]string) {
	for key, val := range query {
		for _, item := range val {
			r.query.Add(key, item)
		}
	}
}

func (r *Request) Body(body []byte) {
	r.body = body
}

func (r *Request) JSON(v interface{}) error {
	r.Header("Content-Type", "application/json")

	body, err := json.Marshal(v)
	if err != nil {
		return err
	}

	r.Body(body)
	return nil
}

func (r *Request) Form(form map[string][]string) {
	param := url.Values{}
	for key, val := range form {
		for _, item := range val {
			param.Add(key, item)
		}
	}

	r.Header("Content-Type", "application/x-www-form-urlencoded")
	r.Body([]byte(param.Encode()))
}

func getHeaders(header http.Header) map[string]string {
	h := make(map[string]string)

	for key := range header {
		h[key] = header.Get(key)
	}

	return h
}
