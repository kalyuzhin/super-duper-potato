package service

import "context"

// Thin – ...
type Thin interface {
	SaveNewPassword(ctx context.Context, userID int64, service string,
		login, loginNonce, password, passwordNonce []byte) error
	GetVaultData(ctx context.Context, userID int64, authKey []byte,
		service string) (login, loginNonce, password, passwordNonce []byte, err error)
}
