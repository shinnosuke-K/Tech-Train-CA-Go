package model

import "time"

type User struct {
	ID       string    `db:"id"`
	Name     string    `db:"name"`
	RegAt    time.Time `db:"reg_at"`
	UpdateAt time.Time `db:"update_at"`
}
