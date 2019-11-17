package wall_worker

import (
	"aaa/wall_worker/common"
)

// 任务调度
type Scheduler1 struct {
	jobEventChan1 chan *common.WallLetterJob //  etcd任务事件队列
	jobEventChan2 chan *common.UpdateNumPublishJob
	jobEventChan3 chan *common.UpdateNumChargeJob
	jobEventChan4 chan *common.UpdateNumFriJob
	jobEventChan5 chan *common.WallFriend
	jobEventChan6 chan *common.LetterByFri
	jobEventChan7 chan string
}

var (
	G_scheduler1 *Scheduler1
)

// 调度协程
func (scheduler *Scheduler1) scheduleLoop() {
	var (
		jobEvent1 *common.WallLetterJob
		jobEvent2 *common.UpdateNumPublishJob
		jobEvent3 *common.UpdateNumChargeJob
		jobEvent4 *common.UpdateNumFriJob
		jobEvent5 *common.WallFriend
		jobEvent6 *common.LetterByFri
		jobEvent7 string
	)

	// 定时任务common.Job
	for {
		select {
		case jobEvent1 = <-scheduler.jobEventChan1: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteWallLetterJob(jobEvent1)
		case jobEvent2 = <-scheduler.jobEventChan2: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumPublishJob(jobEvent2)
		case jobEvent3 = <-scheduler.jobEventChan3: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumChargeJob(jobEvent3)
		case jobEvent4 = <-scheduler.jobEventChan4: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumFriJob(jobEvent4)
		case jobEvent5 = <-scheduler.jobEventChan5: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteWallFriend(jobEvent5)
		case jobEvent6 = <-scheduler.jobEventChan6: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteLetterByFri(jobEvent6)
		case jobEvent7 = <-scheduler.jobEventChan7: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteKill(jobEvent7)
		}

	}
}

// 推送任务变化事件
func (scheduler *Scheduler1) PushJobEvent1(jobEvent *common.WallLetterJob) {
	scheduler.jobEventChan1 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent2(jobEvent *common.UpdateNumPublishJob) {
	scheduler.jobEventChan2 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent3(jobEvent *common.UpdateNumChargeJob) {
	scheduler.jobEventChan3 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent4(jobEvent *common.UpdateNumFriJob) {
	scheduler.jobEventChan4 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent5(jobEvent *common.WallFriend) {
	scheduler.jobEventChan5 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent6(jobEvent *common.LetterByFri) {
	scheduler.jobEventChan6 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent7(jobEvent string) {
	scheduler.jobEventChan7 <- jobEvent
}

// 初始化调度器
func InitScheduler1() (err error) {
	G_scheduler1 = &Scheduler1{
		jobEventChan1: make(chan *common.WallLetterJob, 1000),
		jobEventChan2: make(chan *common.UpdateNumPublishJob, 1000),
		jobEventChan3: make(chan *common.UpdateNumChargeJob, 1000),
		jobEventChan4: make(chan *common.UpdateNumFriJob, 1000),
		jobEventChan5: make(chan *common.WallFriend, 1000),
		jobEventChan6: make(chan *common.LetterByFri, 1000),
		jobEventChan7: make(chan string, 1000),
	}
	// 启动调度协程
	go G_scheduler1.scheduleLoop()
	return
}
