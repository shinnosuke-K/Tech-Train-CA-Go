package persistence

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type userPersistence struct {
	DB *sql.DB
}

func NewUserPersistence(db *sql.DB) repository.UserRepository {
	return &userPersistence{
		DB: db,
	}
}

func (u userPersistence) IsRecord(id string) bool {
	_, err := u.DB.Query("select * from users where user_id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u userPersistence) Add(user *model.User) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("insert into users(user_id, user_name, reg_at, update_at) values (?, ?, ?, ?)", user.ID, user.Name, user.RegAt, user.UpdateAt)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (u userPersistence) Get(id string) (*model.User, error) {
	rows, err := u.DB.Query("select * from users where user_id = ?", id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.RegAt, &user.UpdateAt); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &user, nil
}

func (u userPersistence) Update(user *model.User) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("update users set user_name=?, update_at=? where user_id=?", user.Name, user.UpdateAt, user.ID)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
