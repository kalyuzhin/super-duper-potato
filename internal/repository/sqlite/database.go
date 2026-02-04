package sqlite

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

const (
	driverName = "sqlite3"
)

// DB – wrapper for sqlite storage
type DB struct {
	DB *sql.DB
}

// NewDB – ...
func NewDB(filePath string) (*DB, error) {
	db, err := sql.Open(driverName, filePath)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}

// Exec – ...
func (db *DB) Exec(_ context.Context, q string, args ...any) (sql.Result, error) {
	return db.DB.Exec(q, args)
}

// QueryRow – ...
func (db *DB) QueryRow(_ context.Context, q string, args ...any) *sql.Row {
	return db.DB.QueryRow(q, args)
}
