package tasks

import (
	"aaa/letter_box_worker/common"
	"aaa/letter_box_worker/tasks/DB"
	"fmt"
	"testing"
)

func TestAddLetterJob(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.LetterJob{
		Name:     "ma",
		LetterID: "2021-03-01",
		UUID1:    "asdasd1",
		UUID2:    "aaaaaaa",
		UserName: "assad",
		Content:  "12asdasd3",
		Label:    "asd",
		WallType: "love0",
	}

	err := AddLetterJob(job)
	fmt.Println("测试结果是：", err)
}

func TestDeleteDelFriend(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.DelFriend{
		Name:  "ma",
		UUID1: "asdasd1",
		UUID2: "aaaaaaa",
	}

	err := DeleteDelFriend(job)
	fmt.Println("测试结果是：", err)
}

func TestUpdateNumFri(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.UpdateNumFriJob{
		Name:   "ma",
		UUID1:  "job.UserUUID",
		NumFri: 1,
	}

	err := UpdateNumFri(job)
	fmt.Println("测试结果是：", err)
}

func TestUpdateNumPublish(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.UpdateNumPublishJob{
		Name:       "ma",
		UUID1:      "job.UserUUID",
		NumPublish: 1,
	}

	err := UpdateNumPublish(job)
	fmt.Println("测试结果是：", err)
}
