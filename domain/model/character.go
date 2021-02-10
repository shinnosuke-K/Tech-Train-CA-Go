package model

import "time"

type Character struct {
	Id     string    `db:"id"`
	Name   string    `db:"chara_name"`
	RegAt  time.Time `db:"reg_at"`
	Rarity int       `db:"rarity"`
}
