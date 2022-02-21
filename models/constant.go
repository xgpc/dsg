package models

//loginType 登录环境类型

const (
	CodeLoginWeb = 1 + iota
	CodeLoginApp
	CodeLoginMini
	CodeLoginSub
)

var LoginType = map[int]string{
	CodeLoginWeb:  "web登录",
	CodeLoginApp:  "app登录",
	CodeLoginMini: "小程序登录",
	CodeLoginSub:  "公众号登录",
}
