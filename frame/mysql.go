package frame

import (
	"bytes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func MySqlDefault() *gorm.DB {
	return DB
}

func loadMysql() {
	DB = mysqlInit(Config.Mysql)
}

func mysqlInit(conf Mysql) *gorm.DB {
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

	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为1ms
			SlowThreshold: 500 * 1000 * time.Microsecond,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel: logger.Warn,
		},
	)

	db, err := gorm.Open(mysql.Open(connStr.String()), &gorm.Config{
		Logger: slowLogger,
	})
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
