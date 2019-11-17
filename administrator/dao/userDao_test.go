package dao

import (
	"aaa/administrator/dao/DB"
	"fmt"
	"testing"
)

func TestUpdateUserMsg(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	err := UpdateUserMsg()
	fmt.Println("测试结果是：", err)
}
