package dao

import (
	"aaa/mail_box/dao/DB"
	"aaa/mail_box/model"
)

//查询用户信息
func SearchUserByUUID(userID string) (*model.User, error) {
	//1.sql语句
	sqlStr := `select UserUUID, Name, Sex, Text, NumFri, NumCharge, NumPublish from users where UserUUID = ?`
	//2.执行
	row := DB.G_jobDB.QueryRow(sqlStr, userID)
	user := &model.User{}
	err := row.Scan(&user.UserUUID, &user.Name, &user.Sex, &user.Text, &user.NumFri, &user.NumCharge, &user.NumPublish)
	if err != nil {
		return nil, err
	}

	//3.查询Friends
	fris, _ := SearchFriByUUID(userID)
	user.Friends = fris

	return user, nil
}
