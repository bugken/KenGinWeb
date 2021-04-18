package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

// 定义一个初始化数据库的函数
func Init() (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetInt("mysql.port"),
		viper.GetString("mysql.db_name"))
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
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) // 最大空闲连接数
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) // 最大连接数

	return nil
}

func Close() {
	_ = db.Close()
}
