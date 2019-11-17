package common

import (
	"encoding/json"
)

//////////////////////Letter_box///////////////////////////////

//回复信件
type LetterJob struct {
	Name     string `json:"Name"`     //任务名  （发表者的UUID）
	WallType string `json:"WallType"` //墙类型
	LetterID string `json:"LetterID"` //信件id 发表时间
	UUID1    string `json:"Uuid1"`    //发表者uuid
	UUID2    string `json:"Uuid2"`    //默认为0
	UserName string `json:"UserName"` //写信人的名字
	Content  string `json:"Content"`  //信件内容
	Label    string `json:"Label"`    //邮箱签收标识
}

//删除好友
type DelFriend struct {
	Name  string `json:"Name"`
	UUID1 string `json:"Uuid1"`
	UUID2 string `json:"Uuid2"`
}

//修改好友数量
type UpdateNumFriJob struct {
	Name   string `json:"Name"`   //任务名   发表者UUID
	UUID1  string `json:"Uuid1"`  //发表者uuid
	NumFri int    `json:"NumFri"` //好友数量
}

//修改发表次数
type UpdateNumPublishJob struct {
	Name       string `json:"Name"`       //任务名  发表者UUID
	UUID1      string `json:"Uuid1"`      //发表者uuid
	NumPublish int    `json:"NumPublish"` //发表次数
}

// HTTP接口应答
type Response struct {
	Errno int         `json:"Errno"`
	Msg   string      `json:"Msg"`
	Data  interface{} `json:"Data"`
}

// 应答方法
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	// 1, 定义一个response
	var (
		response Response
	)

	response.Errno = errno
	response.Msg = msg
	response.Data = data

	// 2, 序列化json
	resp, err = json.Marshal(response)
	return
}
