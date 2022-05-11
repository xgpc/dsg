package frame

import (
	"bytes"
	"context"
	"github.com/go-redis/redis/v8"
)

var rDB *redis.Client

func RedisDefault() *redis.Client {
	return rDB
}
func LoadRedis() {
	rDB = redisInit(Config.Redis)
}

func redisInit(conf Redis) *redis.Client {
	var addr bytes.Buffer

	addr.WriteString(conf.Host)
	addr.WriteString(":")
	addr.WriteString(conf.Port)

	conn := redis.NewClient(&redis.Options{
		Addr:         addr.String(),
		Password:     conf.Password,
		DB:           conf.Db,
		PoolSize:     conf.PoolSize,
		MinIdleConns: conf.MinIdleConn,
	})
	err := conn.Ping(context.Background()).Err()
	if err != nil {
		panic("Redis启动失败，" + err.Error())
	}
	//使用0号数据库
	conn.Do(context.Background(), "select", 0)
	return conn
}
