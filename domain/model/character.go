package model

import "time"

type Character struct {
	CharaId   int       `gorm:"type:int;column:chara_id;primary_key"`
	CharaName string    `gorm:"type:varchar(255);column:chara_name"`
	RegAt     time.Time `gorm:"type:datetime;column:reg_at"`
	Rarity    int       `gorm:"type:int;column:rarity"`
}
