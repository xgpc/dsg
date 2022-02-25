package frame

import (
	"fmt"
	"io/ioutil"
	path2 "path"

	"gopkg.in/yaml.v3"
)

var Config conf

const defConfigPath = "config.yaml"

// LoadConf 读取Yaml配置文件
func LoadConf(path ...string) {

	var configPath string
	if len(path) == 0 {
		configPath = defConfigPath
	} else {
		configPath = path2.Join(path...)
	}
	c := &Config
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type conf struct {
	App       App       `yaml:"app"`
	Mysql     Mysql     `yaml:"mysql"`
	Redis     Redis     `yaml:"redis"`
	Message   Message   `yaml:"message"`
	Wechat    Wechat    `yaml:"wechat"`
	SysConfig SysConfig `yaml:"sysConfig"`
}

type App struct {
	AppName string `yaml:"appName"`
	CnName  string `yaml:"cnName"`
	Port    string `yaml:"port"`
	RPCAddr string `yaml:"rpcAddr"`
	SysCode uint32 `yaml:"sysCode"`
	TLS     string `yaml:"tls"`
}
type Redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Db          int    `yaml:"db"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"poolSize"`
	MinIdleConn int    `yaml:"MinIdleConn"`
}

type Mysql struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Database        string `yaml:"database"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	MaxIdleConn     int    `yaml:"maxIdleConn"`
	MaxOpenConn     int    `yaml:"maxOpenConn"`
	ConnMaxIdleTime int    `yaml:"connMaxIdleTime"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}
type Message struct {
	AccessKeyId     string
	AccessKeySecret string
}

type Wechat struct {
	AppID string `yaml:"appID"`
}

// SysConfig 系统配置
type SysConfig struct {
	// 定时器
	StartSchedule bool   `yaml:"startSchedule"`
	LogLevel      string `yaml:"logLevel"`
}
