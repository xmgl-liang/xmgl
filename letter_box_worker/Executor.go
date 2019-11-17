package letter_box_worker

import (
	"aaa/letter_box_worker/common"
	"aaa/letter_box_worker/tasks"
	"math/rand"
	"time"
)

// 任务执行器
type Executor struct {
}

var (
	G_executor *Executor
)

// 执行一个任务
func (executor *Executor) ExecuteLetterJob(info *common.LetterJob) {
	go func() {
		var (
			err     error
			result  *common.JobExecuteResult1
			jobLock *JobLock
		)

		// 任务结果
		result = &common.JobExecuteResult1{
			LetterJob: info,
		}

		// 初始化分布式锁  Name需要修饰一下加点其他标记
		jobLock = G_jobMgr.CreateJobLock(info.Name)

		// 记录任务开始时间
		result.StartTime = time.Now()

		// 上锁
		// 随机睡眠(0~1s)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		err = jobLock.TryLock(common.LOCK_1)
		defer jobLock.Unlock()

		if err != nil { // 上锁失败
			result.Err = err
			result.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务启动时间
			result.StartTime = time.Now()

			// 执行任务并捕获输出
			err = tasks.AddLetterJob(info)

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Err = err
		}
		// 任务执行完成后，把执行的结果返回给Scheduler2
		G_scheduler2.PushJobResult1(result)
	}()
}

func (executor *Executor) ExecuteDelFriend(info *common.DelFriend) {
	go func() {
		var (
			err     error
			result  *common.JobExecuteResult2
			jobLock *JobLock
		)

		// 任务结果
		result = &common.JobExecuteResult2{
			DelFriend: info,
		}

		// 初始化分布式锁  Name需要修饰一下加点其他标记
		jobLock = G_jobMgr.CreateJobLock(info.Name)

		// 记录任务开始时间
		result.StartTime = time.Now()

		// 上锁
		// 随机睡眠(0~1s)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		err = jobLock.TryLock(common.LOCK_2)
		defer jobLock.Unlock()

		if err != nil { // 上锁失败
			result.Err = err
			result.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务启动时间
			result.StartTime = time.Now()

			// 执行任务并捕获输出
			err = tasks.DeleteDelFriend(info)

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Err = err
		}
		// 任务执行完成后，把执行的结果返回给Scheduler2
		G_scheduler2.PushJobResult2(result)
	}()
}

func (executor *Executor) ExecuteUpdateNumFriJob(info *common.UpdateNumFriJob) {
	go func() {
		var (
			err     error
			result  *common.JobExecuteResult3
			jobLock *JobLock
		)

		// 任务结果
		result = &common.JobExecuteResult3{
			UpdateNumFriJob: info,
		}

		// 初始化分布式锁  Name需要修饰一下加点其他标记
		jobLock = G_jobMgr.CreateJobLock(info.Name)

		// 记录任务开始时间
		result.StartTime = time.Now()

		// 上锁
		// 随机睡眠(0~1s)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		err = jobLock.TryLock(common.LOCK_3)
		defer jobLock.Unlock()

		if err != nil { // 上锁失败
			result.Err = err
			result.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务启动时间
			result.StartTime = time.Now()

			// 执行任务并捕获输出
			err = tasks.UpdateNumFri(info)

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Err = err
		}
		// 任务执行完成后，把执行的结果返回给Scheduler2
		G_scheduler2.PushJobResult3(result)
	}()
}

func (executor *Executor) ExecuteUpdateNumPublishJob(info *common.UpdateNumPublishJob) {
	go func() {
		var (
			err     error
			result  *common.JobExecuteResult4
			jobLock *JobLock
		)

		// 任务结果
		result = &common.JobExecuteResult4{
			UpdateNumPublishJob: info,
		}

		// 初始化分布式锁  Name需要修饰一下加点其他标记
		jobLock = G_jobMgr.CreateJobLock(info.Name)

		// 记录任务开始时间
		result.StartTime = time.Now()

		// 上锁
		// 随机睡眠(0~1s)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		err = jobLock.TryLock(common.LOCK_4)
		defer jobLock.Unlock()

		if err != nil { // 上锁失败
			result.Err = err
			result.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务启动时间
			result.StartTime = time.Now()

			// 执行任务并捕获输出
			err = tasks.UpdateNumPublish(info)

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Err = err
		}
		// 任务执行完成后，把执行的结果返回给Scheduler2
		G_scheduler2.PushJobResult4(result)
	}()
}

//  初始化执行器
func InitExecutor() (err error) {
	G_executor = &Executor{}
	return
}
