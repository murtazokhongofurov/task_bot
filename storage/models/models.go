package models

import "time"



type User struct {
	TgId int64 
	TgName string
	Step string
	CreateAt *time.Time
}

type UserForList struct {
	TgName string
}

type AllUsers struct {
	Users []UserForList 
}


