package storage

import (
	"database/sql"

	"gitlab.com/task_bot/storage/models"
)

const (
	EnterStartingStep    string = "starting_step"
	ChangeRole 			 string = "change_role"
	RegisteredStep       string = "registered"
)

type StorageI interface {
	Create(*models.User) (*models.User, error)
	GetOrCreate(TgId int64, TgName string) (*models.User, error)
	ChangeStep(TgId int64, step string) error
}	

type storagePg struct {
	db *sql.DB
}

func NewStoragePg(db *sql.DB) StorageI {
	return &storagePg{
		 db: db,
	}
}


