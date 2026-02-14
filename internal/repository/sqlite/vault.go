package sqlite

import (
	"context"

	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// GetVaultDataByService – ...
func (db *DB) GetVaultDataByService(ctx context.Context, service string) (data model.VaultData, err error) {
	q := `
	SELECT id, service, login, login_nonce, password, password_nonce
	FROM vault
	WHERE service = $1;`

	err = db.QueryRow(ctx, q, service).Scan(&data.ID, &data.Service, &data.Login, &data.LoginNonce, &data.Password, &data.PasswordNonce)
	if err != nil {
		return data, errorspkg.Wrap(err, "can't get vault data by service")
	}

	return data, nil
}

// InsertVaultData – ...
func (db *DB) InsertVaultData(ctx context.Context, userID int64, data model.VaultData) error {
	q := `
	INSERT INTO vault(service, user_id, login, login_nonce, password, password_nonce)
	VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := db.Exec(ctx, q, data.Service, userID, data.Login, data.LoginNonce, data.Password, data.PasswordNonce)
	if err != nil {
		return errorspkg.Wrap(err, "can't insert vault data")
	}

	return nil
}

// CheckVaultDataExistsByServiceUserID – ...
func (db *DB) CheckVaultDataExistsByServiceUserID(ctx context.Context, userID int64, service string) (exists bool, err error) {
	q := `
	SELECT EXISTS(SELECT id FROM vault WHERE service = $1 AND user_id = $2);`

	err = db.QueryRow(ctx, q, service, userID).Scan(&exists)
	if err != nil {
		return false, errorspkg.Wrap(checkSQLErrors(err), "can't check vault data exists by service user id")
	}

	return exists, nil
}

// DeleteVaultDataUser – ...
func (db *DB) DeleteVaultDataUser(ctx context.Context, userID int64, service string) error {
	q := `
	DELETE FROM vault
	WHERE user_id = $1 AND service = $2;`

	rows, err := db.Exec(ctx, q, userID, service)
	if err != nil {
		return errorspkg.Wrap(checkSQLErrors(err), "can't delete vault data")
	}

	return checkRowsAffected(rows)
}
