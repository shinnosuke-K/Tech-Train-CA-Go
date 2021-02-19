package repository

import (
	"database/sql"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type GachaRepository interface {
	GetRareRate() ([]*model.Gacha, error)
	GetCharacter() ([]*model.Character, error)
	Store(*sql.Tx, *model.Possession) error
}
