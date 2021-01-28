package repository

import "github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"

type GachaRepository interface {
	GetRareRate() ([]*model.Gacha, error)
	GetCharacter() ([]*model.Character, error)
	Store(*model.Possession) error
}
