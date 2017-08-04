package once

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"errors"
)

type Configuration struct {
	Domain  string
	RedisConf *RedisConfiguration
}

type RedisConfiguration struct {
	Host string
	Port int
	SessionDB int
}

func InitOnce(conf *Configuration) (err error) {
	if conf == nil {
		err = errors.New("Once.initRedisPool: redis configuration is empty")
		return
	}

	initDomain(conf.Domain)

	err = initRedisPool(conf.RedisConf)

	return
}

func initDomain(d string) {
	domain = d
}

func initRedisPool(redisConf *RedisConfiguration) (err error) {
	if redisConf == nil {
		err = errors.New("Once.initRedisPool: configuration is empty")
		return
	}

	redisPool = &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 5,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConf.Host)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("SELECT", redisConf.SessionDB)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return
}
