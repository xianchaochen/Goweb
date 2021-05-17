package repository

import (
	mysql2 "bluebell/common/mysql"
	"bluebell/config"
	"bluebell/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IUserRepository interface {
	Conn() error
	CheckUserExist(userName string) (bool)
	Insert(user *entity.User) (userID int64, err error)
	FindUserByUsername(username string) (user *entity.User)
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
		mysql, errMysql := mysql2.NewMysqlConn(config.GlobalConfig.MysqlConfig)
		if errMysql != nil {
			zap.L().Error(fmt.Sprintf("Conn Mysql failed,err%v\n", errMysql))
			return errMysql
		}
		u.mysqlConn = mysql
	}

	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserManagerRepository) CheckUserExist(username string) (bool) {
	if err := u.Conn(); err != nil {
		return true
	}
	user := u.FindUserByUsername(username)
	if user != nil {
		return true
	}

	return false
}

func (u *UserManagerRepository) Insert(user *entity.User) (userID int64, err error) {
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


func (u *UserManagerRepository) FindUserByUsername(username string) *entity.User {
	if err := u.Conn(); err != nil {
		return nil
	}

	var user entity.User
	sql := "Select id,user_id,username,password,email,gender,create_time,update_time from " + u.table + " Where username=?"
	err := u.mysqlConn.Get(&user, sql, username)
	if err != nil {
		zap.L().Error(fmt.Sprintf("FindUserByUsername failed,err:%v\n", err))
		return nil
	}

	return &user
}
