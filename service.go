// Package dsg
// @Author:        asus
// @Description:   $
// @File:          New
// @Data:          2022/2/2118:09
package dsg

import (
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	// mysql
	_db *gorm.DB
	// redis
	_redis *redis2.Client
	// jwt
	_jetKey string
)

type option func() error

// Default dsg初始化后，可以通过dsg.xxx调用各个子功能
func Default(list ...option) {
	// 加载配置
	for _, opt := range list {
		err := opt()
		if err != nil {
			panic(err)
		}
	}
}
