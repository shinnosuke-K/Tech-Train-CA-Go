package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func Open() (*gorm.DB, error) {
	connectTemplate := "%s:%s@tcp(db:3306)/%s?%s"
	connectPath := fmt.Sprintf(connectTemplate, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTION"))
	db, err := gorm.Open("mysql", connectPath)
	if err != nil {
		log.Fatalln(err)
	}
	return db, nil
}
