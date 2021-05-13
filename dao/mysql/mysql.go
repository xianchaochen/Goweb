package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"web_app/settings"
)

var db *sqlx.DB

func Init(config *settings.MysqlConfig) (err error)  {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("Connect DB failed", zap.Error(err))
		return
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	return
}

func Close()  {
	err := db.Close()
	if err != nil {
		zap.L().Error("Close DB failed", zap.Error(err))
	}

	return
}