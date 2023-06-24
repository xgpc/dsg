package apps

import (
	redis2 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_db         *gorm.DB
	_rdb        *redis2.Client
	routers     map[string]string // Server
	routersUuid map[string]string // Server

	_conf Config // Client
	_uuid string // Client
)

type Config struct {
	RootHost      string `yaml:"root_host"`       // 服务端地址
	Port          string `yaml:"port"`            // 本机访问端口(自动获取)
	HostType      string `yaml:"host_type"`       // 服务器状态 server | client
	RouterAppName string `yaml:"router_app_name"` // 服务名称
	IpAddress     string `yaml:"ip_address"`      // 服务器地址
}

func GetUUid() string {
	if _uuid == "" {
		_uuid = uuid.New().String()
	}
	return _uuid
}

func Init(db *gorm.DB, rdb *redis2.Client, conf Config) {
	_db = db
	_rdb = rdb
	_conf = conf

	switch conf.HostType {
	case "server":
		//添加路由
		reloadRouter()
	case "client":
		runTask()
	}

}

func db() *gorm.DB {
	return _db
}

func redis() *redis2.Client {
	return _rdb
}
