package service

import (
	"bluebell/common/redis"
	"errors"
	"time"
)

const secondsInOneWeek = 3600 * 24 * 7

var (
	ErrOutOfVotingTime = errors.New("投票时间已过")
)

// 每个帖子自发版后一个星期内允许用户投票 到期讲redis保存的赞成和发对票数存储到mysql 删除redis key
func PostVote(user_id int64, postId string, direction float64) (err error) {
	// 判断帖子是否还能投票

	// 获取帖子创建时间
	postTime := redis.Client.ZScore(redis.GetKey(redis.KeyPostTime), postId).Val()
	if float64(time.Now().Unix())-postTime > secondsInOneWeek {
		return ErrOutOfVotingTime
	}

	// 更新帖子分数

	// 记录用户投票信息

	return nil
}
