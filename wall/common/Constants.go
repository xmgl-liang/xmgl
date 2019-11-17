package common

const (
	// 墙服务目录
	JOB_SAVE_WALL_LETTER = "/ma/wall/letter/"

	//用户发表次数减1
	JOB_UPDATE_NUMPUBLISH = "/ma/wall/publish/"

	//墙中信件过期
	JOB_KILLER_LETTER = "/ma/wall/kill/"

	//用户接收次数减1
	JOB_UPDATE_NUMCHARGE = "/ma/wall/charge/"

	//好友数量减1
	JOB_UPDATE_NUMFRI = "/ma/wall/numFri/"

	//建立好友关系
	JOB_WALL_FRIEND = "/ma/wall/friends/"

	//建立信件关系通过好友
	JOB_LETTER_BYFRI = "/ma/wall/letterByFri/"

	// 服务注册目录
	JOB_WORKER_DIR = "/ma/wall/workers/"
)
