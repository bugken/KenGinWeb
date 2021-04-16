package nativeSQL

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	err := InitDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Printf("connect db success.\n")

	//InsertRowDemo()
	//DeleteRowDemo()
	//UpdateRowDemo()
	//QueryRowDemo()
	//QueryMultiRowDemo()

	PrepareInsertDemo()
	PrepareQueryDemo()
}
