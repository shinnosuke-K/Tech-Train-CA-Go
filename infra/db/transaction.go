package db

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"
)

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) repository.Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(txFunc func(*sql.Tx) error) (err error) {
	tx, err := t.db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	defer func() {
		if p := recover(); p != nil {
			logger.Log.Error("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			logger.Log.Error("rollback")
			tx.Rollback()
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				tx.Rollback()
				err = commitErr
			}
		}
	}()

	return txFunc(tx)
}
