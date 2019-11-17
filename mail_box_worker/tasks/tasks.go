package tasks

import (
	"aaa/mail_box_worker/common"
	"aaa/mail_box_worker/tasks/DB"
	"fmt"
)

//SaveMailFriend
func AddMailFriend(job *common.MailFriend) error {
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

//0拒绝  1可接受  -1即刻   2接受
//SaveLetterByFri
func AddLetterByFri(job *common.LetterByFri) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE letter set Label = ? WHERE LetterUUID = ? AND UUID1 = ? AND UUID2 = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.Label, job.LetterID, job.UUID1, job.UUID2)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil

	// var (
	// 	sqlStr string
	// 	err    error
	// )

	// //1.sql语句
	// sqlStr = `INSERT INTO letter (LetterUUID,UUID1,UUID2,UserName,Content,Label,WallType) VALUES (?,?,?,?,?,?,?)`
	// //2.执行
	// _, err = DB.G_jobDB.Exec(sqlStr, job.LetterID, job.UUID1, job.UUID2, job.UserName, job.Content, job.Label, job.WallType)
	// if err != nil {
	// 	fmt.Println("执行出现异常：", err)
	// 	return err
	// }
	// return nil
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
	sqlStr = `UPDATE users set NumFri = NumFri - 1 WHERE UserUUID = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.UUID1)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//0拒绝  1可接受  -1即刻   2接受
//SaveNoLetter
func UpdateNoLetter(job *common.NoLetter) error {
	var (
		sqlStr string
		err    error
	)

	//1.sql语句
	sqlStr = `UPDATE letter set Label = ? WHERE LetterUUID = ? AND UUID1 = ? AND UUID2 = ?`
	//2.执行
	_, err = DB.G_jobDB.Exec(sqlStr, job.Label, job.LetterID, job.UUID1, job.UUID2)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
