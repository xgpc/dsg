// Package util
// @Author:        asus
// @Description:   $
// @File:          struct
// @Data:          2022/4/119:31
//
package util

import "reflect"

func StructToMapByRef(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if objT.Kind() == reflect.Interface || objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
		objV = objV.Elem()
	}

	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	return
}

// ToMapByJson 通过json转为map
func ToMapByJson(obj interface{}) (data map[string]interface{}) {
	marshal, _ := JsonEncode(obj)
	data = make(map[string]interface{})
	JsonDecode(marshal, &data)

	return data
}
