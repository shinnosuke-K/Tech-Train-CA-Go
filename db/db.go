package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() (*gorm.DB, error) {

	connectTemplate := "%s:%s@/%s?%s"
	connectPath := fmt.Sprintf(connectTemplate, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTION"))
	return gorm.Open("mysql", connectPath)
}

type User struct {
	UserId        string    `json:"user_id"`
	Token         string    `json:"token"`
	UserName      string    `json:"user_name"`
	RegTime       time.Time `json:"reg_time"`
	RegTimeJST    time.Time `json:"reg_time_jst"`
	UpdateTime    time.Time `json:"update_time"`
	UpdateTimeJST time.Time `json:"update_time_jst"`
}

func (userInfo *User) IsRecord(DB *gorm.DB) bool {
	var user User
	if err := DB.Where("user_id=?", userInfo.UserId).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return true
	} else {
		return false
	}
}

func (userInfo *User) Insert(DB *gorm.DB) error {
	return DB.Create(&userInfo).Error
}

func Get(DB *gorm.DB, userId string) (*User, error) {
	var getUser User
	if err := DB.Where("user_id=?", userId).First(&getUser).Error; err != nil {
		return nil, err
	} else {
		return &getUser, nil
	}
}
