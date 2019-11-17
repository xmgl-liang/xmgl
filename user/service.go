package user

import (
	"aaa/user/common"
	"aaa/user/dao"
	"aaa/user/model"
	"encoding/json"
	"net/http"
	"fmt"
)

//1.修改信息在etcd
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		postJob string
		job     common.UserMsg
		bytes   []byte
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJob = r.PostForm.Get("userMsg")

	//3.反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	//4.保存到etcd
	if err = G_jobMgr.SaveUserUpdata(&job); err != nil {
		goto ERR
	}
	fmt.Println("111")

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

//2.查询信息
func SearchUserIN(w http.ResponseWriter, r *http.Request) {

	var (
		err        error
		postJobID  string
		bytes      []byte
		userStatue *model.User
	)

	//1.解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}

	//2.取表单中的job字段
	postJobID = r.PostForm.Get("userID") //获取userID

	//3.调用SearchUserByID  检查标签
	userStatue, err = dao.SearchUserByID(postJobID)

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
