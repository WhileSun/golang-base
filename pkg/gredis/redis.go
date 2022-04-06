package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/whilesun/go-admin/pkg/gconf"
	"log"
	"time"
)

type redisConfigObj struct {
	Server      string
	Password    string
	MaxActive   int
	MaxIdle     int
	IdleTimeout int
	Select      int
}

var (
	redisPool   *redis.Pool
	redisConfig *redisConfigObj
)

func Run() {
	if redisPool != nil{
		return
	}
	redisConfig = &redisConfigObj{
		Server:      "localhost:6379",
		Password:    "",
		MaxIdle:     2,
		MaxActive:   5,
		IdleTimeout: 240,
		Select:      0,
	}
	gconf.Config.UnmarshalKey("redis", redisConfig)
	redisPool = initRedisPool()
	testPing()
}

func testPing(){
	redisConn :=redisPool.Get()
	defer redisConn.Close()
	redisConn.Do("PING")
}

//Get 从池里获取单个连接
func Get() redis.Conn {
	rc := redisPool.Get()
	return rc
}

//GetPool 获取进程池
func GetPool() *redis.Pool {
	return redisPool
}

func initRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle, //空闲数
		IdleTimeout: time.Duration(redisConfig.IdleTimeout) * time.Second,
		MaxActive:   redisConfig.MaxActive, //最大数
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.Server)
			if err != nil {
				log.Fatalf("redis Connect Error:%s", err)
				return nil, err
			}
			if redisConfig.Password != "" {
				if _, err := c.Do("AUTH", redisConfig.Password); err != nil {
					c.Close()
					log.Fatalf("gredis AUTH Error:%s", err)
					return nil, err
				}
			}
			_, err = redis.String(c.Do("SELECT", redisConfig.Select))
			if err != nil {
				log.Fatalf("gredis SELECT Error:%s", err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
