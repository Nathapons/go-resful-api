package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-resful-api/internal/configs"
	"github.com/redis/go-redis/v9"
)

var Timeout = 1

func NewRedis(cfg *configs.RedisConfig) (RedisClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Timeout)*time.Second)
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Username: cfg.RedisUser,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDb,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Fatal("Connect to redis successfully")
	return &redisClient{rdb: rdb}, nil
}

func (c *redisClient) SetCtx(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Timeout)*time.Second)
	defer cancel()

	err := c.rdb.Set(ctx, key, value, 0).Err()
	return err
}

func (c *redisClient) GetCtx(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Timeout)*time.Second)
	defer cancel()

	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), value)
	if err != nil {
		return err
	}

	return nil
}

func (c *redisClient) RemoveCtx(key ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Timeout)*time.Second)
	defer cancel()

	err := c.rdb.Del(ctx, key...).Err()
	if err != nil {
		return err
	}

	return nil
}
