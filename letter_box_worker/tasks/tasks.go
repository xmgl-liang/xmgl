package tasks

import (
	"aaa/letter_box_worker/common"
	"aaa/letter_box_worker/tasks/DB"
	"fmt"
)

//SaveJobJobReplyLetter
func AddLetterJob(job *common.LetterJob) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `INSERT INTO letter (LetterUUID,UUID1,UUID2,UserName,Content,Label,WallType) VALUES (?,?,?,?,?,?,?)`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.LetterID, job.UUID1, job.UUID2, job.UserName, job.Content, job.Label, job.WallType)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//SaveDelFriend
func DeleteDelFriend(job *common.DelFriend) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `DELETE FROM friends WHERE MeUUID = ? AND FriUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1, job.UUID2)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil

}

//SaveAddNumFri
func UpdateNumFri(job *common.UpdateNumFriJob) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set NumFri = NumFri + 1 WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//SaveNumPublish
func UpdateNumPublish(job *common.UpdateNumPublishJob) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set NumPublish = NumPublish - 1  WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
