package tasks

import (
	"aaa/wall_worker/common"
	"aaa/wall_worker/tasks/DB"
	"fmt"
)

//SaveJobWallLetter
func AddWallLetterJob(job *common.WallLetterJob) error {
	var (
		sqlStr string
		err    error
		//msg    []byte
	)
	//1.sql语句
	sqlStr = `INSERT INTO wall_letter (LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType) VALUES (?,?,?,?,?,?,?)`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.LetterID, job.UUID1, job.UUID2, job.UserName, job.Content, job.Label, job.WallType)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		//msg = {"执行出现异常"}
		return err
	}
	//msg = {"信件ID" + job.LetterID + " UUID1" + job.UUID1 + " UUID2" + job.UUID2 + " WallType" + job.WallType}
	return nil
}

//SaveNumPublish
func UpdateNumPublish(job *common.UpdateNumPublishJob) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set NumPublish = NumPublish - 1 WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//SaveNumCharge
func UpdateNumCharge(job *common.UpdateNumChargeJob) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE users set NumCharge = NumCharge - 1 WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//SaveNumFri
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

//SaveWallFriend
func AddWallFriend(job *common.WallFriend) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `INSERT INTO friends (MeUUID,FriUUID) VALUES (?,?)`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1, job.UUID2)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//SaveLetterByFri
func AddLetterByFri(job *common.LetterByFri) error {
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
