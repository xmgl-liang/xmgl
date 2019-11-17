package user

import (
	"context"
	"encoding/json"
	"time"
	"fmt"
	"aaa/user/common"
	"aaa/user/configs"

	"github.com/coreos/etcd/clientv3"
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

//////////////////////////////user//////////////////////////////////

//修改用户信息
func (jobMgr *JobMgr) SaveUserUpdata(job *common.UserMsg) (err error) {
	// 把任务保存到JOB_UPDATE_USER = "/ma/users/update/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位
	fmt.Println("2222")
	// etcd的保存key
	jobKey = common.JOB_UPDATE_USER + job.Name
	fmt.Println(jobKey)
	// 任务信息json
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}
	// 保存到etcd
	if _, err = jobMgr.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	return
}
