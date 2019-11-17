package dao

import (
	"aaa/administrator/dao/DB"
	"aaa/administrator/model"
)

//0拒绝  1可接受  -1即刻   2接受
//查询信件  （默认管理员的UUID为-2）
func SearchLetterById(id string) ([]*model.Letter, error) {

	//1.sql语句(结合签收标识和ID)  墙上信件回复，即刻收到  -1标识  ID
	sqlStr := `select LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType from letter where UUID2 = ?`
	//2.执行
	rows, _ := DB.G_jobDB.Query(sqlStr, id)

	var lets []*model.Letter

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
