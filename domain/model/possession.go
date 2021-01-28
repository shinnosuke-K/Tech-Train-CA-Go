package model

import "time"

type Possession struct {
	PosseId string
	UserId  string
	CharaId string
	RegAt   time.Time
}
