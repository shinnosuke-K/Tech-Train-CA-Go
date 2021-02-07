package model

import "time"

type User struct {
	ID       string
	Name     string
	RegAt    time.Time
	UpdateAt time.Time
}
