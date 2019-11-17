package common

const (
	//信箱回信
	JOB_SAVE_REPLY_LETTER = "/ma/letter_box/letter/"

	//信箱删除好友
	JOB_DEL_FRIEND = "/ma/letter_box/del/"

	//信箱好友数加1
	JOB_UPDATE_NUMFRI = "/ma/letter_box/numFri/"

	//信箱用户发表次数减1
	JOB_UPDATE_NUMPUBLISH = "/ma/letter_box/publish/"

	//锁路径
	JOB_LOCK_DIR = "/ma/letter_box/lock/"

	//锁任务类型
	LOCK_1 = "letterJob/"

	LOCK_2 = "delFriend/"

	LOCK_3 = "updateNumFriJob/"

	LOCK_4 = "updateNumPublishJob/"

	// 服务注册目录
	JOB_WORKER_DIR = "/ma/letter_box_worker/workers/"
)
