package model

import "time"

type Possession struct {
	ID      string    `db:"id"`
	UserID  string    `db:"user_id"`
	CharaID string    `db:"chara_id"`
	RegAt   time.Time `db:"reg_at"`
}
