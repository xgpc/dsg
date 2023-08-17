package dsg

import (
	"github.com/xgpc/dsg/v2/pkg/etcd"
	"github.com/xgpc/dsg/v2/pkg/mysql"
	"github.com/xgpc/dsg/v2/pkg/redis"
)

var Conf Config

type Config struct {
	TLS    string       `mapstructure:"tls"`
	Etcd   etcd.Config  `mapstructure:"etcd"`
	DBInfo mysql.DBInfo `mapstructure:"db_info"`
	Redis  redis.Config `mapstructure:"redis"`
	JwtKey string       `mapstructure:"jwt_key"`
	AesKey string       `mapstructure:"aes_key"`
}
