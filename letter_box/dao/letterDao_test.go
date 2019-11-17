package dao

import (
	"aaa/letter_box/dao/DB"
	"fmt"
	"testing"
)

func TestSearchSearchLetByUUID1and2(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID1 := "asdasd1"
	userID2 := "aaaaaaa"

	msg, err := SearchLetByUUID1and2(userID1, userID2)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
