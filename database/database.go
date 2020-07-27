package database

import "database/sql"

type DB struct {
	Conn *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{Conn: db}
}
