package models

import (
	"github.com/enjoy322/ormtool"
	"github.com/enjoy322/ormtool/base"
	"github.com/xgpc/dsg/frame"
)

// GenModel 生成数据库表对应的结构体
func GenModel() {
	//从config.yaml获取配置
	m := frame.Config.Mysql
	ormtool.GenerateMySQL(
		base.MysqlConfig{
			User:     m.User,
			Password: m.Password,
			Host:     m.Host,
			Port:     m.Port,
			Database: m.Database},
		base.Config{
			SavePath:       "./models/model.go",
			IsGenJsonTag:   true,
			IsGenInOneFile: true,
			// 1：不生成数据库基本信息 2：生成简单的数据库字段信息
			GenDBInfoType: 1,
			//# json tag类型，前提：IsGenJsonTag:true. 1.UserName 2.userName 3.user_name 4.user-name
			JsonTagType: 3,
			// 是否生成建表语句
			IsGenCreateSQL: true,
			// 自定义对应类型，优先选择
			CustomType: map[string]string{
				"int":          "int",
				"int unsigned": "uint32",
				"tinyint(1)":   "bool",
			}})
}
