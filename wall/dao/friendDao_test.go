package dao

import (
	"aaa/wall/dao/DB"
	"fmt"
	"testing"
)

func TestSearchFriByUUID(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID := "job.UserUUID"

	msg, err := SearchFriByUUID(userID)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
