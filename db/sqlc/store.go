package db

import (
	"database/sql"
)

// create Store struct to hold queries and db pointers.
type Store struct {
	*Queries
	db *sql.DB
}

// create NewStore function which returns Store pointer to db conncetion.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
