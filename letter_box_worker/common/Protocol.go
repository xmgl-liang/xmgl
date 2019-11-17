package common

import (
	"encoding/json"
	"time"
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

//LetterJob
// 反序列化Job
func UnpackLetterJob(value []byte) (ret *LetterJob, err error) {
	var (
		job *LetterJob
	)

	job = &LetterJob{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult1 struct {
	LetterJob *LetterJob // 执行任务
	TaskPath  string     //任务路径
	Err       error      // 脚本错误原因
	StartTime time.Time  // 启动时间
	EndTime   time.Time  // 结束时间
}

//DelFriend
// 反序列化Job
func UnpackDelFriend(value []byte) (ret *DelFriend, err error) {
	var (
		job *DelFriend
	)

	job = &DelFriend{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult2 struct {
	DelFriend *DelFriend // 执行任务
	TaskPath  string     //任务路径
	Err       error      // 脚本错误原因
	StartTime time.Time  // 启动时间
	EndTime   time.Time  // 结束时间
}

//UpdateNumFriJob
// 反序列化Job
func UnpackUpdateNumFriJob(value []byte) (ret *UpdateNumFriJob, err error) {
	var (
		job *UpdateNumFriJob
	)

	job = &UpdateNumFriJob{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult3 struct {
	UpdateNumFriJob *UpdateNumFriJob // 执行任务
	TaskPath        string           //任务路径
	Err             error            // 脚本错误原因
	StartTime       time.Time        // 启动时间
	EndTime         time.Time        // 结束时间
}

//UpdateNumPublishJob
// 反序列化Job
func UnpackUpdateNumPublishJob(value []byte) (ret *UpdateNumPublishJob, err error) {
	var (
		job *UpdateNumPublishJob
	)

	job = &UpdateNumPublishJob{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult4 struct {
	UpdateNumPublishJob *UpdateNumPublishJob // 执行任务
	TaskPath            string               //任务路径
	Err                 error                // 脚本错误原因
	StartTime           time.Time            // 启动时间
	EndTime             time.Time            // 结束时间
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
