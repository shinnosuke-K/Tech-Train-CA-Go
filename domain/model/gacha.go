package model

type Gacha struct {
	Id          int     `db:"id"`
	Rarity      int     `db:"rarity"`
	Probability float64 `db:"probability`
}
