package mail_box

import (
	"aaa/mail_box/common"
	"aaa/mail_box/dao"
	"aaa/mail_box/model"
	"encoding/json"
	"net/http"
)

//0拒绝  1可接受  -1即刻   2已经接受

//1.接受信件(查询的好友列表去判断)，建立好友关系（etcd），获取收信编入（etcd）历时交流信件的过程
func HandleAcceptLetterMail(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		postJobFri    string
		postJobLetter string
		postJobNumC   string
		postJobNumF   string
		jobFri        common.MailFriend
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
	if err = G_jobMgr.SaveMailFriend(&jobFri); err != nil {
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

//2.拒绝信件，（etcd）屏蔽信件查询
func HandleNoLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		postJob string
		job     common.NoLetter
		bytes   []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("refuse")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	//4.保存到etcd
	if err = G_jobMgr.SaveNoLetter(&job); err != nil {
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

//3.查询信件
func SearchLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		postJobID   string
		postJobTime string
		bytes       []byte
		lets        []*model.Letter
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJobID = r.PostForm.Get("userID")    //获取userID
	postJobTime = r.PostForm.Get("nowTime") //获取当前时间

	//3.调用SearchLetterByIDandTime  检查标签
	lets, err = dao.SearchLetterByIDandTime(postJobID, postJobTime)

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
