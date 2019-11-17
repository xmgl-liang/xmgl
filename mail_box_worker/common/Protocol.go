package common

import (
	"encoding/json"
	"time"
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

//MailFriend
// 反序列化Job
func UnpackMailFriend(value []byte) (ret *MailFriend, err error) {
	var (
		job *MailFriend
	)

	job = &MailFriend{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult1 struct {
	MailFriend *MailFriend // 执行任务
	Output     []byte      // 脚本输出
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
type JobExecuteResult2 struct {
	LetterByFri *LetterByFri // 执行任务
	Output      []byte       // 脚本输出
	Err         error        // 脚本错误原因
	StartTime   time.Time    // 启动时间
	EndTime     time.Time    // 结束时间
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
	Output             []byte              // 脚本输出
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
	Output          []byte           // 脚本输出
	Err             error            // 脚本错误原因
	StartTime       time.Time        // 启动时间
	EndTime         time.Time        // 结束时间
}

//NoLetter
// 反序列化Job
func UnpackNoLetter(value []byte) (ret *NoLetter, err error) {
	var (
		job *NoLetter
	)

	job = &NoLetter{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务执行结果
type JobExecuteResult5 struct {
	NoLetter  *NoLetter // 执行任务
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
