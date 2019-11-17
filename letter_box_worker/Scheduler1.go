package letter_box_worker

import (
	"aaa/letter_box_worker/common"
)

// 任务调度
type Scheduler1 struct {
	jobEventChan1 chan *common.LetterJob //  etcd任务事件队列
	jobEventChan2 chan *common.DelFriend
	jobEventChan3 chan *common.UpdateNumFriJob
	jobEventChan4 chan *common.UpdateNumPublishJob
}

var (
	G_scheduler1 *Scheduler1
)

// 调度协程
func (scheduler *Scheduler1) scheduleLoop() {
	var (
		jobEvent1 *common.LetterJob
		jobEvent2 *common.DelFriend
		jobEvent3 *common.UpdateNumFriJob
		jobEvent4 *common.UpdateNumPublishJob
	)

	// 定时任务common.Job
	for {
		select {
		case jobEvent1 = <-scheduler.jobEventChan1: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteLetterJob(jobEvent1)
		case jobEvent2 = <-scheduler.jobEventChan2: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteDelFriend(jobEvent2)
		case jobEvent3 = <-scheduler.jobEventChan3: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumFriJob(jobEvent3)
		case jobEvent4 = <-scheduler.jobEventChan4: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumPublishJob(jobEvent4)
		}

	}
}

// 推送任务变化事件
func (scheduler *Scheduler1) PushJobEvent1(jobEvent *common.LetterJob) {
	scheduler.jobEventChan1 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent2(jobEvent *common.DelFriend) {
	scheduler.jobEventChan2 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent3(jobEvent *common.UpdateNumFriJob) {
	scheduler.jobEventChan3 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent4(jobEvent *common.UpdateNumPublishJob) {
	scheduler.jobEventChan4 <- jobEvent
}

// 初始化调度器
func InitScheduler1() (err error) {
	G_scheduler1 = &Scheduler1{
		jobEventChan1: make(chan *common.LetterJob, 1000),
		jobEventChan2: make(chan *common.DelFriend, 1000),
		jobEventChan3: make(chan *common.UpdateNumFriJob, 1000),
		jobEventChan4: make(chan *common.UpdateNumPublishJob, 1000),
	}
	// 启动调度协程
	go G_scheduler1.scheduleLoop()
	return
}
