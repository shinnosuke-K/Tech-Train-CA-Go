package model

import "time"

type Possession struct {
	Id      string    `db:"id"`
	UserId  string    `db:"user_id"`
	CharaId string    `db:"chara_id"`
	RegAt   time.Time `db:"reg_at"`
}
