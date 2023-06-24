package util

import (
	"github.com/xgpc/dsg/pkg/json"
)

func StructToMap(obj interface{}) map[string]interface{} {
	body, err := json.Encode(obj)
	if err != nil {
		panic(err)
	}
	var md map[string]interface{}
	err = json.Decode(body, &md)
	if err != nil {
		panic(err)
	}
	return md
}
