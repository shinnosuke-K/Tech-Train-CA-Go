package db

import (
	"database/sql"
	"log"
)

type Transaction interface {
	DoInTx(func(*sql.Tx) error) error
}

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(txFunc func(*sql.Tx) error) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return nil
}
