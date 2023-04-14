package models

import "time"



type User struct {
	TgId int64 
	TgName string
	FullName *string
	PhoneNumber *string
	Step string
	CreateAt *time.Time
}


