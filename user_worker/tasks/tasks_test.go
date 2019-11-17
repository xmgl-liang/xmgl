package tasks

import (
	"aaa/user_worker/common"
	"aaa/user_worker/tasks/DB"
	"fmt"
	"testing"
)

func TestUpdateUserMsg(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.UserMsg{
		Name:     "ma",
		UserUUID: "job.UserUUID",
		UserName: "asdas",
		Sex:      "gaaS",
		Text:     "asfqw",
	}

	err := UpdateUserMsg(job)
	fmt.Println("测试结果是：", err)
}
