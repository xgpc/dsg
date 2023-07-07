package env

import (
	"fmt"
	"github.com/xgpc/dsg/v2/pkg/mysql"
	"github.com/xgpc/dsg/v2/pkg/redis"
	"io/ioutil"
	path2 "path"

	"gopkg.in/yaml.v3"
)

var Config Conf

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

type Conf struct {
	App           *App           `yaml:"app"`
	Mysql         *mysql.DBInfo  `yaml:"db_info"`
	Redis         *redis.Config  `yaml:"redis"`
	Message       *Message       `yaml:"message"`
	Wechat        *Wechat        `yaml:"wechat"`
	Microservices *Microservices `yaml:"microServices"`
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
