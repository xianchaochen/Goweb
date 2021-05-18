package entity

type ParamRegister struct {
	Username   string `json:"username" binding:"required"` // 不能为空
	Password   string `json:"password" binding:"required"` // 不能为空
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 不能为空,必须和上面相等
}

type ParamLogin struct {
	Username string `json:"username" binding:"required" example:"kuso123321"` // 不能为空
	Password string `json:"password" binding:"required" example:"123"` // 不能为空
}

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