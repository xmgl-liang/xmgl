package mail_box_worker

import (
	"aaa/mail_box_worker/common"
)

// 任务调度
type Scheduler1 struct {
	jobEventChan1 chan *common.MailFriend //  etcd任务事件队列
	jobEventChan2 chan *common.LetterByFri
	jobEventChan3 chan *common.UpdateNumChargeJob
	jobEventChan4 chan *common.UpdateNumFriJob
	jobEventChan5 chan *common.NoLetter
}

var (
	G_scheduler1 *Scheduler1
)

// 调度协程
func (scheduler *Scheduler1) scheduleLoop() {
	var (
		jobEvent1 *common.MailFriend
		jobEvent2 *common.LetterByFri
		jobEvent3 *common.UpdateNumChargeJob
		jobEvent4 *common.UpdateNumFriJob
		jobEvent5 *common.NoLetter
	)

	// 定时任务common.Job
	for {
		select {
		case jobEvent1 = <-scheduler.jobEventChan1: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteMailFriend(jobEvent1)
		case jobEvent2 = <-scheduler.jobEventChan2: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteLetterByFri(jobEvent2)
		case jobEvent3 = <-scheduler.jobEventChan3: //监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumChargeJob(jobEvent3)
		case jobEvent4 = <-scheduler.jobEventChan4: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteUpdateNumFriJob(jobEvent4)
		case jobEvent5 = <-scheduler.jobEventChan5: //监听任务变化事4
			// 对内存中维护的任务列表做增删改查
			G_executor.ExecuteNoLetter(jobEvent5)
		}

	}
}

// 推送任务变化事件
func (scheduler *Scheduler1) PushJobEvent1(jobEvent *common.MailFriend) {
	scheduler.jobEventChan1 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent2(jobEvent *common.LetterByFri) {
	scheduler.jobEventChan2 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent3(jobEvent *common.UpdateNumChargeJob) {
	scheduler.jobEventChan3 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent4(jobEvent *common.UpdateNumFriJob) {
	scheduler.jobEventChan4 <- jobEvent
}

func (scheduler *Scheduler1) PushJobEvent5(jobEvent *common.NoLetter) {
	scheduler.jobEventChan5 <- jobEvent
}

// 初始化调度器
func InitScheduler1() (err error) {
	G_scheduler1 = &Scheduler1{
		jobEventChan1: make(chan *common.MailFriend, 1000),
		jobEventChan2: make(chan *common.LetterByFri, 1000),
		jobEventChan3: make(chan *common.UpdateNumChargeJob, 1000),
		jobEventChan4: make(chan *common.UpdateNumFriJob, 1000),
		jobEventChan5: make(chan *common.NoLetter, 1000),
	}
	// 启动调度协程
	go G_scheduler1.scheduleLoop()
	return
}
