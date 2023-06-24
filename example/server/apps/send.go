package apps

type Body struct {
	Method        string `json:"method"` // 请求数据类型 Register | Heartbeat
	Body          string `json:"body"`
	RouterAppName string `json:"router_app_name"`
	IpAddress     string `json:"ip_address"`
	AppKey        string `json:"app_key"`
}

func SendRegister() {

}

func SendHeartbeat() {

}
