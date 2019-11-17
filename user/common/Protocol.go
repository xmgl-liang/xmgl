package common

import (
	"encoding/json"
)

///////////////////////USER///////////////////////////////

//用户信息
type UserMsg struct {
	Name     string `json:"Name"`     //任务名
	UserUUID string `json:"Uuid1"`    //uuid
	UserName string `json:"UserName"` //用户名
	Sex      string `json:"Sex"`      //用户的性别
	Text     string `json:"Text"`     //座右铭
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
