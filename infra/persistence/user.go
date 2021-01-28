package persistence

import (
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

func (u userPersistence) IsRecord(id string) bool {
	var user model.User
	err := u.DB.Where("user_id=?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u userPersistence) Add(user *model.User) error {
	if err := u.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u userPersistence) Get(id string) (*model.User, error) {
	var user model.User
	if err := u.DB.Where("user_id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userPersistence) Update(user *model.User) error {
	if err := u.DB.Model(&model.User{}).Update(user).Error; err != nil {
		return err
	}
	return nil
}