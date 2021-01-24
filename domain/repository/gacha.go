package repository

import "github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"

type GachaRepository interface {
	IsRecord(table, id string) bool
	GetRareRate() ([]*model.Gacha, error)
	GetCharacter() ([]*model.Character, error)
}
