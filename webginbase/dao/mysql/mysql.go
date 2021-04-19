package mysql

import (
	"NetClassGinWeb/webginbase/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// 定义一个初始化数据库的函数
func Init(cfg *settings.MySQLConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return
	}

	// 数据库相关设置
	db.SetMaxIdleConns(cfg.MaxIdleConns) // 最大空闲连接数
	db.SetMaxOpenConns(cfg.MaxOpenConns) // 最大连接数

	return nil
}

func Close() {
	_ = db.Close()
}
