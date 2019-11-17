package common

import (
	"encoding/json"
	"time"
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

//UserMsg
// 反序列化Job
func UnpackLetterByFri(value []byte) (ret *UserMsg, err error) {
	var (
		job *UserMsg
	)

	job = &UserMsg{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult1 struct {
	UserMsg   *UserMsg  // 执行任务
	Output    []byte    // 脚本输出
	Err       error     // 脚本错误原因
	StartTime time.Time // 启动时间
	EndTime   time.Time // 结束时间
}

// 任务执行日志
type JobLog struct {
	JobName   string `json:"JobName" bson:"JobName"`     // 任务名字
	TaskPath  string `json:"TaskPath" bson:"TaskPath"`   // 执行路径
	Err       string `json:"Err" bson:"Err"`             // 错误原因
	StartTime int64  `json:"StartTime" bson:"StartTime"` // 任务执行开始时间
	EndTime   int64  `json:"EndTime" bson:"EndTime"`     // 任务执行结束时间
}

// 日志批次
type LogBatch struct {
	Logs []interface{} // 多条日志
}
