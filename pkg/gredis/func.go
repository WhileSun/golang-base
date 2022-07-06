package gredis

import (
	"github.com/gomodule/redigo/redis"
)

func setExpire(conn redis.Conn, key string, exTime int) {
	conn.Do("expire", key, exTime)
}

func (g *GRedis) Exists(key string) (reply int, err error){
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = redis.Int(conn.Do("EXISTS", key))
	return
}

func (g *GRedis) Keys(pattern string) (reply []string, err error){
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = redis.Strings(conn.Do("KEYS", pattern))
	return
}

func (g *GRedis) Set(key string, exTime int, val interface{}) (reply interface{}, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = conn.Do("SET", key, val)
	setExpire(conn, key, exTime)
	return
}

func (g *GRedis) Get(key string) (reply interface{}, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = conn.Do("GET", key)
	return
}

func (g *GRedis) Del(key string) (reply interface{}, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = conn.Do("DEL", key)
	return
}

func (g *GRedis) Hmset(key string, exTime int, args ...interface{}) (reply interface{}, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	args = append([]interface{}{key}, args...)
	reply, err = conn.Do("HMSET", args...)
	setExpire(conn, key, exTime)
	return
}

func (g *GRedis) Hgetall(key string) (reply map[string]string,err error){
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = redis.StringMap(conn.Do("HGETALL", key))
	return
}

func (g *GRedis) Sadd(key string,  exTime int, args ...interface{}) (reply interface{}, err error){
	conn := redisPool.Get()
	defer conn.Close()
	args = append([]interface{}{key}, args...)
	reply, err = conn.Do("Sadd", args...)
	if exTime>0 {
		setExpire(conn, key, exTime)
	}
	return
}

func (g *GRedis) Sismember(key string,member string) (reply int, err error){
	conn := redisPool.Get()
	defer conn.Close()
	reply, err = redis.Int(conn.Do("SISMEMBER", key, member))
	return
}
