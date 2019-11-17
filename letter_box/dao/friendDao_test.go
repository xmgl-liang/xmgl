package dao

import (
	"aaa/letter_box/dao/DB"
	"fmt"
	"testing"
)

func TestSearchFriByUUID(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID1 := "asdasd1"

	msg, err := SearchFriByUUID(userID1)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
