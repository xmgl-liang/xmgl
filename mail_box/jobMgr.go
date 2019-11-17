package mail_box

import (
	"context"
	"encoding/json"
	"time"

	"aaa/mail_box/common"
	"aaa/mail_box/configs"

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

/////////////////////////mail_box////////////////////////////
// 建立好友关系任务
func (jobMgr *JobMgr) SaveMailFriend(job *common.MailFriend) (err error) {
	// 把任务保存到JOB_FRIEND_MAIL = "/ma/mail_box/friend/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_FRIEND_MAIL + job.Name
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

//0拒绝  1可接受  -1即刻   2接受
// 建立信件关系任务 SaveFriLetter
func (jobMgr *JobMgr) SaveLetterByFri(job *common.LetterByFri) (err error) {
	// 把任务保存到JOB_LETTER_MAIL = "/ma/mail_box/letter/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_LETTER_MAIL + job.Name
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

// 建立接收次数减1任务  user
func (jobMgr *JobMgr) SaveNumCharge(job *common.UpdateNumChargeJob) (err error) {
	// 把任务保存到JOB_UPDATE_NUMCHARGE = "/ma/mail_box/charge/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_UPDATE_NUMCHARGE + job.Name
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

// 建立好友数量减1任务  user
func (jobMgr *JobMgr) SaveNumFri(job *common.UpdateNumFriJob) (err error) {
	// 把任务保存到JOB_UPDATE_NUMFRI = "/ma/mail_box/numFri/任务名" -> json
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

//0拒绝  1可接受  -1即刻   2接受
// 拒绝访问信件
func (jobMgr *JobMgr) SaveNoLetter(job *common.NoLetter) (err error) {
	// 把任务保存到JOB_NO_LETTER_MAIL = "/ma/mail_box/refuse/任务名" -> json
	var (
		jobKey   string
		jobValue []byte
	)

	//etcd是会关注信件记录的过程，所以都能执行到位

	// etcd的保存key
	jobKey = common.JOB_NO_LETTER_MAIL + job.Name
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
