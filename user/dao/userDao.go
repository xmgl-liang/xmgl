package dao

import (
	"aaa/user/dao/DB"
	"aaa/user/model"
)

//查询用户信息
func SearchUserByID(userID string) (*model.User, error) {
	//1.sql语句
	sqlStr := `select UserUUID, Name, Sex, Text, NumFri, NumCharge, NumPublish from users where UserUUID = ?`
	//2.执行
	row := DB.G_jobDB.QueryRow(sqlStr, userID)
	user := &model.User{}
	err := row.Scan(&user.UserUUID, &user.Name, &user.Sex, &user.Text, &user.NumFri, &user.NumCharge, &user.NumPublish)
	if err != nil {
		return nil, err
	}

	return user, nil
}
