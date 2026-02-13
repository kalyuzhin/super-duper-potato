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
		return meta, errorspkg.Wrap(err, "can't get meta by name")
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
		return meta, errorspkg.Wrap(err, "can't get meta by user id")
	}

	return meta, nil
}
