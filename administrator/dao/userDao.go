package dao

import (
	"aaa/administrator/dao/DB"
	"fmt"
)

//12点更新发表接受次数
func UpdateUserMsg() error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set NumCharge = ?, NumPublish = ? `
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, 5, 5)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
