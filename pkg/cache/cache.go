package cache

import "github.com/redis/go-redis/v9"

type (
	redisClient struct {
		rdb *redis.Client
	}

	RedisClient interface {
		SetCtx(key string, value interface{}) error
		GetCtx(key string, value interface{}) error
		RemoveCtx(key ...string) error
	}
)
