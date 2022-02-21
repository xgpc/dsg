package frame

import (
	"bytes"
	"github.com/xgpc/dsg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func MySqlDefault() *gorm.DB {
	return DB
}

func loadMysql() {
	DB = mysqlInit(dsg.Config.Mysql)
}

func mysqlInit(conf dsg.Mysql) *gorm.DB {
	//连接
	mysqlConfig := conf
	var connStr bytes.Buffer
	connStr.WriteString(mysqlConfig.User)
	connStr.WriteString(":")
	connStr.WriteString(mysqlConfig.Password)
	connStr.WriteString("@(")
	connStr.WriteString(mysqlConfig.Host)
	connStr.WriteString(":")
	connStr.WriteString(mysqlConfig.Port)
	connStr.WriteString(")/")
	connStr.WriteString(mysqlConfig.Database)
	connStr.WriteString("?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local")

	db, err := gorm.Open(mysql.Open(connStr.String()), &gorm.Config{})
	if db == nil {
		panic("[error] 连接失败")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。

	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnMaxLifetime) * time.Second)

	//connections are not closed due to a connection's idle time.

	sqlDB.SetConnMaxIdleTime(time.Duration(mysqlConfig.ConnMaxIdleTime) * time.Second)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("mysql 启动失败", err.Error())
	}
	return db
}
