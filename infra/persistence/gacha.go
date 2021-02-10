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

func (g gachaPersistence) GetRareRate() ([]*model.Gacha, error) {

	rows, err := g.DB.Query("select * from gachas")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var gachaRate []*model.Gacha
	for rows.Next() {
		var gacha model.Gacha
		if err := rows.Scan(&gacha.Id, &gacha.Rarity, &gacha.Probability); err != nil {
			return nil, errors.WithStack(err)
		}
		gachaRate = append(gachaRate, &gacha)
	}

	return gachaRate, nil
}

func (g gachaPersistence) GetCharacter() ([]*model.Character, error) {

	rows, err := g.DB.Query("select * from characters")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var characters []*model.Character
	for rows.Next() {
		var chara model.Character
		if err := rows.Scan(&chara.Id, &chara.Name, &chara.RegAt, &chara.Rarity); err != nil {
			return nil, errors.WithStack(err)
		}
		characters = append(characters, &chara)
	}
	return characters, nil
}

func (g gachaPersistence) Store(p *model.Possession) error {

	tx, err := g.DB.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = tx.Exec("insert into possessions(id, user_id, chara_id, reg_at) values (?,?,?,?)", p.ID, p.UserID, p.CharaID, p.RegAt)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
