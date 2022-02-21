package sysService

import (
	"context"
	"github.com/xgpc/dsg/frame"
	"time"
)

func GetSetSysVersion() (int, error) {
	conn := frame.RedisDefault()
	result, err := conn.Exists(context.Background(), "sysVersion").Result()
	if err != nil {
		return 0, err
	}
	//不存在
	if result <= 0 {
		InitSysVersion()
	}
	return Get()
}

func Get() (int, error) {
	i, err := frame.RedisDefault().Get(context.Background(), "sysVersion").Int()
	return i, err
}

func InitSysVersion() {
	unix := time.Now().Unix()
	frame.RedisDefault().Set(context.Background(), "sysVersion", unix, time.Hour*24*300)
}
