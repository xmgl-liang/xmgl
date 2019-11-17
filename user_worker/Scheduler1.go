package user_worker

import (
	"aaa/user_worker/common"
)

// 任务调度
type Scheduler1 struct {
	jobEventChan1 chan *common.UserMsg //  etcd任务事件队列
}

var (
	G_scheduler1 *Scheduler1
)

// 调度协程
func (scheduler *Scheduler1) scheduleLoop() {
	var (
		jobEvent1 *common.UserMsg
	)

	// 定时任务common.Job
	for {
		select {
		case jobEvent1 = <-scheduler.jobEventChan1: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUserMsg(jobEvent1)
		}

	}
}

// 推送任务变化事件
func (scheduler *Scheduler1) PushJobEvent1(jobEvent *common.UserMsg) {
	scheduler.jobEventChan1 <- jobEvent
}

// 初始化调度器
func InitScheduler1() (err error) {
	G_scheduler1 = &Scheduler1{
		jobEventChan1: make(chan *common.UserMsg, 1000),
	}
	// 启动调度协程
	go G_scheduler1.scheduleLoop()
	return
}
