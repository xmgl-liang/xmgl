package dao

import (
	"aaa/letter_box/dao/DB"
	"aaa/letter_box/model"
	"fmt"
	"testing"
)

func TestSearchUserByUUID(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	userID1 := "asdasd1"

	msg, err := SearchUserByUUID(userID1)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}

func TestSearchFriMsgByID(t *testing.T) {
	//连接数据库
	if err := DB.InitDB(); err != nil {
		fmt.Println("测试结果是：", err)
	}

	var fris []*model.Friend

	fri := &model.Friend{
		MeUUID:  "asdasd1",
		FriUUID: "aaaaaaa",
	}

	fris = append(fris, fri)

	user1 := &model.User{
		Friends: fris,
	}

	msg, err := SearchFriMsgByID(user1)
	fmt.Println("测试结果是：", err)
	fmt.Println("测试结果是：", msg)
}
