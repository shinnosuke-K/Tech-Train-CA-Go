package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type gachaPersistence struct {
	DB *gorm.DB
}

func NewGachaPersistence(db *gorm.DB) repository.GachaRepository {
	return &gachaPersistence{
		DB: db,
	}
}

func (g gachaPersistence) GetRareRate() ([]*model.Gacha, error) {
	panic("implement me")
}

func (g gachaPersistence) GetCharacter() ([]*model.Character, error) {
	panic("implement me")
}
