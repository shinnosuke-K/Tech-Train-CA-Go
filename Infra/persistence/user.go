package persistence

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type userPersistence struct {
	DB *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{
		DB: db,
	}
}

func (u userPersistence) IsRecord(ctx context.Context, id int) bool {
	var user model.User
	err := u.DB.Where("user_id=?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u userPersistence) Add(user *model.User) error {
	panic("implement me")
}

func (u userPersistence) Get(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}

func (u userPersistence) Update(ctx context.Context) error {
	panic("implement me")
}
