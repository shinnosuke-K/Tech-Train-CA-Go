package repository

import (
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type UserRepository interface {
	IsRecord(id int) bool
	Add(user *model.User) error
	Get(id int) (*model.User, error)
	Update() error
}
