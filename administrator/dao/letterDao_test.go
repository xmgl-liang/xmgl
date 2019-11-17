package dao

import (
	"aaa/administrator/dao/DB"
	"fmt"
	"testing"
)

func TestSearchSearchLetByUUID1and2(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID1 := "-2"

	msg, err := SearchLetterById(userID1)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
