package model

type Gacha struct {
	Id          int     `gorm:"type:int;column:id;primary_key"`
	Rarity      int     `gorm:"type:int;column:rarity"`
	Probability float64 `gorm:"type:float;column:probability"`
}
