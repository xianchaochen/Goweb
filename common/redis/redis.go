package redis

import (
	"bluebell/config"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

const (
	Prefix             = "bluebell:"
	KeyPostTime        = "post:time"   // zset 帖子及发帖时间
	KeyPostScore       = "post:score"  // zset 帖子及投票分数
	KeyPostVotedPrefix = "post:voted:" // zset 记录用户及投票类型,参数是post_id
)

func init() {
	Client, _ = NewRedisConn(config.GlobalConfig.RedisConfig)
}

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

func GetKey(key string) string  {
	return Prefix +key
}




