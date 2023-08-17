/**
 * @Author: smono
 * @Description:
 * @File:  conf
 * @Version: 1.0.0
 * @Date: 2022/6/30 9:16
 */

package util

import (
	"github.com/spf13/viper"
	"path"
)

const defConfigPath = "config.yml"

// LoadYmlConf 读取Yaml配置文件
/*
 * 使用viper进行配置文件读取, 遇到_ 导致无法读取, 需要配置tag mapstructure
 */
func LoadYmlConf(out interface{}, fileName ...string) {

	var configPath string
	if len(fileName) == 0 {
		configPath = defConfigPath
	} else {
		configPath = path.Join(fileName...)
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(out)
	if err != nil {
		panic(err)
	}
}
