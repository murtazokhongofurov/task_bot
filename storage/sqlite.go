package storage

import (
	"database/sql"
	"errors"

	"gitlab.com/task_bot/storage/models"
)




func (b *storagePg) Create(user *models.User) (*models.User, error) {
	var res models.User
	query := `
	INSERT INTO 
		users(tg_id, tg_name, full_name, phone_number, step)
	VALUES
		(?, ?, ?, ?, ?)
	RETURNING
		full_name, phone_number, created_at
	`
	err := b.db.QueryRow(query, user.TgId, user.TgName, user.FullName, user.PhoneNumber, user.Step).
	Scan(&res.FullName, &res.PhoneNumber, &res.CreateAt)
	if err != nil {
		return &models.User{}, err
	}

	return &res, nil
}

func (b *storagePg) Get(id int64) (*models.User, error) {
	var res models.User
	query := `
	SELECT
		tg_id, full_name, phone_number, step, created_at
	FROM 
		users
	WHERE 
		tg_id=?
	`

	err := b.db.QueryRow(query, id).Scan(
		&res.TgId, &res.FullName, &res.PhoneNumber, &res.Step, &res.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}


func (b *storagePg) GetOrCreate(TgId int64, TgName string) (*models.User, error) {
	user, err := b.Get(TgId)
	if errors.Is(err, sql.ErrNoRows) {
		u, err := b.Create(&models.User{
			TgId: TgId,
			TgName: TgName,
			Step: EnterFullnameStep,
		})
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
	result, err := b.db.Exec(query, step, TgId)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}
	return nil
}