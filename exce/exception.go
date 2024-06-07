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
	CodeSysConfigError
	CodeSysDatabaseError
	CodeSysCacheError
	CodeSysFileError
	CodeSysDependencyError
)

// 用户类
const (
	CodeUserError DsgError = 2001 + iota
	CodeUserNoLogin
	CodeUserNoAuth
	CodeUserTokenError
	CodeUserPasswordError
	CodePermissionError
	CodePermissionDenied
)

// 请求相关
const (
	CodeRequestError DsgError = 3001 + iota
	CodeRequestTimeout
	CodeRequestTooMany
	CodeRequestAPINotFound
	CodeRequestServiceNotAvailable
)

// 数据库相关
const (
	CodeDBError DsgError = 4001 + iota
	CodeDBInsertError
	CodeDBUpdateError
	CodeDBDeleteError
	CodeDBQueryError
)

// CodeUnknownError 其他通用错误
const (
	CodeUnknownError DsgError = 6001
)

var ErrString = map[DsgError]string{
	CodeOK:                 "ok",
	CodeSysBusy:            "[error] 系统繁忙，请稍后重试",
	CodeSysConfigError:     "[error] 系统配置错误",
	CodeSysDatabaseError:   "[error] 数据库访问错误",
	CodeSysCacheError:      "[error] 缓存访问错误",
	CodeSysFileError:       "[error] 文件操作错误",
	CodeSysDependencyError: "[error] 依赖服务不可用",
	CodeUserError:          "[error] 用户类错误",
	CodeUserNoLogin:        "[error] 用户未登录",
	CodeUserNoAuth:         "[error] 抱歉，您的角色没有此功能的操作权限，请联系技术人员(开发人员级别权限)",
	CodeUserTokenError:     "[error] token无效或已过期",
	CodeUserPasswordError:  "[error] 密码错误",
	CodePermissionError:    "[error] 权限错误(接口权限)",
	CodePermissionDenied:   "[error] 权限被拒绝(层级权限)",

	CodeRequestError:               "[error] 请求参数错误(给出明确的错误问题)",
	CodeRequestTimeout:             "[error] 请求超时",
	CodeRequestTooMany:             "[error] 请求过多",
	CodeRequestAPINotFound:         "[error] 请求接口不存在",
	CodeRequestServiceNotAvailable: "[error] 请求服务不存在",

	CodeDBError:       "[error] 数据库错误",
	CodeDBInsertError: "[error] 数据插入失败",
	CodeDBUpdateError: "[error] 数据更新失败",
	CodeDBDeleteError: "[error] 数据删除失败",
	CodeDBQueryError:  "[error] 数据查询失败",

	CodeUnknownError: "[error] 未知错误",
}

type SysException struct {
	Code int
	Msg  string
}

// ThrowSys 抛出自定义错误
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
	default:
		e.Code = temp.Code()
		e.Msg = fmt.Sprint(args)
	}
	panic(e)
}

func DealException(ctx iris.Context, err interface{}) {
	switch err.(type) {
	case SysException:
		sysErr := err.(SysException)
		err := ctx.JSON(map[string]interface{}{"Code": sysErr.Code, "Msg": sysErr.Msg})
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
			"Code": CodeSysBusy.Code(),
			"Msg":  CodeSysBusy.Error(),
		}
		err := ctx.JSON(res)
		if err != nil {
			return
		}
	}
}
