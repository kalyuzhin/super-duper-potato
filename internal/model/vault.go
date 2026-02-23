package model

import "time"

// VaultData – ...
type VaultData struct {
	ID            uint64    `db:"id"`
	Service       string    `db:"service"`
	Login         []byte    `db:"login"`
	LoginNonce    []byte    `db:"login_nonce"`
	Password      []byte    `db:"password"`
	PasswordNonce []byte    `db:"password_nonce"`
	Secret        []byte    `db:"secret"`
	SecretNonce   []byte    `db:"secret_nonce"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// VaultDataDTO – ...
type VaultDataDTO struct {
	Service       string `db:"service"`
	Login         []byte `db:"login"`
	LoginNonce    []byte `db:"login_nonce"`
	Password      []byte `db:"password"`
	PasswordNonce []byte `db:"password_nonce"`
}
