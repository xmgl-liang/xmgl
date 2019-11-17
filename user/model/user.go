package model

type User struct {
	UserUUID string //用户的UUID
	Name     string //用户的名字
	Sex      string //用户的性别
	Text     string //座右铭
	//Friends    []*Friend //用户的好友组
	NumFri     int //用户的好友数量
	NumCharge  int //用户的收信数
	NumPublish int //用户的发表数
}
