package repository

import (
	"database/sql"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type UserRepository interface {
	IsRecord(id string) bool
	Add(tx *sql.Tx, user *model.User) error
	Get(id string) (*model.User, error)
	Update(tx *sql.Tx, user *model.User) error
}
