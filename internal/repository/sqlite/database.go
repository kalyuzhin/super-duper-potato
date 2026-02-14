package sqlite

import (
	"context"
	"database/sql"
	"strings"

	_ "modernc.org/sqlite"

	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

const (
	driverName = "sqlite"
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
	return db.DB.Exec(q, args...)
}

// QueryRow – ...
func (db *DB) QueryRow(_ context.Context, q string, args ...any) *sql.Row {
	return db.DB.QueryRow(q, args...)
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
	number, err := rows.RowsAffected()
	if number == 0 {
		return errorspkg.New("0 rows affected")
	}

	return err
}
