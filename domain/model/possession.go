package model

import "time"

type Possession struct {
	PosseId string    `gorm:"type:varchar(36);column:posse_id;primary_key"`
	UserId  string    `gorm:"type:varchar(36);column:user_id;foreign_key"`
	CharaId string    `gorm:"type:varchar(255);column:chara_id"`
	RegAt   time.Time `gorm:"type:datetime;column:reg_at"`
}
