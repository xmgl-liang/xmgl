package common

import (
	"encoding/json"
)

//////////////////mail_box////////////////////

//好友建立
type MailFriend struct {
	Name  string `json:"Name"`
	UUID1 string `json:"Uuid1"`
	UUID2 string `json:"Uuid2"`
}

//信件建立
type LetterByFri struct {
	Name     string `json:"Name"`     //任务名  （接收者的UUID）
	WallType string `json:"WallType"` //墙类型
	LetterID string `json:"LetterID"` //信件id 发表时间
	UUID1    string `json:"Uuid1"`    //发表者uuid
	UUID2    string `json:"Uuid2"`    //接收者uuid
	UserName string `json:"UserName"` //写信人的名字
	Content  string `json:"Content"`  //信件内容
	Label    string `json:"Label"`    //邮箱签收标识
}

//修改接受次数
type UpdateNumChargeJob struct {
	Name      string `json:"Name"`      //任务名  发表者UUID
	UUID1     string `json:"Uuid1"`     //发表者uuid
	NumCharge int    `json:"NumCharge"` //发表次数
}

//修改好友数量
type UpdateNumFriJob struct {
	Name   string `json:"Name"`   //任务名   发表者UUID
	UUID1  string `json:"Uuid1"`  //发表者uuid
	NumFri int    `json:"NumFri"` //好友数量
}

//信件拒绝
type NoLetter struct {
	Name     string `json:"Name"`     //任务名   拒绝者UUID
	UUID1    string `json:"Uuid1"`    //发表者
	UUID2    string `json:"Uuid2"`    //拒绝者
	LetterID string `json:"LetterID"` //信件id 时间
	Label    string `json:"Label"`    //信件接收标识
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
