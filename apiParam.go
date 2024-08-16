package dsg

import (
	"encoding/json"
	"github.com/xgpc/dsg/v2/exce"
	"github.com/xgpc/dsg/v2/pkg/validator"
	"net/http"
	"reflect"
)

const (
	CodeSuccess = 0
)

var (
	resEmptySlice []interface{}
)

type ResJson struct {
	Code int
	Msg  string
	Data interface{}
}

func init() {
	resEmptySlice = make([]interface{}, 0)
}
func (p *Base) Init(data interface{}) {

	// 请求参数
	switch p.ctx.Method() {
	case http.MethodGet, http.MethodDelete:
		p.initParamGet(data)
	default:
		p.initParamPost(data)
	}
	p.ValidateParam(data)
}

func (p *Base) CheckBody(data interface{}) map[string]interface{} {
	ReqByte, err := p.Ctx().GetBody()
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	if err := json.Unmarshal(ReqByte, &data); err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var mdData map[string]interface{}
	if err := json.Unmarshal(ReqByte, &mdData); err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 把data转为json, 然后再转为map[string], 如果mdData中的key不在data中, 则删除
	marshal, err := json.Marshal(data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	var tmp map[string]interface{}
	err = json.Unmarshal(marshal, &tmp)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	for k := range mdData {
		if _, ok := tmp[k]; !ok {
			delete(mdData, k)
		}
	}

	p.ValidateParam(data)
	return mdData
}

// ValidateParam 参数校验
func (p *Base) ValidateParam(data interface{}) {
	// 加载验证参数
	err := validator.GetTranslations().ValidateParam(data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

// InitAndBackParam Init返回前端传回的参数
func (p *Base) InitAndBackParam(data interface{}) (param map[string]interface{}) {
	param = make(map[string]interface{})
	err := p.ctx.ReadJSON(&param)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	// 请求参数
	switch p.ctx.Method() {
	case "GET", "PUT", "DELETE":
		p.initParamGet(data)
	}

	marshal, err := json.Marshal(param)
	if err != nil {
		exce.ThrowSys(err)
	}
	err = json.Unmarshal(marshal, &data)
	if err != nil {
		exce.ThrowSys(err)
	}
	p.ValidateParam(data)
	return
}

func (p *Base) ReplaceParamColumn(data, old interface{}, requestParam map[string]interface{}) map[string]interface{} {
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

func (p *Base) initParamPost(data interface{}) {
	err := p.ctx.ReadJSON(&data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

func (p *Base) initParamGet(data interface{}) {
	err := p.ctx.ReadQuery(data)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
}

func (p *Base) Page() int64 {
	return p.ctx.URLParamInt64Default("page", 1)
}

func (p *Base) PageSize() int64 {
	return p.ctx.URLParamInt64Default("page_size", 10)
}

func (p *Base) Skip() int64 {
	return (p.Page() - 1) * p.PageSize()
}
