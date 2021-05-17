package mysql

import (
	"bluebell/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewMysqlConn(cfg *config.MysqlConfig) ( db *sqlx.DB,err error)  {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("Connect DB failed", zap.Error(err))
		return
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return
}

func Close(db *sqlx.DB)  {
	err := db.Close()
	if err != nil {
		zap.L().Error("Close DB failed", zap.Error(err))
	}
	return
}