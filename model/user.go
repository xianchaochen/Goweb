package model

type User struct {
	ID         int64  `db:"id"`
	UserID     int64  `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Email      string `db:"email"`
	Gender     int    `db:"gender"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
