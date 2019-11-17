package dao

import (
	"aaa/mail_box/dao/DB"
	"fmt"
	"testing"
)

func TestSearchLetterByIDandTime(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID := "aaaaaaa"
	time1 := "2021-02-01 00:00:00"

	msg, err := SearchLetterByIDandTime(userID, time1)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
