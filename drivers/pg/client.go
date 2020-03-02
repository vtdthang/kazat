package pgdatabase

import (
	"database/sql"
	"sync"
)

var once sync.Once

// DBCon is connection pool of database/sql
type DBCon struct {
	db *sql.DB
}

// NewDB to open connection to pg
func NewDB(connectionString string) (*DBCon, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DBCon{db}, nil
}
