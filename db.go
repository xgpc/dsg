package dsg

import (
	redis2 "github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/v2/pkg/mysql"
	"github.com/xgpc/dsg/v2/pkg/redis"
	"gorm.io/gorm"
)

var (
	// mysql
	_db *gorm.DB
	// redis
	_redis *redis2.Client
)

// DB 默认
func DB() *gorm.DB {
	return _db
}
func Redis() *redis2.Client {
	return _redis
}

func OptionMysql(info mysql.DBInfo) func() error {
	return func() error {
		_db = mysql.New(info)
		return nil
	}
}

func OptionRedis(config redis.Config) func() error {
	return func() error {
		_redis = redis.New(config)
		return nil
	}
}
