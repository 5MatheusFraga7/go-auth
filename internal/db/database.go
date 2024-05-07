package db

import "database/sql"

type Database interface {
	Connect() error
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (*sql.Result, error)
}
