package frame

import (
	"encoding/json"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/validatorService"
	"reflect"
)

const (
	ResponseTypeJson = 1
	ResponseTypeXml  = 3
	CodeSuccess      = 0
)

var (
	resEmptySlice []interface{}
)

type ResJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	resEmptySlice = make([]interface{}, 0)
}
func (this *Base) Init(data interface{}) {

	// 请求参数
	switch this.ctx.Method() {
	case "GET":
		this.initParamGet(data)
	default:
		this.initParamPost(data)
		this.initParamGet(data)
	}
	this.ValidateParam(data)
}

// ValidateParam 参数校验
func (this *Base) ValidateParam(data interface{}) {
	// 加载验证参数
	err := validatorService.GetTranslations().ValidateParam(data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

// InitAndBackParam Init返回前端传回的参数
func (this *Base) InitAndBackParam(data interface{}) (param map[string]interface{}) {
	param = make(map[string]interface{})
	err := this.ctx.ReadJSON(&param)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 请求参数
	switch this.ctx.Method() {
	case "GET", "PUT", "DELETE":
		this.initParamGet(data)
	}

	marshal, err := json.Marshal(param)
	if err != nil {
		exce.ThrowSys(err)
	}
	err = json.Unmarshal(marshal, &data)
	if err != nil {
		exce.ThrowSys(err)
	}
	this.ValidateParam(data)
	return
}

func (this *Base) ReplaceParamColumn(data, old interface{}, requestParam map[string]interface{}) map[string]interface{} {
	columnName := reflect.TypeOf(data)
	columnValue := reflect.ValueOf(data)
	finalParam := make(map[string]interface{})
	oldcolumnName := reflect.TypeOf(old)
	oldcolumnValue := reflect.ValueOf(old)

	for i := 0; i < oldcolumnName.NumField(); i++ {
		name := oldcolumnName.Field(i).Name

		if _, ok := requestParam[name]; ok {
			value := oldcolumnValue.Field(i).Interface()
			requestParam[name] = value
		}
	}

	for i := 0; i < columnName.NumField(); i++ {
		name := columnName.Field(i).Name
		value := columnValue.Field(i).String()
		if _, ok := requestParam[name]; ok {
			finalParam[value] = requestParam[name]
		}
	}
	return finalParam
}

func (this *Base) initParamPost(data interface{}) {
	err := this.ctx.ReadJSON(&data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

func (this *Base) initParamGet(data interface{}) {
	if this.ctx.Params().Len() == 0 {
		return
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		panic(exce.CodeRequestError)
	}

	for i := 0; i < val.NumField(); i++ {
		key := val.Type().Field(i).Name
		keyType := val.Type().Field(i).Type.Kind()
		switch keyType {
		case reflect.String:
			if this.ctx.Params().Exists(key) {
				value := this.ctx.Params().Get(key)
				val.Field(i).SetString(value)
			}
		case reflect.Uint32:
			if this.ctx.Params().Exists(key) {
				value, _ := this.ctx.Params().GetUint64(key)
				val.Field(i).SetUint(value)
			}
		case reflect.Int:
			if this.ctx.Params().Exists(key) {
				value, _ := this.ctx.Params().GetInt64(key)
				val.Field(i).SetInt(value)
			}
		case reflect.Bool:
			if this.ctx.Params().Exists(key) {
				value, _ := this.ctx.Params().GetBool(key)
				val.Field(i).SetBool(value)
			}
		case reflect.Float64:
			if this.ctx.Params().Exists(key) {
				value, _ := this.ctx.Params().GetFloat64(key)
				val.Field(i).SetFloat(value)
			}
		case reflect.Int64:
			if this.ctx.Params().Exists(key) {
				value, _ := this.ctx.Params().GetInt64(key)
				val.Field(i).SetInt(value)
			}
		}

	}
}

func (this *Base) Page() int {
	return this.ctx.Params().GetIntDefault("Page", 1)
}

func (this *Base) PageSize() int {
	return this.ctx.Params().GetIntDefault("PageSize", 10)
}

func (this *Base) Skip() int {
	return (this.Page() - 1) * this.PageSize()
}
