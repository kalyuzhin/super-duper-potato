package sqlite

import (
	"context"

	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// GetUserAuthKey – ...
func (db *DB) GetUserAuthKey(ctx context.Context, userID int64) (authKey []byte, err error) {
	q := `
	SELECT auth_key
	FROM users
	WHERE id = $1;`

	err = db.DB.QueryRow(q, userID).Scan(&authKey)
	if err != nil {
		return authKey, errorspkg.Wrap(checkSQLErrors(err), "can't get user by id")
	}

	return authKey, nil
}

// CheckUserExists – ...
func (db *DB) CheckUserExists(ctx context.Context, userID int64) (exists bool, err error) {
	q := `
	SELECT EXISTS(SELECT id FROM users WHERE id = 1);`

	err = db.DB.QueryRow(q, userID).Scan(&exists)
	if err != nil {
		return false, errorspkg.Wrap(checkSQLErrors(err), "can't check user exists")
	}

	return exists, nil
}

// InsertUser – ...
func (db *DB) InsertUser(ctx context.Context, userID int64, authHash []byte) error {
	q := `
	INSERT INTO users(id, auth_key)
	VALUES ($1, $2);`

	rows, err := db.DB.Exec(q, userID, authHash)
	if err != nil {
		return errorspkg.Wrap(checkSQLErrors(err), "can't insert user")
	}

	return checkRowsAffected(rows)
}
