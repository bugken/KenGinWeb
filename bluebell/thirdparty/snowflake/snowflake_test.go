package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	if err := Init("2021-04-19", 1); err != nil {
		fmt.Printf("init snowflake error:%v\n", err.Error())
		return
	}

	fmt.Printf("Generate ID:%v\n", GenID())
}
