package exce

import (
	"github.com/kataras/iris/v12"
)

const (
	CodeOK = 0
)
const (
	CodeSysBusy = 1001 + iota
	CodeIllegalRequest
	CodeToHome
	CodeToResetPwd
	Code404
)

//用户类

const (
	CodeUserError = 2001 + iota
	CodeUserNoLogin
	CodeUserOperateError
	CodeUserPasswordError
	CodeUserVerifyCodeError
	CodeUserOtherLogin
	CodeUserInvalidLogin
	CodeUserNotExist
	CodeUserExisted
	CodeUserNoAuth
)

// 权限类

const (
	//请求

	CodeReqMethodError = 3001 + iota
	CodeReqNotExistAPI
	//	参数

	CodeRequestError = 4001 + iota
	CodeReqParamError
	CodeReqParamMissing
	CodeReqParamTypeError
)

var ErrString = map[int]string{
	CodeOK: "ok",
	//系统
	CodeSysBusy: "[error] 系统繁忙，请稍后重试",

	//	用户
	CodeUserNoLogin:         "[error] 请先登录",
	CodeUserPasswordError:   "[error] 密码错误",
	CodeUserVerifyCodeError: "[error] 验证码错误",
	CodeUserOtherLogin:      "[error] 账号已在其他设备登录",
	CodeUserInvalidLogin:    "[error] 登录失效，请重新登录",
	CodeUserNotExist:        "[error] 用户不存在",
	CodeUserExisted:         "[error] 用户已存在",
	CodeUserNoAuth:          "[error] 抱歉，您的角色没有此功能的操作权限，请联系技术人员",

	//请求
	CodeReqMethodError: "[error] 请求方式错误",
	CodeReqNotExistAPI: "[error] 接口不存在",

	//参数
	CodeReqParamError:     "[error] 参数错误",
	CodeReqParamMissing:   "[error] 参数缺失",
	CodeReqParamTypeError: "[error] 参数类型错误",
}

type SysException struct {
	Code int
	Msg  string
}

func ThrowSys(code int, msg string) {
	if _, ok := ErrString[code]; !ok {
		code = CodeSysBusy
	}
	panic(SysException{
		Code: code,
		Msg:  msg,
	})
}

func DealException(ctx iris.Context, err interface{}) {
	var e interface{}
	switch err.(type) {
	case SysException:
		sysErr := err.(SysException)
		_, err := ctx.JSON(map[string]interface{}{"Code": sysErr.Code, "Msg": sysErr.Msg})
		if err != nil {
			return
		}
		e = &sysErr
	default:
		ExceptionCode(ctx, err)
		return
	}
	defer logFile(LogLvError, e, nil)
}

func ExceptionCode(ctx iris.Context, t interface{}) {
	var e interface{}
	switch t.(type) {
	case int:
		var res = map[string]interface{}{
			"Code": t.(int),
			"Msg":  backMsg(t.(int)),
		}
		_, err := ctx.JSON(res)
		if err != nil {
			return
		}
		e = &res
	default:
		var res = map[string]interface{}{
			"Code": CodeSysBusy,
			"Msg":  backMsg(CodeSysBusy),
		}
		_, err := ctx.JSON(res)
		if err != nil {
			return
		}
		e = &res
	}
	defer logFile(LogLvError, e, nil)
}

func backMsg(code int) string {
	if v, ok := ErrString[code]; ok {
		return v
	}
	return ""
}
