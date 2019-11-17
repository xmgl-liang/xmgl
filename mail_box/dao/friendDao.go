package dao

import (
	"aaa/mail_box/dao/DB"
	"aaa/mail_box/model"
)

//查询好友
func SearchFriByUUID(userID string) ([]*model.Friend, error) {
	//1.sql语句
	sqlStr := `select MeUUID, FriUUID from friends where MeUUID = ?`
	//2.执行
	rows, _ := DB.G_jobDB.Query(sqlStr, userID)

	var fris []*model.Friend

	for rows.Next() {
		fri := &model.Friend{}

		err := rows.Scan(&fri.MeUUID, &fri.FriUUID)
		if err != nil {
			return nil, err
		}

		fris = append(fris, fri)
	}

	return fris, nil
}
