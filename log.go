package dsg

import (
	"encoding/json"
	"github.com/xgpc/dsg/env"
	"github.com/xgpc/dsg/exce"
	"reflect"
)

type ErrorLevel int

const (
	LogLevelInfo ErrorLevel = 1 + iota
	LogLevelError
)

var ErrLevel = map[ErrorLevel]string{
	LogLevelInfo:  "info",
	LogLevelError: "error",
}

func LogInfo(msg string, data interface{}) {
	go logHandle(LogLevelInfo, msg, &data)
}

func LogError(msg string, data interface{}) {
	go logHandle(LogLevelError, msg, &data)
}

func logHandle(level ErrorLevel, msg string, data *interface{}) {
	sysLogLevel := env.Config.SysConfig.LogLevel
	if sysLogLevel == ErrLevel[level] || level == LogLevelError {
		exce.Write("[" + ErrLevel[level] + "] " + msg + ":" + dealCon(data) + "\n")
	}
}

func dealCon(content *interface{}) string {
	switch reflect.TypeOf(content).Kind() {
	case reflect.String:
		return reflect.ValueOf(content).String()
	default:
		j, err := json.Marshal(content)
		if err != nil {
			panic(err.Error())
		}
		return string(j)
	}
	return ""
}
