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
	var gachaRate []*model.Gacha
	if err := g.DB.Find(&gachaRate).Error; err != nil {
		return nil, err
	}
	return gachaRate, nil
}

func (g gachaPersistence) GetCharacter() ([]*model.Character, error) {
	var characters []*model.Character
	if err := g.DB.Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}
