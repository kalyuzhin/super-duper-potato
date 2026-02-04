package model

import "time"

// VaultData – ...
type VaultData struct {
	ID        uint64 `db:"id"`
	Service   string `db:"service"`
	Login     []byte `db:"login"`
	Nonce     []byte `db:"nonce"`
	Password  []byte `db:"password"`
	Secret    []byte `db:"secret"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
