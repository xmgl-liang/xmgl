package dao

import (
	"aaa/mail_box/dao/DB"
	"aaa/mail_box/model"
)

//0拒绝  1可接受  -1即刻   2接受
//查询信件
func SearchLetterByIDandTime(userID string, time string) ([]*model.Letter, error) {

	//1.sql语句(结合时间和签收标识和ID)
	sqlStr := `select LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType from letter where UUID2 = ? AND LetterUUID <= ? AND Label = ?`
	//2.执行
	rows, _ := DB.G_jobDB.Query(sqlStr, userID, time, "1")

	var lets []*model.Letter

	for rows.Next() {
		let := &model.Letter{}

		err := rows.Scan(&let.LetterID, &let.UUID1, &let.UUID2, &let.UserName, &let.Content, &let.Label, &let.WallType)
		if err != nil {
			return nil, err
		}

		lets = append(lets, let)
	}

	//1.sql语句(结合签收标识和ID)  墙上信件回复，即刻收到  -1标识  ID
	sqlStr = `select LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType from letter where Label = ? AND UUID2 = ?`
	//2.执行
	rows, _ = DB.G_jobDB.Query(sqlStr, "-1", userID)

	for rows.Next() {
		let := &model.Letter{}

		err := rows.Scan(&let.LetterID, &let.UUID1, &let.UUID2, &let.UserName, &let.Content, &let.Label, &let.WallType)
		if err != nil {
			return nil, err
		}

		lets = append(lets, let)
	}

	return lets, nil
}
