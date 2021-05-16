package common

import (
	"bluebell/config"
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedisConn(cfg *config.RedisConfig) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}
