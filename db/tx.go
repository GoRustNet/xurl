package db

import "github.com/jmoiron/sqlx"

func Tx() (*sqlx.Tx, error) {
	return db.Beginx()
}
