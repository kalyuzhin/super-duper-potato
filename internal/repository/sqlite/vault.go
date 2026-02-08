package sqlite

import (
	"context"

	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// GetVaultDataByService – ...
func (db *DB) GetVaultDataByService(ctx context.Context, service string) (data model.VaultData, err error) {
	q := `
	SELECT id, service, login, login_nonce, password, password_nonce, created_at, updated_at
	FROM vault
	WHERE service = $1;`

	err = db.QueryRow(ctx, q, service).Scan(&data.ID, &data.Service, &data.Login, &data.LoginNonce, &data.Password, &data.PasswordNonce, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, errorspkg.Wrap(err, "can't get vault data by service")
	}

	return data, nil
}

// InsertVaultData – ...
func (db *DB) InsertVaultData(ctx context.Context, data model.VaultData) error {
	q := `
	INSERT INTO vault(service, login, login_nonce, password, password_nonce)
	VALUES ($1, $2, $3, $4, $5);`

	_, err := db.Exec(ctx, q, data.Service, data.Login, data.LoginNonce, data.Password, data.PasswordNonce)
	if err != nil {
		return errorspkg.Wrap(err, "can't insert vault data")
	}

	return nil
}
