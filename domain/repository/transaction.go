package repository

import "database/sql"

type Transaction interface {
	DoInTx(func(*sql.Tx) error) error
}
