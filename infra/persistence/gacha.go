package persistence

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type gachaPersistence struct {
	DB *sql.DB
}

func NewGachaPersistence(db *sql.DB) repository.GachaRepository {
	return &gachaPersistence{
		DB: db,
	}
}

func (g gachaPersistence) GetCharacter() ([]*model.Character, error) {

	rows, err := g.DB.Query("select * from characters")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var characters []*model.Character
	for rows.Next() {
		var chara model.Character
		if err := rows.Scan(&chara.ID, &chara.Name, &chara.RegAt, &chara.Rarity, &chara.Weight); err != nil {
			return nil, errors.WithStack(err)
		}
		characters = append(characters, &chara)
	}
	return characters, nil
}

func (g gachaPersistence) Store(tx *sql.Tx, p *model.Possession) error {

	_, err := tx.Exec("insert into possessions(id, user_id, chara_id, reg_at) values (?,?,?,?)", p.ID, p.UserID, p.CharaID, p.RegAt)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("insert into possessions_composite_key(user_id, possession_id, chara_id, reg_at) values (?,?,?,?)", p.UserID, p.ID, p.CharaID, p.RegAt)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("insert into possessions_index(id, user_id, chara_id, reg_at) values (?,?,?,?)", p.ID, p.UserID, p.CharaID, p.RegAt)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("insert into possessions_not_index(id, user_id, chara_id, reg_at) values (?,?,?,?)", p.ID, p.UserID, p.CharaID, p.RegAt)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
