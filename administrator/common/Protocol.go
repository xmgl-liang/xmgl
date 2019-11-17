package common

import (
	"encoding/json"
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

//信件建立
type Letter struct {
	Name     string `json:"Name"`      //任务名  （接收者的UUID）
	WallType string `json:"WallType"`  //墙类型
	LetterID string `json:"LetterID"` //信件id 发表时间
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

// 任务执行日志
type JobLog struct {
	JobName   string `json:"JobName" bson:"JobName"`     // 任务名字
	TaskPath  string `json:"TaskPath" bson:"TaskPath"`   // 执行路径
	Err       string `json:"Err" bson:"Err"`             // 错误原因
	StartTime int64  `json:"StartTime" bson:"StartTime"` // 任务执行开始时间
	EndTime   int64  `json:"EndTime" bson:"EndTime"`     // 任务执行结束时间
}

// 任务日志过滤条件
type JobLogFilter struct {
	TaskPath string `bson:"TaskPath"`
}

// 任务日志排序规则
type SortLogByStartTime struct {
	SortOrder int `bson:"StartTime"` // {startTime: -1}
}

//workers ip struct
type WorkersIP struct {
	WorkerIP string 
}

// 提取worker的IP
func ExtractWorkerIP(regKey string, workerDir string) string {
	return strings.TrimPrefix(regKey, workerDir)
}
