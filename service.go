// Package dsg
// @Author:        asus
// @Description:   $
// @File:          New
// @Data:          2022/2/2118:09
//
package dsg

import (
	redis2 "github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/pkg/mysql"
	"github.com/xgpc/dsg/pkg/redis"
	"gorm.io/gorm"
)

var (
	// mysql
	_db *gorm.DB
	// redis
	_redis *redis2.Client
	// jwt
)

type option func() error

func MysqlOption(info mysql.DBInfo) func() error {
	return func() error {
		_db = mysql.New(info)
		return nil
	}
}

func RedisOption(config redis.Config) func() error {
	return func() error {
		_redis = redis.New(config)
		return nil
	}
}

// Default dsg初始化后，可以通过dsg.xxx调用各个子功能
func Default(optins ...option) {
	// 加载配置
	for _, opt := range optins {
		err := opt()
		if err != nil {
			panic(err)
		}

	}
}
