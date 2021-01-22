package repository

import (
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type UserRepository interface {
	IsRecord(id string) bool
	Add(user *model.User) error
	Get(id string) (*model.User, error)
	Update(user *model.User) error
}
