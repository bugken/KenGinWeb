package sql

import (
	"fmt"
	"testing"
)

func TestFuncX(t *testing.T){
	err := InitDBX() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init dbx failed,err:%v\n", err)
		return
	}
	fmt.Printf("connect dbx success.\n")

	//QueryRowDemoX()
	//QueryMultiRowDemoX()

	//InsertRowDemoX()
	//DeleteRowDemoX()
	//UpdateRowDemoX()

	//InsertUserDemoX()
	//NamedQueryX()
	TransactionDemoX()
}

