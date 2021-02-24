package db

import (
	"database/sql"

	"github.com/pkg/errors"
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
		return errors.Wrap(err, "failed to begin transaction")
	}

	if err := txFunc(tx); err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return errors.Wrapf(err, "failed to rollback, error: %s", rollBackErr.Error())
		}
		return errors.Wrap(err, "failed to exec function (rollback was success)")
	}

	if err := tx.Commit(); err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return errors.Wrapf(err, "failed to rollback, error: %s", rollBackErr.Error())
		}
		return errors.Wrap(err, "failed to commit (rollback was success)")
	}
	return nil
}
