package models

import "time"



type User struct {
	TgId int64 
	TgName string
	Step string
	CreateAt *time.Time
}

type Admin struct {
	TgId int64
	TgName string
	Step string
}

type UserForList struct {
	TgName string
}

type AllUsers struct {
	Users []UserForList 
}

type UserTgIds struct {
	TgId int
}

type TgIdsList struct {
	TgIds []UserTgIds 
}


