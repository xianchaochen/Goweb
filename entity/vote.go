package entity

// 帖子投票数据
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`                // 帖子id string
	Direction int8  `json:"direction,string" binding:"oneof=-1 1 0"` // 赞成1or反对-1 0取消投票
}
