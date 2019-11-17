package letter_box

import (
	"aaa/letter_box/common"
	"aaa/letter_box/dao"
	"aaa/letter_box/model"
	"encoding/json"
	"net/http"
	"fmt"
)

//1.查询好友列表
func SearchFris(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		postJobID  string
		bytes      []byte
		userStatue *model.User
		fris       []*model.User
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段，获取用户ID
	postJobID = r.PostForm.Get("userID") //获取userID

	//3.查询用户信息SearchUserByUUID
	userStatue, err = dao.SearchUserByUUID(postJobID)

	//4.通过用户信息中好友的UUID进行好友的查询，调用SearchFriByUUID
	fris, err = dao.SearchFriMsgByID(userStatue)

	//5.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", fris); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 6, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

//2.查询历时信件
func SearchLetsOld(w http.ResponseWriter, r *http.Request) {
	//1.获取用户ID和好友ID
	var (
		err        error
		postUserID string
		postFriID  string
		bytes      []byte
		lets1      []*model.Letter //用户本人写的信
		//lets2      []*model.Letter //用户好友回的信
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段，获取用户ID
	postUserID = r.PostForm.Get("userID") //获取userID
	postFriID = r.PostForm.Get("friID")   //获取friID
	fmt.Println("okokko111")

	//3.查询信件  调用SearchLetByUUID1and2
	//需要信件数据都不为空哦！！
	lets1, err = dao.SearchLetByUUID1and2(postUserID, postFriID)
	fmt.Println(lets1)
	fmt.Println("okokko")

	//4.返回正确应答（{"errno": 0, "msg": ""}）
	if bytes, err = common.BuildResponse(0, "success", lets1); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	//5, 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

//3.回复信件（etcd）
func HandleReplyLetter(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		postJob  string
		postNumP string
		job      common.LetterJob
		jobNumP  common.UpdateNumPublishJob
		bytes    []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("letterJob")
	postNumP = r.PostForm.Get("numPublish")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}
	//在前端判断是否为0，然后减少1传回后端
	if err = json.Unmarshal([]byte(postNumP), &jobNumP); err != nil {
		goto ERR
	}

	//4.保存到etcd
	if err = G_jobMgr.SaveJobReplyLetter(&job); err != nil {
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

//4.删除好友（etcd）
func HandleDelFri(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		postMsg     string
		postJobNumF string
		jobMsg      common.DelFriend
		jobNumF     common.UpdateNumFriJob
		bytes       []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postMsg = r.PostForm.Get("delMsg")
	postJobNumF = r.PostForm.Get("numFri")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postMsg), &jobMsg); err != nil {
		goto ERR
	}
	if err = json.Unmarshal([]byte(postJobNumF), &jobNumF); err != nil {
		goto ERR
	}

	//4.保存到etcd  建立好友关系 如果已经是好友关系就随它录入数据库失败
	if err = G_jobMgr.SaveDelFriend(&jobMsg); err != nil {
		goto ERR
	}
	if err = G_jobMgr.SaveAddNumFri(&jobNumF); err != nil {
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
