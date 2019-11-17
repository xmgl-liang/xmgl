package letter_box

import (
	"aaa/letter_box/common"
	"aaa/letter_box/configs"
	"context"
	"encoding/json"
	"time"

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

/////////////////////////letter_box//////////////////////////
// 保存回信任务
func (jobMgr *JobMgr) SaveJobReplyLetter(job *common.LetterJob) (err error) {
	// 把任务保存到JOB_SAVE_REPLY_LETTER = "/ma/letter_box/letter/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	// etcd的保存key
	jobKey = common.JOB_SAVE_REPLY_LETTER + job.Name
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

// 删除好友关系任务
func (jobMgr *JobMgr) SaveDelFriend(job *common.DelFriend) (err error) {
	// 把任务保存到JOB_DEL_FRIEND = "/ma/letter_box/del/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_DEL_FRIEND + job.Name
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

// 建立好友数量加1任务  user
func (jobMgr *JobMgr) SaveAddNumFri(job *common.UpdateNumFriJob) (err error) {
	// 把任务保存到JOB_UPDATE_NUMFRI = "/ma/letter_box/numFri/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_UPDATE_NUMFRI + job.Name
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

// 建立发表次数减1任务  user
func (jobMgr *JobMgr) SaveNumPublish(job *common.UpdateNumPublishJob) (err error) {
	// 把任务保存到JOB_UPDATE_NUMPUBLISH = "/ma/letter_box/publish/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_UPDATE_NUMPUBLISH + job.Name
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
