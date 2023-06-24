// Package guzzle
// @Author:        asus
// @Description:   $
// @File:          response
// @Data:          2022/1/2017:59
package guzzle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	status  int
	headers http.Header
	cookies []*http.Cookie
	body    []byte
}

func NewResponse(response *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		status:  response.StatusCode,
		headers: response.Header,
		cookies: response.Cookies(),
		body:    body,
	}, nil
}

func (r *Response) Status() int {
	return r.status
}

func (r *Response) Header(key string) string {
	return r.headers.Get(key)
}

func (r *Response) Headers() map[string]string {
	return getHeaders(r.headers)
}

func (r *Response) Cookie(key string) string {
	for _, cookie := range r.cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}

	return ""
}

func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.Body())
}

func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}
