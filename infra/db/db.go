package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() (*gorm.DB, error) {
	connectTemplate := "%s:%s@tcp(db:3306)/%s?%s"
	connectPath := fmt.Sprintf(connectTemplate, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTION"))
	db, err := gorm.Open("mysql", connectPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestOpen() (*sql.DB, error) {
	connectTemplate := "%s:%s@tcp(db:3306)/%s?%s"
	connectPath := fmt.Sprintf(connectTemplate, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTION"))
	db, err := sql.Open("mysql", connectPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
