package entity

import "time"

type Post struct {
	ID int64 `json:"id" db:"post_id"` //前端数字int太大会有不精确，可以 `json:"id,string" db:"post_id"` 来回转string
	AuthID int64 `json:"auth_id" db:"auth_id"`
	CommunityID int64 `json:"community_id" db:"community_id"`
	Status int64 `json:"status" db:"status"`
	Title int64 `json:"title" db:"title"`
	Content int64 `json:"content" db:"content"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

// 帖子详情结构体 service查，拼接一个
type ApiPostDetail struct {
	AuthName string `json:"auth_name"`
	// 嵌入帖子结构体
	*Post `json:"post"`
	*Community `json:"community"` // 加了json会分层

}
