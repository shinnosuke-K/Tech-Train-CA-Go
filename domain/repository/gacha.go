package repository

import (
	"database/sql"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type GachaRepository interface {
	GetCharacter() ([]*model.Character, error)
	Store(*sql.Tx, *model.Possession) error
}
