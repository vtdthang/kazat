package pgdb

import (
	"database/sql"
	"sync"
)

var once sync.Once

// NewDB to open connection to pg
func NewDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
