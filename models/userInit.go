package models

import (
	"github.com/xgpc/dsg/frame"
	"log"
	"os"
)

const (
	checkSQL  = `select count(*) from information_schema.tables where table_name=? and TABLE_SCHEMA=?;`
	createSQL = `create table user
(
    id     int unsigned auto_increment comment '主键'
        primary key,
    open_id varchar(32) null comment '微信openID',

	created_at int unsigned comment '创建时间',
	updated_at int unsigned comment '更新时间',
    constraint user_id_uindex
        unique (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment '用户表';`
	tableName = "user"
)

func InitUser() {
	db := frame.DB
	dbName := frame.Config.Mysql.Database
	var i int64
	err := db.Raw(checkSQL, tableName, dbName).Scan(&i).Error
	if err != nil {
		log.Fatal("生成user表错误,", err)
	}

	if i < 1 {
		frame.DB.Exec(createSQL)
		GenModel()
		os.Mkdir("models/cmd", 0777)
		//	todo os.OpenFile()

	}
}
