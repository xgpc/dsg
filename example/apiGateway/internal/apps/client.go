package apps

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var _quit chan bool

func runTask() {
	// 注册服务
	ticker := time.NewTicker(time.Minute * 5)
	_quit = make(chan bool)

	// 首次需要手动注册, 如果注册失败也就不需要启动了
	register()

	go func() {
		for {

			select {

			case <-ticker.C:

				sendRequest(_conf.RootHost, getBody(), nil)

			case <-_quit:
				fmt.Println("quit")
			}
		}
	}()

}

func register() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}()

	sendRequest(_conf.RootHost, getBody(), nil)
}

func sendRequest(url string, body map[string]interface{}, Headers map[string]string) []byte {

	// 防止低筒崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	data := Request(url, http.MethodPost, body, Headers)
	return data
}

func Quit() {
	time.Sleep(time.Second * 5)
	_quit <- true
}

func getBody() map[string]interface{} {
	md := map[string]interface{}{}

	//Method        string `json:"method"` // 请求数据类型 Register | Heartbeat
	//Body          string `json:"body"`
	//RouterAppName string `json:"router_app_name"`
	//IpAddress     string `json:"ip_address"`
	//AppKey        string `json:"app_key"`

	md["app_key"] = GetUUid()
	md["method"] = "Register"
	md["body"] = ""
	md["router_app_name"] = _conf.RouterAppName
	md["ip_address"] = _conf.IpAddress + ":" + _conf.Port

	return md
}
