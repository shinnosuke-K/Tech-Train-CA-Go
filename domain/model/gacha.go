package model

type Gacha struct {
	ID      int `db:"id"`
	Rarity  int `db:"rarity"`
	Weights int `db:"weights`
}
