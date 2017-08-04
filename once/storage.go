package once

import "github.com/garyburd/redigo/redis"

var (
	domain string
	redisPool *redis.Pool
)