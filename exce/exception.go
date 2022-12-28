package exce

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"reflect"
	"runtime"
)

type DsgError int

func (d DsgError) Error() string {
	return ErrString[d]
}
func (d DsgError) Code() int {
	return int(reflect.ValueOf(d).Int())
}

const (
	CodeOK = 0
)
const (
	CodeSysBusy DsgError = 1001 + iota
)

//用户类
const (
	CodeUserError DsgError = 2001 + iota
	CodeUserNoLogin
	CodeUserNoAuth
)

const (
	//	请求

	CodeRequestError DsgError = 3001 + iota
)

var ErrString = map[DsgError]string{
	CodeOK: "ok",
	//系统
	CodeSysBusy: "[error] 系统繁忙，请稍后重试",

	//	用户
	CodeUserError:   "[error] 用户类错误",
	CodeUserNoLogin: "[error] 请先登录",
	CodeUserNoAuth:  "[error] 抱歉，您的角色没有此功能的操作权限，请联系技术人员",

	CodeRequestError: "[error] 请求错误，请检查参数",
}

type SysException struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ThrowSys 抛出错误
func ThrowSys(err error, args ...interface{}) {
	num := len(args)
	var e SysException

	var t = reflect.TypeOf(err).Kind()
	if t != reflect.Int {
		defer func() {
			if err := recover(); err != nil {
				log.Println("日志写入失败：" + err.(error).Error())
			}
		}()
		pc, _, line, ok := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		if !ok {

		}
		if err != nil {
			errMsg := fmt.Sprintf("[error] at %s:%d Cause by: %s\n", f.Name(), line, err.Error())
			Write(errMsg)
		}
		e.Code = CodeSysBusy.Code()
		e.Msg = ErrString[CodeSysBusy]
		panic(e)
	}
	temp := err.(DsgError)
	switch num {
	case 0:
		e.Code = temp.Code()
		e.Msg = temp.Error()
	case 1:
		msg := args[0]
		e.Code = temp.Code()
		e.Msg = reflect.ValueOf(msg).String()
	}
	panic(e)
}

func DealException(ctx iris.Context, err interface{}) {
	switch err.(type) {
	case SysException:
		sysErr := err.(SysException)
		err := ctx.JSON(map[string]interface{}{"code": sysErr.Code, "msg": sysErr.Msg})
		if err != nil {
			return
		}
	default:
		ExceptionCode(ctx, err)
		return
	}
}

func ExceptionCode(ctx iris.Context, t interface{}) {
	switch t.(type) {
	default:
		var res = map[string]interface{}{
			"code": CodeSysBusy.Code(),
			"msg":  CodeSysBusy.Error(),
		}
		err := ctx.JSON(res)
		if err != nil {
			return
		}
	}
}
