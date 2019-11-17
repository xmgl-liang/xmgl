package tasks

import (
	"aaa/user_worker/common"
	"aaa/user_worker/tasks/DB"
	"fmt"
)

//SaveNumPublish
func UpdateUserMsg(job *common.UserMsg) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set Name = ?, Sex = ?, Text = ? WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UserName, job.Sex, job.Text, job.UserUUID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
