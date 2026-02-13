package sqlite

import (
	"context"
	"database/sql"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
	_ "modernc.org/sqlite"
	"strings"
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

func checkSQLErrors(err error) error {
	switch {
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return errorspkg.NewC(err.Error(), errorspkg.NotFound)
	default:
		return nil
	}
}

func checkRowsAffected(rows sql.Result) error {
	_, err := rows.RowsAffected()

	return err
}
