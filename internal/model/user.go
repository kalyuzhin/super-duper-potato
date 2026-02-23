package model

import "time"

// User – ...
type User struct {
	ID        uint64
	AuthKey   []byte
	UpdatedAt time.Time
	CreatedAt time.Time
}
