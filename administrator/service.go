package administrator

import (
	"aaa/administrator/common"
	"aaa/administrator/dao"
	"aaa/administrator/model"
	"net/http"
	"strconv"
)

//1.查询墙中信件,从etcd中查询
// 列举所有wallLetter任务
func HandleListLettersByWallType(w http.ResponseWriter, r *http.Request) {
	var (
		postJob string
		jobList []*common.WallLetterJob
		bytes   []byte
		err     error
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("wallType")

	//3.获取任务列表
	if jobList, err = G_jobMgr.ListLettersByWallType(string(postJob)); err != nil {
		goto ERR
	}

	//4.正常应答
	if bytes, err = common.BuildResponse(0, "success", jobList); err == nil {
		w.Write(bytes)
	}
	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

// 强制杀死某个任务
// POST JOB_KILLER_LETTER = "/ma/wall/kill/"  name=job1
func HandleJobKill(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		name  string
		bytes []byte
	)

	// 解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	// 要杀死的任务名
	name = r.PostForm.Get("name")

	// 杀死任务
	if err = G_jobMgr.KillJob(name); err != nil {
		goto ERR
	}

	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		w.Write(bytes)
	}
	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

func SearchLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		postJob string
		bytes   []byte
		lets    []*model.Letter
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("id1") //获取userID

	//3.调用SearchLetterByIDandTime  检查标签
	lets, err = dao.SearchLetterById(postJob)

	//4.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", lets); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 5, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

// 查询任务日志
func HandleJobLog(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		taskPath   string // 任务路径
		skipParam  string // 从第几条开始
		limitParam string // 返回多少条
		skip       int
		limit      int
		logArr     []*common.JobLog
		bytes      []byte
	)

	// 解析GET参数
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	// 获取请求参数 /job/administrator4/log?name=job10&skip=0&limit=10
	taskPath = r.Form.Get("taskPath")
	skipParam = r.Form.Get("skip")
	limitParam = r.Form.Get("limit")
	if skip, err = strconv.Atoi(skipParam); err != nil {
		skip = 0
	}
	if limit, err = strconv.Atoi(limitParam); err != nil {
		limit = 20
	}

	if logArr, err = G_logMgr.ListLog(taskPath, skip, limit); err != nil {
		goto ERR
	}

	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", logArr); err == nil {
		w.Write(bytes)
	}
	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

// 获取健康worker节点列表
func HandleWorkerList(w http.ResponseWriter, r *http.Request) {
	var (
		workerArr []*common.WorkersIP
		err       error
		postJob   string
		bytes     []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("workerDir") //获取workerDir

	//3.ListWorkers
	if workerArr, err = G_workerMgr.ListWorkers(postJob); err != nil {
		goto ERR
	}

	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", workerArr); err == nil {
		w.Write(bytes)
	}
	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

//考虑要不要加入管理员客服服务
