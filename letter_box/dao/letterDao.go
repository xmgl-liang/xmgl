package dao

import (
	"aaa/letter_box/dao/DB"
	"aaa/letter_box/model"
	"fmt"
)

//查询信件
func SearchLetByUUID1and2(id1 string, id2 string) ([]*model.Letter, error) {
	//1.sql语句(结合时间和签收标识和ID)
	sqlStr := `select LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType from letter WHERE UUID1 = ? AND UUID2 = ?`
	//2.执行
	rows, _ := DB.G_jobDB.Query(sqlStr, id1, id2)

	var lets []*model.Letter

	for rows.Next() {
		let := &model.Letter{}

		err := rows.Scan(&let.LetterID, &let.UUID1, &let.UUID2, &let.UserName, &let.Content, &let.Label, &let.WallType)
		if err != nil {
			return nil, err
		}
		fmt.Println(let)
		fmt.Println("no")
		lets = append(lets, let)
	}
	
	sqlStr = `select LetterUUID, UUID1, UUID2, UserName, Content, Label, WallType from letter WHERE UUID1 = ? AND UUID2 = ? AND Label = ?`

	rows, _ = DB.G_jobDB.Query(sqlStr, id2, id1, "2")

	for rows.Next() {
		let := &model.Letter{}

		err := rows.Scan(&let.LetterID, &let.UUID1, &let.UUID2, &let.UserName, &let.Content, &let.Label, &let.WallType)
		if err != nil {
			return nil, err
		}
		fmt.Println(let)
		fmt.Println("no1")
		lets = append(lets, let)
	}


	return lets, nil
}
