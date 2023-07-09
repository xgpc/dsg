package env

import (
	"github.com/xgpc/dsg/v2/pkg/mysql"
	"github.com/xgpc/dsg/v2/pkg/redis"
)

var Config Conf

type Conf struct {
	App     *App          `yaml:"app"`
	Mysql   *mysql.DBInfo `yaml:"db_info"`
	Redis   *redis.Config `yaml:"redis"`
	Message *Message      `yaml:"message"`
	Wechat  *Wechat       `yaml:"wechat"`
}

type Microservices struct {
	FileAddr string `yaml:"fileAddr"`
	RPCAddr  string `yaml:"rpcAddr"`
}

type App struct {
	AppName string `yaml:"appName"`
	CnName  string `yaml:"cnName"`
	Port    string `yaml:"port"`
	SysCode uint32 `yaml:"sysCode"`
	TLS     string `yaml:"tls"`
}

type Message struct {
	AccessKeyId     string
	AccessKeySecret string
}

type Wechat struct {
	AppID string `yaml:"appID"`
}
