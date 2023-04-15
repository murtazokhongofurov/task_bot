package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"gitlab.com/task_bot/storage/models"
)




func (b *storagePg) Create(user *models.User) (*models.User, error) {
	var res models.User
	statment, err := b.db.Prepare(`INSERT INTO users(tg_id, tg_name, step) VALUES(?, ?, ?)`)
	if err != nil {
		return &models.User{}, err
	}
	result, err := statment.Exec(user.TgId, user.TgName, user.Step)
	if err != nil {
		return &models.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil , err
	}

	err = b.db.QueryRow(`SELECT tg_id, tg_name, step FROM users WHERE tg_id = ?`, id).Scan(&res.TgId, &res.TgName, &res.Step)
	if err != nil {
		return &models.User{}, err
	}
	return &res, nil
}

func (b *storagePg) Get(TgId int64) (*models.User, error) {
	var res models.User
	err := b.db.QueryRow(`SELECT tg_id, tg_name, step FROM users WHERE tg_id = ?`, TgId).Scan(&res.TgId, &res.TgName, &res.Step)
	if err != nil {
		return &models.User{}, err
	}

	return &res, nil
}


func (b *storagePg) GetOrCreate(TgId int64, TgName string) (*models.User, error) {
	user, err := b.Get(TgId)
	if errors.Is(err, sql.ErrNoRows) {
		u, err := b.Create(&models.User{
			TgId: TgId,
			TgName: TgName,
			Step: EnterStartingStep,
		})
		fmt.Println()
		if err != nil {
			return nil, err
		}
		user = u
	}else if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *storagePg) ChangeStep(TgId int64, step string) error {
	query := `UPDATE users SET step=? WHERE tg_id=?`
	statment, err := b.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := statment.Exec(step,TgId)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}
	return nil
}