package dsg

import (
	"github.com/xgpc/dsg/pkg/mysql"
	"github.com/xgpc/dsg/pkg/redis"
)

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
