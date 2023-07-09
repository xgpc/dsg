package dsg

import (
	"github.com/xgpc/dsg/v2/pkg/etcd"
	"github.com/xgpc/dsg/v2/pkg/mysql"
	"github.com/xgpc/dsg/v2/pkg/redis"
)

type Config struct {
	TLS    string       `yaml:"tls"`
	Etcd   etcd.Config  `yaml:"etcd"`
	DBInfo mysql.DBInfo `yaml:"db_info"`
	Redis  redis.Config `yaml:"redis"`
	JwtKey string       `yaml:"jwt_key"`
	AesKey string       `yaml:"aes_key"`
}
