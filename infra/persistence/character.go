package persistence

import (
	"database/sql"
	"strings"

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

func (c characterPersistence) GetCharacters(ids []interface{}) ([]*model.Character, error) {

	query := "select * from characters where id in (?" + strings.Repeat(",?", len(ids)-1) + ")"

	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := stmt.Query(ids...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var charas []*model.Character
	for rows.Next() {
		var chara model.Character
		if err := rows.Scan(&chara.ID, &chara.Name, &chara.RegAt, &chara.Rarity); err != nil {
			return nil, errors.WithStack(err)
		}
		charas = append(charas, &chara)
	}

	return charas, nil
}

func (c characterPersistence) GetPossession(userID string) ([]*model.Possession, error) {

	rows, err := c.DB.Query("select * from possessions where user_id = ?", userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var possessions []*model.Possession
	for rows.Next() {
		var pos model.Possession
		if err := rows.Scan(&pos.ID, &pos.UserID, &pos.CharaID, &pos.RegAt); err != nil {
			return nil, errors.WithStack(err)
		}

		possessions = append(possessions, &pos)
	}

	return possessions, nil
}
