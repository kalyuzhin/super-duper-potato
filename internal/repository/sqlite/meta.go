package sqlite

import "context"

// GetMetaByName – ...
func (db *DB) GetMetaByName(ctx context.Context, name string) {
	q := `
	SELECT name, kdf_type, salt, time, threads, memory, key_length
	FROM meta
	WHERE name = $1;`

	db.QueryRow(ctx, q, name).Scan()
}
