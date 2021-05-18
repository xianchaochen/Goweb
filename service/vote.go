package service

import (
	"bluebell/common/redis"
	"errors"
	"math"
	"time"
)

const (
	secondsInOneWeek = 3600 * 24 * 7
	scorePerVote = 430
)


var (
	ErrOutOfVotingTime = errors.New("投票时间已过")
)

// 每个帖子自发版后一个星期内允许用户投票 到期讲redis保存的赞成和发对票数存储到mysql 删除redis key
func PostVote(user_id string, postId string, value float64) (err error) {
	// 判断帖子是否还能投票
	// 获取帖子创建时间
	postTime := redis.Client.ZScore(redis.GetKey(redis.KeyPostTime), postId).Val()
	if float64(time.Now().Unix())-postTime > secondsInOneWeek {
		return ErrOutOfVotingTime
	}

	// 更新帖子分数
	// 获取用户之前的投票记录
	cacheVal := redis.Client.ZScore(redis.GetKey(redis.KeyPostVotedPrefix+postId), user_id).Val()
	var op float64
	if value > cacheVal  {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(cacheVal - value)

	// 事务
	pipeline := redis.Client.TxPipeline()
	pipeline.ZIncrBy(redis.GetKey(redis.KeyPostScore), op*diff*scorePerVote, postId)
	// 记录用户投票信息
	key := redis.GetKey(redis.KeyPostVotedPrefix+postId)
	if value == 0 {
		pipeline.ZRem(key, postId)
	} else {
		pipeline.ZAdd(key, redis.ZData(value, user_id))
	}

	_, err = pipeline.Exec()
	return err
}
