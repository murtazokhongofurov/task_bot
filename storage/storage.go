package storage

import (
	"database/sql"

	"gitlab.com/task_bot/storage/models"
)

const (
	EnterStartingStep    string = "starting_step"
	AdminRole 			 string = "admin_role"
	SendMessage 		 string = "send_message_step"
	StatusStep 			 string = "status_step"
)

type StorageI interface {
	Create(*models.User) (*models.User, error)
	GetOrCreate(TgId int64, TgName string) (*models.User, error)
	ChangeStep(TgId int64, step string) error
	GetAllUsers(page, limit int) (*models.AllUsers, error)
	GetAllTgIds() (*models.TgIdsList, error)
	GetUserCount() (*models.TgUserCount, error)
}	

type storagePg struct {
	db *sql.DB
}

func NewStoragePg(db *sql.DB) StorageI {
	return &storagePg{
		 db: db,
	}
}


