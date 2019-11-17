package model

type Letter struct {
	LetterID  string //信件的UUID 发表时间
	UUID1    string //发表人的UUID
	UUID2    string //默认为0
	UserName string //写信人的名字
	Content  string //信件内容
	Label    string //信件标识
	WallType string //墙的类型
}
