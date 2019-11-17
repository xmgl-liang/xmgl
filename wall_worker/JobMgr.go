package wall_worker

import (
	"aaa/wall_worker/common"
	"aaa/wall_worker/configs"
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

// 任务管理器
type JobMgr struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	watcher clientv3.Watcher
}

var (
	// 单例
	G_jobMgr *JobMgr
)

// 初始化管理器
func InitJobMgr() (err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		watcher clientv3.Watcher
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints:   configs.G_config.EtcdEndpoints,                                     // 集群地址
		DialTimeout: time.Duration(configs.G_config.EtcdDialTimeout) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)

	// 赋值单例
	G_jobMgr = &JobMgr{
		client:  client,
		kv:      kv,
		lease:   lease,
		watcher: watcher,
	}

	// 启动任务监听
	G_jobMgr.watchJobWallLetter()
	G_jobMgr.watchNumPublish()
	G_jobMgr.watchNumCharge()
	G_jobMgr.watchNumFri()
	G_jobMgr.watchWallFriend()
	G_jobMgr.watchLetterByFri()
	G_jobMgr.watchKiller()

	return
}

// 创建任务执行锁
func (jobMgr *JobMgr) CreateJobLock(jobName string) (jobLock *JobLock) {

	jobLock = InitJobLock(jobName, jobMgr.kv, jobMgr.lease)
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchJobWallLetter() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.WallLetterJob
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_SAVE_WALL_LETTER/目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_SAVE_WALL_LETTER, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackWallLetterJob(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent1(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_SAVE_WALL_LETTER目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_SAVE_WALL_LETTER, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackWallLetterJob(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent1(job)
			}
		}
	}()
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchNumPublish() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.UpdateNumPublishJob
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_UPDATE_NUMPUBLISH目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_UPDATE_NUMPUBLISH, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackUpdateNumPublishJob(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent2(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_UPDATE_NUMPUBLISH目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_UPDATE_NUMPUBLISH, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackUpdateNumPublishJob(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent2(job)
			}
		}
	}()
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchNumCharge() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.UpdateNumChargeJob
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_UPDATE_NUMCHARGE目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_UPDATE_NUMCHARGE, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackUpdateNumChargeJob(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent3(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_UPDATE_NUMCHARGE目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_UPDATE_NUMCHARGE, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackUpdateNumChargeJob(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent3(job)
			}
		}
	}()
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchNumFri() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.UpdateNumFriJob
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_UPDATE_NUMFRI目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_UPDATE_NUMFRI, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackUpdateNumFriJob(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent4(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_UPDATE_NUMFRI目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_UPDATE_NUMFRI, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackUpdateNumFriJob(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent4(job)
			}
		}
	}()
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchWallFriend() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.WallFriend
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_WALL_FRIEND目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_WALL_FRIEND, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackWallFriend(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent5(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_WALL_FRIEND目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_WALL_FRIEND, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackWallFriend(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent5(job)
			}
		}
	}()
	return
}

// 监听任务变化
func (jobMgr *JobMgr) watchLetterByFri() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.LetterByFri
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
	)

	// 1, get一下JOB_LETTER_BYFRI目录下的所有任务，并且获知当前集群的revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_LETTER_BYFRI, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.UnpackLetterByFri(kvpair.Value); err == nil {
			// 同步给scheduler(调度协程)
			G_scheduler1.PushJobEvent6(job)
		}
	}

	// 2, 从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听JOB_LETTER_BYFRI目录的后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_LETTER_BYFRI, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.UnpackLetterByFri(watchEvent.Kv.Value); err != nil {
						continue
					}
				case mvccpb.DELETE: // 任务被删除了
				}
				// 变化推给scheduler
				G_scheduler1.PushJobEvent6(job)
			}
		}
	}()
	return
}

func (jobMgr *JobMgr) watchKiller() {
	var (
		watchChan clientv3.WatchChan
		watchResp clientv3.WatchResponse
		watchEvent *clientv3.Event
		jobName string
	)
	// 监听/cron/killer目录
	go func() { // 监听协程
		// 监听/cron/killer/目录的变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_KILLER_LETTER, clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: 
				case mvccpb.DELETE: // killer标记过期, 被自动删除
					jobName = common.ExtractKillerName(string(watchEvent.Kv.Key))
					// 事件推给scheduler
					G_scheduler1.PushJobEvent7(jobName)
				}
			}
		}
	}()
}

func (jobMgr *JobMgr) DeleteJob(name string) (err error) {
	var (
		jobKey    string
	)

	// etcd中保存任务的key
	jobKey = common.JOB_SAVE_WALL_LETTER + name

	// 从etcd中删除它
	if _, err = jobMgr.kv.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return
	}

	return
}