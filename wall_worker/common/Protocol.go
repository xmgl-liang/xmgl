package common

import (
	"encoding/json"
	"time"
	"strings"
)

///////////////////////////WALL///////////////////////////////

// 定时任务
//墙任务
type WallLetterJob struct {
	Name     string `json:"Name"`     //任务名  （发表者的UUID）
	LetterID string `json:"LetterID"` //信件id 发表时间
	UUID1    string `json:"Uuid1"`    //发表者uuid
	UUID2    string `json:"Uuid2"`    //默认为0
	UserName string `json:"UserName"` //写信人的名字
	Content  string `json:"Content"`  //信件内容
	Label    string `json:"Label"`    //邮箱签收标识
	WallType string `json:"WallType"` //墙类型
}

//修改发表次数
type UpdateNumPublishJob struct {
	Name       string `json:"Name"`       //任务名  发表者UUID
	UUID1      string `json:"Uuid1"`      //发表者uuid
	NumPublish int    `json:"NumPublish"` //发表次数
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

//墙友建立
type WallFriend struct {
	Name  string `json:"Name"`
	UUID1 string `json:"Uuid1"` //使用者
	UUID2 string `json:"Uuid2"` //好友
}

//信件建立
type LetterByFri struct {
	Name     string `json:"Name"`      //任务名  （接收者的UUID）
	WallType string `json:"WallType"`  //墙类型
	LetterID string `json:"LettterID"` //信件id 发表时间
	UUID1    string `json:"Uuid1"`     //发表者uuid
	UUID2    string `json:"Uuid2"`     //接收者uuid
	UserName string `json:"UserName"`  //写信人的名字
	Content  string `json:"Content"`   //信件内容
	Label    string `json:"Label"`     //邮箱签收标识
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

//WallLetterJob
// 反序列化Job
func UnpackWallLetterJob(value []byte) (ret *WallLetterJob, err error) {
	var (
		job *WallLetterJob
	)

	job = &WallLetterJob{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult1 struct {
	WallLetterJob *WallLetterJob // 执行任务
	TaskPath      string         //任务路径
	Err           error          // 脚本错误原因
	StartTime     time.Time      // 启动时间
	EndTime       time.Time      // 结束时间
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
type JobExecuteResult2 struct {
	UpdateNumPublishJob *UpdateNumPublishJob // 执行任务
	TaskPath            string               //任务路径
	Err                 error                // 脚本错误原因
	StartTime           time.Time            // 启动时间
	EndTime             time.Time            // 结束时间
}

//UpdateNumChargeJob
// 反序列化Job
func UnpackUpdateNumChargeJob(value []byte) (ret *UpdateNumChargeJob, err error) {
	var (
		job *UpdateNumChargeJob
	)

	job = &UpdateNumChargeJob{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult3 struct {
	UpdateNumChargeJob *UpdateNumChargeJob // 执行任务
	TaskPath           string              //任务路径
	Err                error               // 脚本错误原因
	StartTime          time.Time           // 启动时间
	EndTime            time.Time           // 结束时间
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
type JobExecuteResult4 struct {
	UpdateNumFriJob *UpdateNumFriJob // 执行任务
	TaskPath        string           //任务路径
	Err             error            // 脚本错误原因
	StartTime       time.Time        // 启动时间
	EndTime         time.Time        // 结束时间
}

//WallFriend
// 反序列化Job
func UnpackWallFriend(value []byte) (ret *WallFriend, err error) {
	var (
		job *WallFriend
	)

	job = &WallFriend{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult5 struct {
	WallFriend *WallFriend // 执行任务
	TaskPath   string      //任务路径
	Err        error       // 脚本错误原因
	StartTime  time.Time   // 启动时间
	EndTime    time.Time   // 结束时间
}

//LetterByFri
// 反序列化Job
func UnpackLetterByFri(value []byte) (ret *LetterByFri, err error) {
	var (
		job *LetterByFri
	)

	job = &LetterByFri{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult6 struct {
	LetterByFri *LetterByFri // 执行任务
	TaskPath    string       //任务路径
	Err         error        // 脚本错误原因
	StartTime   time.Time    // 启动时间
	EndTime     time.Time    // 结束时间
}

type JobExecuteResult7 struct {
	KillKey 	string // 执行任务
	TaskPath    string       //任务路径
	Err         error        // 脚本错误原因
	StartTime   time.Time    // 启动时间
	EndTime     time.Time    // 结束时间
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

// 从 "/ma/wall/kill/ job提取job10
func ExtractKillerName(killerKey string) string {
	return strings.TrimPrefix(killerKey, JOB_KILLER_LETTER)
}
