package model

import "time"

// KDFType – ...
type KDFType string

const (
	KDFTypeArgon2 = "argon2"
)

// MetaData – ...
type MetaData struct {
	ID           uint64    `db:"id"`
	Name         string    `db:"name"`
	KDFType      KDFType   `db:"kdf_type"`
	Salt         []byte    `db:"salt"`
	KDFTime      uint8     `db:"time"`
	KDFThreads   uint8     `db:"threads"`
	KDFMemory    uint32    `db:"memory"`
	KDFKeyLength uint8     `db:"key_length"`
	CreatedAt    time.Time `db:"created_at"`
}
