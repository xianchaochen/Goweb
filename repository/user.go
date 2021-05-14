package repository

import (
	"bluebell/common"
	"bluebell/config"
	"bluebell/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IUserRepository interface {
	Conn() error
	CheckUserExist(userName string) (bool, error)
	Insert(user *model.User) (userID int64, err error)
}

func NewUserRepository(userTable string) IUserRepository {
	if userTable == "" {
		userTable = "user"
	}
	return &UserManagerRepository{userTable, nil}
}

type UserManagerRepository struct {
	table     string
	mysqlConn *sqlx.DB
}

func (u *UserManagerRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, errMysql := common.NewMysqlConn(config.GlobalConfig.MysqlConfig)
		if errMysql != nil {
			zap.L().Error("数据库连接失败")
			return errMysql
		}
		u.mysqlConn = mysql
	}

	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserManagerRepository) CheckUserExist(username string) (bool, error) {
	if err := u.Conn(); err != nil {
		return false, err
	}
	sql := "select count(id) from " + u.table + " where username = ?"
	var result int
	err := u.mysqlConn.Get(&result, sql, username)
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (u *UserManagerRepository) Insert(user *model.User) (userID int64, err error) {
	if err = u.Conn(); err != nil {
		return 0, err
	}
	sql := "Insert into " + u.table + "(user_id,username,password) values(?,?,?)"
	_, err = u.mysqlConn.Exec(sql, user.UserID, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	return user.UserID, nil
}
