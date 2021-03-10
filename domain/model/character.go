package model

import "time"

type Character struct {
	ID     string    `db:"id"`
	Name   string    `db:"name"`
	RegAt  time.Time `db:"reg_at"`
	Rarity int       `db:"rarity"`
	Weight int       `db:"weight"`
}
