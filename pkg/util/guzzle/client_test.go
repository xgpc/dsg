// Package guzzle
// @Author:        asus
// @Description:   $
// @File:          client_test
// @Data:          2022/1/2016:50
//
package guzzle

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(WithHost("http://localhost"))

	response, err := client.RequestJSON(map[string]interface{}{
		"id":   "1",
		"name": "zhangsan",
	}).Post("/api/test")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response.String())
}
