package user_worker

import (
	"aaa/user_worker/common"
	"aaa/user_worker/tasks"
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
func (executor *Executor) ExecuteUserMsg(info *common.UserMsg) {
	go func() {
		var (
			err     error
			output  []byte
			result  *common.JobExecuteResult1
			jobLock *JobLock
		)

		// 任务结果
		result = &common.JobExecuteResult1{
			UserMsg: info,
			Output:  make([]byte, 0),
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
			err = tasks.UpdateUserMsg(info)

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Output = output
			result.Err = err
		}
		// 任务执行完成后，把执行的结果返回给Scheduler，Scheduler会从executingTable中删除掉执行记录
		G_scheduler2.PushJobResult1(result)
	}()
}

//  初始化执行器
func InitExecutor() (err error) {
	G_executor = &Executor{}
	return
}
