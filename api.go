package dsg

import (
	"encoding/base64"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2/exce"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"time"
)

//type ApiInterface interface {
//    Handler() int
//}

//type Api struct {
//    Method     string
//    Name       string
//    Login      bool
//    AuthKey    string
//    Controller func(base *Base) ApiInterface
//}

type ResSign struct {
	RandomStr string
	SignAt    int64
	Sign      string
}

type Base struct {
	// 请求
	ctx iris.Context `json:"-"`
}

func NewBase(ctx iris.Context) *Base {
	return &Base{ctx: ctx}
}

func (p *Base) Ctx() iris.Context {
	return p.ctx
}

func (p *Base) SetMyId(id uint32) {
	p.ctx.Values().Set("mid", id)
}

func (p *Base) MyId() uint32 {
	id, err := p.ctx.Values().GetUint32("mid")
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, err.Error())
	}
	return id
}

func (p *Base) MyIdToString() string {
	id := p.MyId()
	return strconv.Itoa(int(id))
}

func (p *Base) Token() (token string) {
	token = p.ctx.GetHeader("Token")
	return
}

// Key 从header中获取前端公钥
func (p *Base) Key() (rsa []byte) {
	s := p.ctx.GetHeader("Key")
	if s == "" {
		exce.ThrowSys(exce.CodeRequestError, "需要携带秘钥访问")
	}
	rsa, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(exce.CodeSysBusy)
	}
	return
}

func (p *Base) Success() {
	p.resp(map[string]interface{}{
		"Code": CodeSuccess,
		"Msg":  "ok",
	})
}

func (p *Base) SuccessWithData(data interface{}) {
	var res = &ResJson{Code: CodeSuccess, Msg: "ok"}

	res.Data = data

	p.resp(res)
}

func (p *Base) SuccessWithList(list interface{}, total interface{}) {
	if reflect.TypeOf(list).Kind() == reflect.Slice && reflect.ValueOf(list).Len() < 1 {
		list = resEmptySlice
	}
	var res = &ResJson{
		CodeSuccess, "查询成功", iris.Map{
			"Total": total,
			"List":  list,
		},
	}
	p.resp(res)
}

func (p *Base) resp(data interface{}) {
	err := p.Ctx().JSON(data)
	if err != nil {
		panic(err)
		return
	}
	return
}

// 时间

func (p *Base) Now() time.Time {
	return time.Now()
}

func (p *Base) NowUnix() int64 {
	return time.Now().Unix()
}

// DB 默认
func (p *Base) DB() *gorm.DB {
	return _db
}
func (p *Base) Redis() *redis.Client {
	return _redis
}
