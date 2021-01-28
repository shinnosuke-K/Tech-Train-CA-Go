package model

import "time"

type Character struct {
	CharaId   string
	CharaName string
	RegAt     time.Time
	Rarity    int
}
