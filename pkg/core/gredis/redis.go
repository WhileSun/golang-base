package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"log"
	"time"
)

type RedisConfig struct {
	Server      string
	Password    string
	MaxActive   int
	MaxIdle     int
	IdleTimeout int
	Select      int
}

type GRedis struct {
	RedisPool *redis.Pool
}

func New(redisKey string) *GRedis {
	if redisKey == "" {
		redisKey = "redis"
	}
	redisConfig := &RedisConfig{
		Server:      "localhost:6379",
		Password:    "",
		MaxIdle:     2,
		MaxActive:   5,
		IdleTimeout: 240,
		Select:      0,
	}
	gconfig.Config.UnmarshalKey(redisKey, redisConfig)
	redisPool := redisConfig.initRedisPool()
	return &GRedis{RedisPool: redisPool}
}

func (redisConfig *RedisConfig) initRedisPool() *redis.Pool {
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
			// 密码验证
			if redisConfig.Password != "" {
				if _, err := c.Do("AUTH", redisConfig.Password); err != nil {
					c.Close()
					log.Fatalf("gredis AUTH Error:%s", err)
					return nil, err
				}
			}
			// SELECT获取
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
