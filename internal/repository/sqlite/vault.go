package sqlite

import (
	"context"
	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// GetVaultDataByService – ...
func (db *DB) GetVaultDataByService(ctx context.Context, service string) (data model.VaultData, err error) {
	q := `
	SELECT id, service, login, nonce, password, secret, created_at, updated_at
	FROM vault
	WHERE service = $1;`

	err = db.QueryRow(ctx, q, service).Scan(&data.ID, &data.Service, &data.Login, &data.Nonce, &data.Password, &data.Secret, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, errorspkg.Wrap(err, "can't get vault data by service")
	}

	return data, nil
}
