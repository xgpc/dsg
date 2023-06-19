package redis

import (
	"bytes"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Control struct {
	Conn *redis.Client
}

type Config struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Db          int    `yaml:"db"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"poolSize"`
	MinIdleConn int    `yaml:"MinIdleConn"`
}

// New 返回redis client 需要自行保存
func New(conf Config) *redis.Client {
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

// MakeControl 通过redis client 创建控制器
func MakeControl(client *redis.Client) *Control {
	return &Control{Conn: client}
}

func (p *Control) ErrRedisGet(key string) (string, bool) {
	if p.Conn == nil {
		// LogError("redis get -> redis nil", key)
		return "", false
	}

	str, err := p.Conn.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		// LogError("redis get -> error:"+err.Error(), key)
		return "", false
	}

	return str, true
}

func (p *Control) RedisGet(key string) string {
	str, flag := p.ErrRedisGet(key)
	if !flag {
		//log.ThrowSys("RedisGet", key)
	} //	todo
	return str
}

//
//// ***************************** set *****************************
//

func (p *Control) ErrRedisSet(key, value string, sec int) bool {
	if p == nil {
		// LogError("redis set -> redis nil", key+" => "+value)
		return false
	}

	err := p.Conn.Set(ctx, key, value, time.Duration(sec)*time.Second).Err()
	if err != nil {
		// LogError("redis set -> error:"+err.Error(), key+" => "+value)
		return false
	}
	return true
}

func (p *Control) RedisSet(key, value string, sec int) {
	if !p.ErrRedisSet(key, value, sec) {
		//log.ThrowSys("RedisSet", key+" => "+value)
	} //	todo
}

//// ***************************** del *****************************

func (p *Control) ErrRedisDel(key string) bool {
	if p.Conn == nil {
		// LogError("redis del -> redis nil", key)
		return false
	}

	err := p.Conn.Del(ctx, key).Err()
	if err != nil {
		// LogError("redis del -> error:"+err.Error(), key)
		return false
	}
	return true
}

func (p *Control) RedisDel(key string) {
	if !p.ErrRedisDel(key) {
		//log.ThrowSys("RedisDel", key)
	} //	todo
}

//// ***************************** scan *****************************

func (p *Control) ErrRedisScan(cursor uint64, match string, count int64) ([]string, uint64, bool) {
	if p == nil {
		// LogError("redis scan -> redis nil", match)
		return nil, 0, false
	}

	keys, newCursor, err := p.Conn.Scan(ctx, cursor, match, count).Result()
	if err == redis.Nil {
		return []string{}, 0, true
	}

	if err != nil {
		// LogError("redis scan -> error:"+err.Error(), match)
		return nil, 0, false
	}

	return keys, newCursor, true
}

func (p *Control) RedisScan(cursor uint64, match string, count int64) ([]string, uint64) {
	keys, newCursor, flag := p.ErrRedisScan(cursor, match, count)
	if !flag {
		//log.ThrowSys("RedisScan", match)
	} //	todo
	return keys, newCursor
}

//// ***************************** hscan *****************************

func (p *Control) ErrRedisHScan(key string, cursor uint64, match string, count int64) ([]string, uint64, bool) {
	if p == nil {
		// LogError("redis hscan -> redis nil", match)
		return nil, 0, false
	}

	keys, newCursor, err := p.Conn.HScan(ctx, key, cursor, match, count).Result()
	if err == redis.Nil {
		return []string{}, 0, true
	}

	if err != nil {
		// LogError("redis hscan -> error:"+err.Error(), match)
		return nil, 0, false
	}

	return keys, newCursor, true
}

func (p *Control) RedisHScan(key string, cursor uint64, match string, count int64) ([]string, uint64) {
	keys, newCursor, flag := p.ErrRedisHScan(key, cursor, match, count)
	if !flag {
		//log.ThrowSys("RedisHScan", match)
	} //	todo
	return keys, newCursor
}

//// ***************************** hget *****************************

func (p *Control) ErrRedisHGet(key, field string) (string, bool) {
	if p == nil {
		// LogError("redis hget -> redis nil", key+" => "+field)
		return "", false
	}

	str, err := p.Conn.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		// LogError("redis hget -> error:"+err.Error(), key+" => "+field)
		return "", false
	}

	return str, true
}

func (p *Control) RedisHGet(key, field string) string {
	str, flag := p.ErrRedisHGet(key, field)
	if !flag {
		//log.ThrowSys("RedisHGet", key+" => "+field)
	} //	todo
	return str
}

//// ***************************** hdel *****************************

func (p *Control) ErrRedisHDel(key, field string) bool {
	if p == nil {
		// LogError("redis hdel -> redis nil", key+" => "+field)
		return false
	}
	err := p.Conn.HDel(ctx, key, field).Err()
	if err != nil {
		// LogError("redis hdel -> error:"+err.Error(), key+" => "+field)
		return false
	}

	return true
}

func (p *Control) RedisHDel(key, field string) {
	if !p.ErrRedisHDel(key, field) {
		//log.ThrowSys("RedisHDel", key+" => "+field)
	} //	todo
}

//// ***************************** hset *****************************

func (p *Control) ErrRedisHSet(key, field, value string) bool {
	if p == nil {
		// LogError("redis hset -> redis nil", key+" => "+field)
		return false
	}

	err := p.Conn.HSet(ctx, key, field, value).Err()
	if err != nil {
		// LogError("redis hset -> error:"+err.Error(), key+" => "+field)
		return false
	}
	return true
}

func (p *Control) RedisHSet(key, field, value string) {
	if !p.ErrRedisHSet(key, field, value) {
		//log.ThrowSys("RedisHSet", key+" => "+field)
	} //	todo
}

//// ***************************** rpush *****************************

func (p *Control) ErrRedisPush(key, value string) bool {
	if p == nil {
		// LogError("redis rpush -> redis nil", key)
		return false
	}

	err := p.Conn.RPush(ctx, key, value).Err()
	if err != nil {
		// LogError("redis rpush -> error:"+err.Error(), key+" => "+value)
		return false
	}
	return true
}

func (p *Control) RedisPush(key, value string) {
	if !p.ErrRedisPush(key, value) {
		//log.ThrowSys("RedisPush", key)
		//	todo
	}
}

//// ***************************** lpop *****************************

func (p *Control) ErrRedisPop(key string) (string, bool) {
	if p == nil {
		// LogError("redis lpop -> redis nil", key)
		return "", false
	}

	bs, err := p.Conn.LPop(ctx, key).Bytes()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		// LogError("redis lpop -> error:"+err.Error(), key)
		return "", false
	}
	return string(bs), true
}

func (p *Control) RedisPop(key string) string {
	str, flag := p.ErrRedisPop(key)
	if !flag {
		//log.ThrowSys("RedisPop", key)
		//	todo
	}
	return str
}

// ***************************** lrange *****************************

func (p *Control) ErrRedisLRange(key string, start, end int64) ([]string, bool) {
	if p == nil {
		// LogError("redis lrange -> redis nil", key)
		return []string{}, false
	}

	datum, err := p.Conn.LRange(ctx, key, start, end).Result()
	if err == redis.Nil {
		return []string{}, true
	}
	if err != nil {
		// LogError("redis lrange -> error:"+err.Error(), key)
		return []string{}, false
	}
	return datum, true
}

func (p *Control) RedisLRange(key string, start, end int64) []string {
	datum, flag := p.ErrRedisLRange(key, start, end)
	if !flag {
		//log.ThrowSys("RedisLRange", key)
	} //	todo
	return datum
}

// ***************************** expire *****************************

func (p *Control) ErrRedisExpire(key string, lifetime int) bool {
	if p == nil {
		// LogError("redis expire -> redis nil", key)
		return false
	}

	err := p.Conn.Expire(ctx, key, time.Duration(lifetime)*time.Second).Err()
	if err != nil {
		// LogError("redis expire -> error:"+err.Error(), key)
		return false
	}
	return true
}

func (p *Control) RedisExpire(key string, lifetime int) {
	if !p.ErrRedisExpire(key, lifetime) {
		//log.ThrowSys("RedisExpire", key)
	} //	todo
}

// ***************************** exists *****************************

func (p *Control) ErrRedisExists(key string) (bool, bool) {
	if p == nil {
		// LogError("redis exists -> redis nil", key)
		return false, false
	}

	v, err := p.Conn.Exists(ctx, key).Result()
	if err != nil {
		// LogError("redis exists -> error:"+err.Error(), key)
		return false, false
	}
	return v > 0, true
}

func (p *Control) RedisExists(key string) bool {
	f, flag := p.ErrRedisExists(key)
	if !flag {
		//log.ThrowSys("RedisExists", key)
	} //	todo
	return f
}

//// ***************************** hexists *****************************

func (p *Control) ErrRedisHExists(key, field string) (bool, bool) {
	if p == nil {
		// LogError("redis hexists -> redis nil", key)
		return false, false
	}

	v, err := p.Conn.HExists(ctx, key, field).Result()
	if err != nil {
		// LogError("redis hexists -> error:"+err.Error(), key)
		return false, false
	}
	return v, true
}

func (p *Control) RedisHExists(key, field string) bool {
	f, flag := p.ErrRedisHExists(key, field)
	if !flag {
		//log.ThrowSys("RedisHExists", key)
	} //	todo
	return f
}

// ***************************** incr *****************************

func (p *Control) ErrRedisIncr(key string) (int64, bool) {
	if p == nil {
		// LogError("redis incr -> redis nil", key)
		return 0, false
	}

	v, err := p.Conn.Incr(ctx, key).Result()
	if err != nil {
		// LogError("redis incr -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (p *Control) RedisIncr(key string) int64 {
	v, flag := p.ErrRedisIncr(key)
	if !flag {
		//log.ThrowSys("RedisIncr", key)
	} //	todo
	return v
}

// ***************************** sadd *****************************

func (p *Control) ErrRedisSadd(key string, members ...interface{}) (int64, bool) {
	if p == nil {
		// LogError("redis Sadd -> redis nil", key)
		return 0, false
	}

	v, err := p.Conn.SAdd(ctx, key, members...).Result()
	if err != nil {
		// LogError("redis Sadd -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (p *Control) RedisSadd(key string, members ...interface{}) int64 {
	v, flag := p.ErrRedisSadd(key, members...)
	if !flag {
		//log.ThrowSys("RedisSadd", key)
	} //	todo
	return v
}

func (p *Control) ErrRedisSMembers(key string) ([]string, bool) {
	if p == nil {
		// LogError("redis SMembers -> redis nil", key)
		return nil, false
	}

	v, err := p.Conn.SMembers(ctx, key).Result()
	if err != nil {
		// LogError("redis SMembers -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (p *Control) RedisSMembers(key string) []string {
	v, flag := p.ErrRedisSMembers(key)
	if !flag {
		//log.ThrowSys("RedisSMembers", key)
	} //	todo
	return v
}

// set 类型
//***************************** 有序合集 set *****************************
func (p *Control) ErrRedisZADD(key string, members ...*redis.Z) (int64, bool) {
	if p == nil {
		// LogError("redis ZAdd -> redis nil", key)
		return 0, false
	}

	v, err := p.Conn.ZAdd(ctx, key, members...).Result()
	if err != nil {
		// LogError("redis ZAdd -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (p *Control) RedisZADD(key string, members ...*redis.Z) int64 {
	v, flag := p.ErrRedisZADD(key, members...)
	if !flag {
		//log.ThrowSys("RedisZIncrBy", key)
		return 0
	}
	return v
}

func (p *Control) ErrRedisZIncrBy(key string, increment float64, member string) (float64, bool) {
	if p == nil {
		// LogError("redis ZIncrBy -> redis nil", key)
		return 0, false
	}

	v, err := p.Conn.ZIncrBy(ctx, key, increment, member).Result()
	if err != nil {
		// LogError("redis ZIncrBy -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (p *Control) RedisZIncrBy(key string, increment float64, member string) float64 {

	v, flag := p.ErrRedisZIncrBy(key, increment, member)
	if !flag {
		//log.ThrowSys("RedisZIncrBy", key)
		return 0
	}
	return v
}

//***************************** ZRange *****************************
func (p *Control) ErrRedisZRange(key string, start, stop int64) ([]string, bool) {
	if p == nil {
		// LogError("redis ZRange -> redis nil", key)
		return nil, false
	}

	v, err := p.Conn.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		// LogError("redis ZRange -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (p *Control) RedisZRange(key string, start, stop int64) []string {

	v, flag := p.ErrRedisZRange(key, start, stop)
	if !flag {
		//log.ThrowSys("RedisZRange", key)
		return nil
	}
	return v
}

//***************************** ZRevRange *****************************
func (p *Control) ErrRedisZRevRange(key string, start, stop int64) ([]string, bool) {
	if p == nil {
		// LogError("redis ZRevRange -> redis nil", key)
		return nil, false
	}

	v, err := p.Conn.ZRevRange(ctx, key, start, stop).Result()
	if err != nil {
		// LogError("redis ZRevRange -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (p *Control) RedisZRevRange(key string, start, stop int64) []string {

	v, flag := p.ErrRedisZRevRange(key, start, stop)
	if !flag {
		//log.ThrowSys("RedisZRange", key)
		return nil
	}
	return v
}

//***************************** ZRevRangeWithScores *****************************
func (p *Control) ErrRedisZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, bool) {
	if p == nil {
		// LogError("redis ZRevRangeWithScores -> redis nil", key)
		return nil, false
	}

	v, err := p.Conn.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		// LogError("redis ZRevRangeWithScores -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (p *Control) RedisZRevRangeWithScores(key string, start, stop int64) []redis.Z {

	v, flag := p.ErrRedisZRevRangeWithScores(key, start, stop)
	if !flag {
		//log.ThrowSys(1000, "RedisZRange", key)
		return nil
	}
	return v
}
