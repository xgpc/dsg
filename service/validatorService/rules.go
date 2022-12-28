// Package frame
// @Author:        asus
// @Description:   $
// @File:          registerValidation
// @Data:          2021/12/311:16
//
package validatorService

import (
	"regexp"

	"github.com/xgpc/dsg/util"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//Rules 校验方法
var Rules = map[string]validator.Func{
	"mobile": func(fl validator.FieldLevel) bool {
		res, _ := regexp.MatchString("^1[3-9]\\d{9}$", fl.Field().String())
		return res
	},
}

//RulesMsg 校验返回的错误信息
var RulesMsg = map[string]util.RegisterTrans{
	"mobile": {
		RegisterTranslationsFunc: func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}不是手机号格式", true)
		},
		TranslationFunc: func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		},
	},
}
