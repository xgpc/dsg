package frame

import (
	"encoding/base64"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/cryptService"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type ApiInterface interface {
	Handler() int
}

type Api struct {
	Method     string
	Name       string
	Login      bool
	AuthKey    string
	Controller func(base *Base) ApiInterface
}

type ResSign struct {
	RandomStr string
	SignAt    int64
	Sign      string
}

type Base struct {
	// 请求
	ctx             iris.Context `json:"-"`
	Method          string       `json:"-"`
	Path            string       `json:"-"`
	isMustJsonParam bool
	Time            time.Time `json:"-"`
	TimeStamp       int64     `json:"-"`

	// 参数签名
	RandomStr string `valid:"length(32|32)"`
	SignAt    int64  `valid:"timestamp"`
	Sign      string `valid:"-"`

	// 会话
	//Session *Session `json:"-"`
	//MyId    uint32   `json:"-"`
	dbBegin bool
}

func NewBase(ctx iris.Context) *Base {
	return &Base{ctx: ctx}
}

type UpBase struct {
	Base
	param map[string]interface{}
}

func (this *Base) Ctx() iris.Context {
	return this.ctx
}

func (this *Base) MyId() int {
	return this.ctx.Values().GetIntDefault("mid", 0)
}

func (this *Base) Token() (token string) {
	token = this.ctx.GetHeader("Token")
	return
}

// Key 从header中获取前端公钥
func (this *Base) Key() (rsa []byte) {
	s := this.ctx.GetHeader("Key")
	if s == "" {
		exce.ThrowSys(exce.CodeReqParamMissing, "需要携带秘钥访问")
	}
	rsa, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(exce.CodeSysBusy)
	}
	return
}

func (this *Base) Success() {
	this.resp(map[string]interface{}{
		"Code": CodeSuccess,
		"Msg":  "ok",
	})
}

func (this *Base) SuccessWithData(data interface{}) {
	var res = &ResJson{Code: CodeSuccess, Msg: "ok"}
	if data == nil {
		res.Data = resEmptySlice
	} else {
		res.Data = data
	}
	this.resp(res)
}

func (this *Base) SuccessWithList(list interface{}, total int64) {
	if reflect.TypeOf(list).Kind() == reflect.Slice && reflect.ValueOf(list).Len() < 1 {
		list = resEmptySlice
	}
	var res = &ResJson{
		CodeSuccess, "查询成功", iris.Map{
			"Total": total,
			"List":  list,
		},
	}
	this.resp(res)
}

func (this *Base) resp(data interface{}) {
	_, err := this.ctx.JSON(data)
	if err != nil {
		panic(err)
		return
	}
	return
}

func (this *Base) CancelMustJsonParam() {
	this.isMustJsonParam = false
}

// DB 默认
func (this *Base) DB() *gorm.DB {
	return MySqlDefault()
}
func (this *Base) Redis() *redis.Client {
	return RedisDefault()
}

// CryptSend 加密：后端--->前端
func (this *Base) CryptSend(data []byte) {
	AesKey, err := cryptService.GenKeyByte(16)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	iv, err := cryptService.GenKeyByte(16)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	encrypted, err := cryptService.AesEncrypt(data, AesKey, iv)
	if err != nil {
		panic(err)
	}
	//对AES秘钥加密
	rsaEncrypt := cryptService.RSAEncrypt(AesKey, this.Key())

	this.SuccessWithData(map[string]interface{}{
		"Key":  rsaEncrypt,
		"Tag":  iv,
		"Data": encrypted,
	})

}

// CryptReceive 加密：后端--->前端
func (this *Base) CryptReceive() []byte {
	var param struct {
		// RSA公钥加密AES后的密文
		Key []byte `validate:"required"`
		// AES偏移量
		IV []byte `validate:"required" json:"Tag"`
		// AES加密数据后的密文
		Data []byte `validate:"required"`
	}
	this.Init(&param)

	//------>
	// 使用RSA私钥解出AES秘钥
	decryptAESKey := cryptService.RSADecrypt(param.Key, cryptService.RSAKey.Private)
	origin, err := cryptService.AesDecrypt(param.Data, decryptAESKey, param.IV)
	if err != nil {
		fmt.Println(err)
	}
	//<----------
	return origin
}
