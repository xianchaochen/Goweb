package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"web_app/settings"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error)  {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:      config.DB,
		PoolSize: config.PoolSize,
	})

	_, err = rdb.Ping().Result()

	return
}

func Close()  {
	err := rdb.Close()
	if err != nil {
		zap.L().Error("Close rdb failed", zap.Error(err))
	}
	return
}
