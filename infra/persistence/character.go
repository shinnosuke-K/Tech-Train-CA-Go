package persistence

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type characterPersistence struct {
	DB *sql.DB
}

func NewCharaPersistence(db *sql.DB) repository.CharacterRepository {
	return &characterPersistence{
		DB: db,
	}
}

func (c characterPersistence) GetCharacter(id string) (*model.Character, error) {

	rows, err := c.DB.Query("select * from characters where id = ?", id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var chara model.Character
	for rows.Next() {
		if err := rows.Scan(&chara.Id, &chara.Name, &chara.RegAt, &chara.Rarity); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &chara, nil
}

func (c characterPersistence) GetPossession(userId string) ([]*model.Possession, error) {

	rows, err := c.DB.Query("select * from possessions where user_id = ?", userId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var possessions []*model.Possession
	for rows.Next() {
		var pos model.Possession
		if err := rows.Scan(&pos.Id, &pos.UserId, &pos.CharaId, &pos.RegAt); err != nil {
			return nil, errors.WithStack(err)
		}

		possessions = append(possessions, &pos)
	}

	return possessions, nil
}
