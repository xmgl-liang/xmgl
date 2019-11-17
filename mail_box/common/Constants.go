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

	// 服务注册目录
	JOB_WORKER_DIR = "/ma/mail_box/workers/"
)
