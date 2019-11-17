package tasks

import (
	"aaa/wall_worker/common"
	"aaa/wall_worker/tasks/DB"
	"fmt"
	"testing"
)

func TestAddWallLetterJob(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.WallLetterJob{
		Name:     "ma",
		LetterID: "2019-03-01",
		UUID1:    "asdasd",
		UUID2:    "0",
		UserName: "assad",
		Content:  "123",
		Label:    "asd",
		WallType: "love0",
	}

	err := AddWallLetterJob(job)
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

func TestUpdateNumCharge(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.UpdateNumChargeJob{
		Name:      "ma",
		UUID1:     "job.UserUUID",
		NumCharge: 1,
	}

	err := UpdateNumCharge(job)
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

func TestAddWallFriend(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.WallFriend{
		Name:  "ma",
		UUID1: "job.UserUUID",
		UUID2: "job.UserUUID2",
	}

	err := AddWallFriend(job)
	fmt.Println("测试结果是：", err)
}
