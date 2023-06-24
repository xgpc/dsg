package apps

import (
	bytes2 "bytes"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/util"
	"io"
	"net/http"
	"time"
)

func Request(url, Method string, body map[string]interface{}, Headers map[string]string) []byte {
	var data []byte
	if body != nil {
		marshal, err := util.JsonEncode(body)
		if err != nil {
			exce.ThrowSys(exce.CodeRequestError, err.Error())
		}

		data = marshal
	} else {
		// 没有的时候 返回空数据
		marshal, err := util.JsonEncode("{}")
		if err != nil {
			exce.ThrowSys(exce.CodeRequestError, err.Error())
		}
		data = marshal
	}

	bodyBuffer := bytes2.NewBuffer(data)

	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(Method, url, bodyBuffer)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	for k, v := range Headers {
		req.Header.Add(k, v)
	}

	rsp, err := client.Do(req)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	defer rsp.Body.Close()
	bytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	if rsp.StatusCode != http.StatusOK {
		exce.ThrowSys(exce.CodeRequestError, rsp.Status)
	}
	return (bytes)
}
