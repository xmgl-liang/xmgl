package dao

import (
	"aaa/mail_box/dao/DB"
	"fmt"
	"testing"
)

func TestSearchUserByUUID(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID := "job.UserUUID"

	msg, err := SearchUserByUUID(userID)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
