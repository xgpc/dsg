package dsg

import (
	"encoding/base64"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2/exce"
	"gorm.io/gorm"
	"reflect"
	"strconv"
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

type UpBase struct {
	Base
	param map[string]interface{}
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
	id, err := p.ctx.Values().GetUint32("mid")
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, err.Error())
	}
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
		"code": CodeSuccess,
		"msg":  "ok",
	})
}

func (p *Base) SuccessWithData(data interface{}) {
	var res = &ResJson{Code: CodeSuccess, Msg: "ok"}
	if data == nil {
		res.Data = resEmptySlice
	} else {
		res.Data = data
	}
	p.resp(res)
}

func (p *Base) SuccessWithList(list interface{}, total interface{}) {
	if reflect.TypeOf(list).Kind() == reflect.Slice && reflect.ValueOf(list).Len() < 1 {
		list = resEmptySlice
	}
	var res = &ResJson{
		CodeSuccess, "查询成功", iris.Map{
			"total": total,
			"list":  list,
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

// DB 默认
func (p *Base) DB() *gorm.DB {
	return _db
}
func (p *Base) Redis() *redis.Client {
	return _redis
}

//// CryptSend 加密：后端--->前端
//func (p *Base) CryptSend(data []byte) {
//    AesKey, err := cryptService.GenKeyByte(16)
//    if err != nil {
//        exce.ThrowSys(err)
//    }
//    iv, err := cryptService.GenKeyByte(16)
//    if err != nil {
//        exce.ThrowSys(err)
//    }
//    encrypted, err := cryptService.AesEncrypt(data, AesKey, iv)
//    if err != nil {
//        exce.ThrowSys(err)
//    }
//    //对AES秘钥加密
//    rsaEncrypt := cryptService.RSAEncrypt(AesKey, p.Key())
//
//    p.SuccessWithData(map[string]interface{}{
//        "key":  rsaEncrypt,
//        "tag":  iv,
//        "data": encrypted,
//    })
//
//}

//// CryptReceive 加密：后端--->前端
//func (p *Base) CryptReceive() []byte {
//    var param struct {
//        // RSA公钥加密AES后的密文
//        Key []byte `validate:"required"`
//        // AES偏移量
//        IV []byte `validate:"required" json:"Tag"`
//        // AES加密数据后的密文
//        Data []byte `validate:"required"`
//    }
//    p.Init(&param)
//
//    //------>
//    // 使用RSA私钥解出AES秘钥
//    decryptAESKey := cryptService.RSADecrypt(param.Key, cryptService.RSAKey.Private)
//    origin, err := cryptService.AesDecrypt(param.Data, decryptAESKey, param.IV)
//    if err != nil {
//        exce.ThrowSys(err)
//    }
//    //<----------
//    return origin
//}
