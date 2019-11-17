package tasks

import (
	"aaa/mail_box_worker/common"
	"aaa/mail_box_worker/tasks/DB"
	"fmt"
	"testing"
)

func TestAddMailFriend(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.MailFriend{
		Name:  "ma",
		UUID1: "asdasd1",
		UUID2: "aaaaaaa",
	}

	err := AddMailFriend(job)
	fmt.Println("测试结果是：", err)
}

func TestAddLetterByFri(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.LetterByFri{
		Name:     "ma",
		LetterID: "2017-03-01",
		UUID1:    "asdasd1",
		UUID2:    "aaaaaaa",
		UserName: "assad",
		Content:  "12asa3",
		Label:    "asd",
		WallType: "love0",
	}

	err := AddLetterByFri(job)
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

func TestUpdateNoLetter(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	job := &common.NoLetter{
		Name:     "ma",
		UUID1:    "asdasd1",
		UUID2:    "aaaaaaa",
		LetterID: "2017-03-01 00:00:00",
		Label:    "1",
	}

	err := UpdateNoLetter(job)
	fmt.Println("测试结果是：", err)
}
