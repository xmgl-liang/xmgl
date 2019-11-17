package wall

import (
	"aaa/wall/common"
	"aaa/wall/dao"
	"aaa/wall/model"
	"encoding/json"
	"fmt"
	"net/http"
)

//1.发表信件到etcd中 设置自动过期   发表次数减少1
func HandleJobWallLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		postJob  string
		postNumP string
		job      common.WallLetterJob
		jobNumP  common.UpdateNumPublishJob
		bytes    []byte
	)
	fmt.Println("111")
	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("wallJob")
	postNumP = r.PostForm.Get("numPublish")
	fmt.Println("222")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}
	fmt.Println("333")
	//在前端判断是否为0，然后减少1传回后端
	if err = json.Unmarshal([]byte(postNumP), &jobNumP); err != nil {
		goto ERR
	}
	fmt.Println("444")
	//4.保存到etcd
	if err = G_jobMgr.SaveJobWallLetter(&job); err != nil {
		goto ERR
	}
	if err = G_jobMgr.SaveNumPublish(&jobNumP); err != nil {
		goto ERR
	}

	//5.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 6, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

//2.查询墙中信件,从etcd中查询
// 列举所有wallLetter任务
func HandleListLettersByWallType(w http.ResponseWriter, r *http.Request) {
	var (
		postJob  string
		jobList  []*common.WallLetterJob
		bytes    []byte
		err      error
		wallType common.WALLType
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("wallType")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJob), &wallType); err != nil {
		goto ERR
	}

	//3.获取任务列表
	if jobList, err = G_jobMgr.ListLettersByWallType(wallType.WallType); err != nil {
		goto ERR
	}

	// for i, j := range jobList {
	// 	fmt.Println(i, j)
	// }

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

//3.接受信件(查询的好友列表去判断)，建立好友关系（etcd），获取收信编入（etcd）历时交流信件的过程
func HandleAcceptLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		postJobFri    string
		postJobLetter string
		postJobNumC   string
		postJobNumF   string
		jobFri        common.WallFriend
		jobLetter     common.LetterByFri
		jobNumC       common.UpdateNumChargeJob
		jobNumF       common.UpdateNumFriJob
		bytes         []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJobFri = r.PostForm.Get("msgJobFri")
	postJobLetter = r.PostForm.Get("msgJobLetter")
	postJobNumC = r.PostForm.Get("numCharge")
	postJobNumF = r.PostForm.Get("numFri")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJobFri), &jobFri); err != nil {
		goto ERR
	}
	if err = json.Unmarshal([]byte(postJobLetter), &jobLetter); err != nil {
		goto ERR
	}
	if err = json.Unmarshal([]byte(postJobNumC), &jobNumC); err != nil {
		goto ERR
	}
	if err = json.Unmarshal([]byte(postJobNumF), &jobNumF); err != nil {
		goto ERR
	}

	//4.保存到etcd  建立好友关系 如果已经是好友关系就随它录入数据库失败
	if err = G_jobMgr.SaveWallFriend(&jobFri); err != nil {
		goto ERR
	}
	if err = G_jobMgr.SaveLetterByFri(&jobLetter); err != nil {
		goto ERR
	}
	if err = G_jobMgr.SaveNumCharge(&jobNumC); err != nil {
		goto ERR
	}
	if err = G_jobMgr.SaveNumFri(&jobNumF); err != nil {
		goto ERR
	}

	//5.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 6, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

//4.查询用户信息
func SearchUserMsg(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		postJob    string
		bytes      []byte
		userStatue *model.User
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("userID") //获取userID

	//3.调用SearchUserByUUID  检查标签
	userStatue, err = dao.SearchUserByUUID(postJob)

	//4.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", userStatue); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 5, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}
