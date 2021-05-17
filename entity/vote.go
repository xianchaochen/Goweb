package entity

// 帖子投票数据
type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`                // 帖子id string
	Direction int8  `json:"direction,string" binding:"required,oneof=-1 1 0"` // 赞成1or反对-1 0取消投票
}
