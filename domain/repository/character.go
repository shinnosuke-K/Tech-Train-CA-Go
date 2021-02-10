package repository

import "github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"

type CharacterRepository interface {
	GetCharacter(id string) (*model.Character, error)
	GetPossession(userID string) ([]*model.Possession, error)
}
