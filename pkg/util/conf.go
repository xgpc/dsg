/**
 * @Author: smono
 * @Description:
 * @File:  conf
 * @Version: 1.0.0
 * @Date: 2022/6/30 9:16
 */

package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
)

const defConfigPath = "config.yml"

// LoadYmlConf 读取Yaml配置文件
func LoadYmlConf(out interface{}, fileName ...string) {

	var configPath string
	if len(fileName) == 0 {
		configPath = defConfigPath
	} else {
		configPath = path.Join(fileName...)
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		panic(err.Error())
	}
}
