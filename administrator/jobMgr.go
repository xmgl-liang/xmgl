package administrator

import (
	"context"
	"encoding/json"
	"time"

	"aaa/administrator/common"
	"aaa/administrator/configs"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

//任务管理器
type JobMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var (
	//单例
	G_jobMgr *JobMgr
)

//初始化管理器
func InitJobMgr() (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)

	//初始化配置
	config = clientv3.Config{
		Endpoints:   configs.G_config.EtcdEndpoints,                                     // 集群地址
		DialTimeout: time.Duration(configs.G_config.EtcdDialTimeout) * time.Millisecond, // 连接超时
	}

	//建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	//得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	//赋值单例
	G_jobMgr = &JobMgr{
		client: client,
		kv:     kv,
		lease:  lease,
	}

	return
}

// 列举任务（查询任务）
func (jobMgr *JobMgr) ListLettersByWallType(wallType string) (jobList []*common.WallLetterJob, err error) {
	var (
		dirKey  string
		getResp *clientv3.GetResponse
		kvPair  *mvccpb.KeyValue
		job     *common.WallLetterJob
	)

	// 任务保存的目录
	dirKey = common.JOB_SAVE_WALL_LETTER + wallType + "/"

	// 获取目录下所有任务信息
	if getResp, err = jobMgr.kv.Get(context.TODO(), dirKey, clientv3.WithPrefix()); err != nil {
		return
	}

	// 初始化数组空间
	jobList = make([]*common.WallLetterJob, 0)
	// len(jobList) == 0

	// 遍历所有任务, 进行反序列化
	for _, kvPair = range getResp.Kvs {
		job = &common.WallLetterJob{}
		if err = json.Unmarshal(kvPair.Value, job); err != nil {
			err = nil
			continue
		}
		jobList = append(jobList, job)
	}
	return
}

// 杀死任务
func (jobMgr *JobMgr) KillJob(name string) (err error) {
	// 更新一下JOB_KILLER_LETTER = "/ma/wall/kill/任务名"
	var (
		killerKey      string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
	)

	// 通知worker杀死对应任务
	killerKey = common.JOB_KILLER_LETTER + name

	// 让worker监听到一次put操作, 创建一个租约让其稍后自动过期即可
	if leaseGrantResp, err = jobMgr.lease.Grant(context.TODO(), 1); err != nil {
		return
	}

	// 租约ID
	leaseId = leaseGrantResp.ID

	// 设置killer标记
	if _, err = jobMgr.kv.Put(context.TODO(), killerKey, "", clientv3.WithLease(leaseId)); err != nil {
		return
	}
	return
}
