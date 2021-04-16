package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

func InitDBX() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	dbx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	dbx.SetMaxOpenConns(20)
	dbx.SetMaxIdleConns(10)
	return
}
