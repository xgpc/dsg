package etcd

import (
	"time"
)

// 123
type Config struct {
	Name    string
	Address string
	Port    int

	Endpoints            []string      `yaml:"endpoints"`
	AutoSyncInterval     time.Duration `yaml:"autoSyncInterval"`
	DialTimeout          time.Duration `yaml:"dialTimeout"`    // 秒
	DefLeaseSecond       time.Duration `yaml:"defLeaseSecond"` // 续租时间(秒)
	DialKeepAliveTime    time.Duration `yaml:"dialKeepAliveTime"`
	DialKeepAliveTimeout time.Duration `yaml:"dialKeepAliveTimeout"`
}
