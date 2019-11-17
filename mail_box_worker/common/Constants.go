package common

const (
	//邮箱建立好友关系
	JOB_FRIEND_MAIL = "/ma/mail_box/friend/"

	//邮箱建立信件关系（历时）
	JOB_LETTER_MAIL = "/ma/mail_box/letter/"

	//用户接收次数减1
	JOB_UPDATE_NUMCHARGE = "/ma/mail_box/charge/"

	//好友数量减1
	JOB_UPDATE_NUMFRI = "/ma/mail_box/numFri/"

	//拒绝访问信件
	JOB_NO_LETTER_MAIL = "/ma/mail_box/refuse/"

	//锁路径
	JOB_LOCK_DIR = "/ma/mail_box/lock/"

	//锁任务类型
	LOCK_1 = "mailFriend/"

	LOCK_2 = "letterByFri/"

	LOCK_3 = "updateNumChargeJob/"

	LOCK_4 = "updateNumFriJob/"

	LOCK_5 = "noLetter/"

	// 服务注册目录
	JOB_WORKER_DIR = "/ma/mail_box_worker/workers/"
)
