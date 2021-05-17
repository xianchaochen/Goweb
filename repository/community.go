package repository

import (
	mysql2 "bluebell/common/mysql"
	"bluebell/config"
	"bluebell/entity"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ICommunityRepository interface {
	Conn() error
	SelectCommunityList() (list []*entity.Community, err error)
}

func NewCommunityRepository(table string) ICommunityRepository {
	if table == "" {
		table = "community"
	}
	return &CommunityManagerRepository{table, nil}
}

type CommunityManagerRepository struct {
	table     string
	mysqlConn *sqlx.DB
}

func (c *CommunityManagerRepository) Conn() (err error) {
	if c.mysqlConn == nil {
		mysql, errMysql := mysql2.NewMysqlConn(config.GlobalConfig.MysqlConfig)
		if errMysql != nil {
			zap.L().Error(fmt.Sprintf("Conn Mysql failed,err%v\n", errMysql))
			return errMysql
		}
		c.mysqlConn = mysql
	}
	if c.table == "" {
		c.table = "community"
	}
	return
}

func (c *CommunityManagerRepository)SelectCommunityList() (list []*entity.Community, err error) {
	if err := c.Conn(); err != nil {
		return nil, err
	}
	sqlString := "Select community_id,community_name from " + c.table
	if err = c.mysqlConn.Select(&list, sqlString);err!=nil {
		fmt.Println("%+v", err)
		if err == sql.ErrNoRows {
			err = nil
		}
		return nil, err
	}
	return list, err
}