package service

import (
	"context"

	"github.com/kalyuzhin/password-manager/internal/model"
)

type Storage interface {
	GetMetaByUserID(ctx context.Context, userID int64) (meta model.MetaData, err error)
	InsertMeta(ctx context.Context, userID int64, meta model.MetaData) error
	GetVaultDataByService(ctx context.Context, userID int64, service string) (data model.VaultData, err error)
	InsertVaultData(ctx context.Context, userID int64, data model.VaultData) error
	DeleteVaultDataUser(ctx context.Context, userID int64, service string) error
	GetUserAuthKey(ctx context.Context, userID int64) (authKey []byte, err error)
	CheckUserExists(ctx context.Context, userID int64) (exists bool, err error)
	InsertUser(ctx context.Context, userID int64, authHash []byte) error
}
