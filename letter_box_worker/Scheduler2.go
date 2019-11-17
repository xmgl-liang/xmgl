package letter_box_worker

import (
	"aaa/letter_box_worker/common"
	"fmt"
)

// 任务调度
type Scheduler2 struct {
	jobResultChan1 chan *common.JobExecuteResult1 // 任务结果队列
	jobResultChan2 chan *common.JobExecuteResult2
	jobResultChan3 chan *common.JobExecuteResult3
	jobResultChan4 chan *common.JobExecuteResult4
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
			JobName:   result.LetterJob.Name,
			TaskPath:  common.JOB_SAVE_REPLY_LETTER,
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

	fmt.Println("任务执行完成:", result.LetterJob.Name, common.JOB_SAVE_REPLY_LETTER, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult2(result *common.JobExecuteResult2) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.DelFriend.Name,
			TaskPath:  common.JOB_DEL_FRIEND,
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

	fmt.Println("任务执行完成:", result.DelFriend.Name, common.JOB_DEL_FRIEND, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult3(result *common.JobExecuteResult3) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.UpdateNumFriJob.Name,
			TaskPath:  common.JOB_UPDATE_NUMFRI,
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

	fmt.Println("任务执行完成:", result.UpdateNumFriJob.Name, common.JOB_UPDATE_NUMFRI, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult4(result *common.JobExecuteResult4) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.UpdateNumPublishJob.Name,
			TaskPath:  common.JOB_UPDATE_NUMPUBLISH,
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

	fmt.Println("任务执行完成:", result.UpdateNumPublishJob.Name, common.JOB_UPDATE_NUMPUBLISH, result.Err)
}

// 调度协程
func (scheduler *Scheduler2) scheduleLoop() {
	var (
		jobResult1 *common.JobExecuteResult1
		jobResult2 *common.JobExecuteResult2
		jobResult3 *common.JobExecuteResult3
		jobResult4 *common.JobExecuteResult4
	)

	// 定时任务common.Job
	for {
		select {
		case jobResult1 = <-scheduler.jobResultChan1: // 监听任务执行结果
			scheduler.handleJobResult1(jobResult1)
		case jobResult2 = <-scheduler.jobResultChan2: // 监听任务执行结果
			scheduler.handleJobResult2(jobResult2)
		case jobResult3 = <-scheduler.jobResultChan3: // 监听任务执行结果
			scheduler.handleJobResult3(jobResult3)
		case jobResult4 = <-scheduler.jobResultChan4: // 监听任务执行结果
			scheduler.handleJobResult4(jobResult4)
		}
	}
}

// 回传任务执行结果
func (scheduler *Scheduler2) PushJobResult1(jobResult *common.JobExecuteResult1) {
	scheduler.jobResultChan1 <- jobResult
}

func (scheduler *Scheduler2) PushJobResult2(jobResult *common.JobExecuteResult2) {
	scheduler.jobResultChan2 <- jobResult
}

func (scheduler *Scheduler2) PushJobResult3(jobResult *common.JobExecuteResult3) {
	scheduler.jobResultChan3 <- jobResult
}

func (scheduler *Scheduler2) PushJobResult4(jobResult *common.JobExecuteResult4) {
	scheduler.jobResultChan4 <- jobResult
}

// 初始化调度器
func InitScheduler2() (err error) {
	G_scheduler2 = &Scheduler2{
		jobResultChan1: make(chan *common.JobExecuteResult1, 1000),
		jobResultChan2: make(chan *common.JobExecuteResult2, 1000),
		jobResultChan3: make(chan *common.JobExecuteResult3, 1000),
		jobResultChan4: make(chan *common.JobExecuteResult4, 1000),
	}
	// 启动调度协程
	go G_scheduler2.scheduleLoop()
	return
}
