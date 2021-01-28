package model

import "time"

type Character struct {
	Id     string    `gorm:"type:int;column:chara_id;primary_key"`
	Name   string    `gorm:"type:varchar(255);column:chara_name"`
	RegAt  time.Time `gorm:"type:datetime;column:reg_at"`
	Rarity int       `gorm:"type:int;column:rarity"`
}
