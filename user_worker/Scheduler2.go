package user_worker

import (
	"aaa/user_worker/common"
	"fmt"
)

// 任务调度
type Scheduler2 struct {
	jobResultChan1 chan *common.JobExecuteResult1 // 任务结果队列
}

var (
	G_scheduler2 *Scheduler2
)

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult1(result *common.JobExecuteResult1) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.UserMsg.Name,
			TaskPath:  common.JOB_UPDATE_USER,
			StartTime: result.StartTime.UnixNano() / 1000 / 1000,
			EndTime:   result.EndTime.UnixNano() / 1000 / 1000,
		}
		if result.Err != nil {
			jobLog.Err = result.Err.Error()
		} else {
			jobLog.Err = ""
		}
		G_logSink.Append(jobLog)
	}

	fmt.Println("任务执行完成:", result.UserMsg.Name, common.JOB_UPDATE_USER, result.Err)
}

// 调度协程
func (scheduler *Scheduler2) scheduleLoop() {
	var (
		jobResult1 *common.JobExecuteResult1
	)

	// 定时任务common.Job
	for {
		select {
		case jobResult1 = <-scheduler.jobResultChan1: // 监听任务执行结果
			scheduler.handleJobResult1(jobResult1)
		}
	}
}

// 回传任务执行结果
func (scheduler *Scheduler2) PushJobResult1(jobResult *common.JobExecuteResult1) {
	scheduler.jobResultChan1 <- jobResult
}

// 初始化调度器
func InitScheduler2() (err error) {
	G_scheduler2 = &Scheduler2{
		jobResultChan1: make(chan *common.JobExecuteResult1, 1000),
	}
	// 启动调度协程
	go G_scheduler2.scheduleLoop()
	return
}
