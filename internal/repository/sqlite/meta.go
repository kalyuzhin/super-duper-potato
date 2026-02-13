package sqlite

import (
	"context"

	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// GetMetaByName – ...
func (db *DB) GetMetaByName(ctx context.Context, name string) (meta model.MetaData, err error) {
	q := `
	SELECT name, kdf_type, salt, time, threads, memory, key_length
	FROM meta
	WHERE name = $1;`

	err = db.QueryRow(ctx, q, name).Scan(&meta.Name, &meta.KDFType, &meta.Salt, &meta.KDFTime, &meta.KDFThreads,
		&meta.KDFMemory, &meta.KDFKeyLength)
	if err != nil {
		return meta, errorspkg.Wrap(checkSQLErrors(err), "can't get meta by meta")
	}

	return meta, nil
}

// GetMetaByUserID – ...
func (db *DB) GetMetaByUserID(ctx context.Context, userID int64) (meta model.MetaData, err error) {
	q := `
	SELECT name, kdf_type, salt, time, threads, memory, key_length
	FROM meta
	WHERE user_id = $1;`

	err = db.QueryRow(ctx, q, userID).Scan(&meta.Name, &meta.KDFType, &meta.Salt, &meta.KDFTime, &meta.KDFThreads,
		&meta.KDFMemory, &meta.KDFKeyLength)
	if err != nil {
		return meta, errorspkg.Wrap(checkSQLErrors(err), "can't get meta by user id")
	}

	return meta, nil
}

// InsertMeta – ...
func (db *DB) InsertMeta(ctx context.Context, userID int64, meta model.MetaData) error {
	q := `
	INSERT INTO meta(name, user_id, kdf_type, salt, time, threads, memory, key_length)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := db.Exec(ctx, q, meta.Name, userID, meta.KDFType, meta.Salt, meta.KDFTime, meta.KDFThreads,
		meta.KDFMemory, meta.KDFKeyLength)
	if err != nil {
		return errorspkg.Wrap(checkSQLErrors(err), "can't insert meta")
	}

	return nil
}
