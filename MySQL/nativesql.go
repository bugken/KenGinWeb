package nativeSQL

import (
	"database/sql"
	"fmt"
	//_ 是包引用操作,匿名导入,只会执行包下各模块中的init方法,并不会真正的导入包,所以不可以调用包内的其他方法.
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	age  int
	name string
}

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}

	// 数据库相关设置
	//db.SetConnMaxIdleTime(time.Second * 5) // 超时时间
	//db.SetMaxIdleConns(2)                  // 最大空闲连接数
	//db.SetMaxOpenConns(10)                 // 最大连接数
	return nil
}

// 查询单条数据示例
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多条数据示例
func QueryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}
