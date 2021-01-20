package model

import "time"

type User struct {
	UserId   string    `gorm:"type:varchar(32);column:user_id;primary_key"`
	UserName string    `gorm:"type:varchar(255);column:user_name"`
	RegAt    time.Time `gorm:"type:datetime;column:reg_at"`
	UpdateAt time.Time `gorm:"type:datetime;column:update_at"`
}
