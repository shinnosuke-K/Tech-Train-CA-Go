package model

type Gacha struct {
	ID          int     `db:"id"`
	Rarity      int     `db:"rarity"`
	Probability float64 `db:"probability`
}
