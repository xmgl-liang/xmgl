package wall_worker

import (
	"aaa/wall_worker/common"
	"fmt"
)

// 任务调度
type Scheduler2 struct {
	jobResultChan1 chan *common.JobExecuteResult1 // 任务结果队列
	jobResultChan2 chan *common.JobExecuteResult2
	jobResultChan3 chan *common.JobExecuteResult3
	jobResultChan4 chan *common.JobExecuteResult4
	jobResultChan5 chan *common.JobExecuteResult5
	jobResultChan6 chan *common.JobExecuteResult6
	jobResultChan7 chan *common.JobExecuteResult7
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
			JobName:   result.WallLetterJob.Name,
			TaskPath:  common.JOB_SAVE_WALL_LETTER,
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

	fmt.Println("任务执行完成:", result.WallLetterJob.Name, common.JOB_SAVE_WALL_LETTER, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult2(result *common.JobExecuteResult2) {
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

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult3(result *common.JobExecuteResult3) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.UpdateNumChargeJob.Name,
			TaskPath:  common.JOB_UPDATE_NUMCHARGE,
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

	fmt.Println("任务执行完成:", result.UpdateNumChargeJob.Name, common.JOB_UPDATE_NUMCHARGE, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult4(result *common.JobExecuteResult4) {
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
func (scheduler *Scheduler2) handleJobResult5(result *common.JobExecuteResult5) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.WallFriend.Name,
			TaskPath:  common.JOB_WALL_FRIEND,
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

	fmt.Println("任务执行完成:", result.WallFriend.Name, common.JOB_WALL_FRIEND, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult6(result *common.JobExecuteResult6) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.LetterByFri.Name,
			TaskPath:  common.JOB_LETTER_BYFRI,
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

	fmt.Println("任务执行完成:", result.LetterByFri.Name, common.JOB_LETTER_BYFRI, result.Err)
}

// 处理任务结果
func (scheduler *Scheduler2) handleJobResult7(result *common.JobExecuteResult7) {
	var (
		jobLog *common.JobLog
	)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &common.JobLog{
			JobName:   result.KillKey,
			TaskPath:  common.JOB_KILLER_LETTER,
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

	fmt.Println("任务执行完成:", result.KillKey, common.JOB_KILLER_LETTER, result.Err)
}

// 调度协程
func (scheduler *Scheduler2) scheduleLoop() {
	var (
		jobResult1 *common.JobExecuteResult1
		jobResult2 *common.JobExecuteResult2
		jobResult3 *common.JobExecuteResult3
		jobResult4 *common.JobExecuteResult4
		jobResult5 *common.JobExecuteResult5
		jobResult6 *common.JobExecuteResult6
		jobResult7 *common.JobExecuteResult7
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
		case jobResult5 = <-scheduler.jobResultChan5: // 监听任务执行结果
			scheduler.handleJobResult5(jobResult5)
		case jobResult6 = <-scheduler.jobResultChan6: // 监听任务执行结果
			scheduler.handleJobResult6(jobResult6)
		case jobResult7 = <-scheduler.jobResultChan7: // 监听任务执行结果
			scheduler.handleJobResult7(jobResult7)
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

func (scheduler *Scheduler2) PushJobResult5(jobResult *common.JobExecuteResult5) {
	scheduler.jobResultChan5 <- jobResult
}

func (scheduler *Scheduler2) PushJobResult6(jobResult *common.JobExecuteResult6) {
	scheduler.jobResultChan6 <- jobResult
}

func (scheduler *Scheduler2) PushJobResult7(jobResult *common.JobExecuteResult7) {
	scheduler.jobResultChan7 <- jobResult
}

// 初始化调度器
func InitScheduler2() (err error) {
	G_scheduler2 = &Scheduler2{
		jobResultChan1: make(chan *common.JobExecuteResult1, 1000),
		jobResultChan2: make(chan *common.JobExecuteResult2, 1000),
		jobResultChan3: make(chan *common.JobExecuteResult3, 1000),
		jobResultChan4: make(chan *common.JobExecuteResult4, 1000),
		jobResultChan5: make(chan *common.JobExecuteResult5, 1000),
		jobResultChan6: make(chan *common.JobExecuteResult6, 1000),
		jobResultChan7: make(chan *common.JobExecuteResult7, 1000),
	}
	// 启动调度协程
	go G_scheduler2.scheduleLoop()
	return
}
