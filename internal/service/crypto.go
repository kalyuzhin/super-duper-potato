package service

import (
	"context"
	"github.com/kalyuzhin/password-manager/internal/lib/crypto"
	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// Service – ...
type Service struct {
	cryptoStorage sqlite.DB
}

// NewService – ...
func NewService(storage sqlite.DB) *Service {
	return &Service{cryptoStorage: storage}
}

// SaveNewPassword – ...
func (s *Service) SaveNewPassword(ctx context.Context, userID int64, masterPassword, service, login, password string) error {
	meta, err := s.cryptoStorage.GetMetaByUserID(ctx, userID)
	var key []byte
	if err == nil {
		key = s.getExistingArgon2Key(ctx, masterPassword, meta.Salt)
	}
	if err != nil && errorspkg.Code(err) != errorspkg.NotFound {
		return err
	}
	if err != nil && errorspkg.Code(err) == errorspkg.NotFound {
		key, err = s.getNewArgon2Key(ctx, userID, masterPassword)
		if err != nil {
			return err
		}
	}

	passwordEnc, passwordNonce, err := crypto.Encrypt(password, key)
	if err != nil {
		return err
	}

	loginEnc, loginNonce, err := crypto.Encrypt(login, key)
	if err != nil {
		return err
	}

	err = s.cryptoStorage.InsertVaultData(ctx, userID, model.VaultData{
		Service:       service,
		Password:      passwordEnc,
		PasswordNonce: passwordNonce,
		Login:         loginEnc,
		LoginNonce:    loginNonce,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) getNewArgon2Key(ctx context.Context, userID int64, masterPassword string) (key []byte, err error) {
	key, salt, err := crypto.GenerateArgon2Key(masterPassword)
	if err != nil {
		return nil, err
	}

	meta := model.MetaData{
		KDFType:      model.KDFTypeArgon2,
		KDFKeyLength: 32,
		KDFMemory:    32 * 1024,
		KDFThreads:   4,
		KDFTime:      3,
		Name:         "my",
		Salt:         salt,
	}

	err = s.cryptoStorage.InsertMeta(ctx, userID, meta)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (s *Service) getExistingArgon2Key(ctx context.Context, masterPassword string, salt []byte) []byte {
	return crypto.GetArgonKey(masterPassword, salt)
}

// GenerateNewSecurePassword – ...
func (s *Service) GenerateNewSecurePassword(ctx context.Context, length uint8) (string, error) {
	return crypto.GenerateRandomSecurePassword(length)
}

// GetVaultData – ...
func (s *Service) GetVaultData(ctx context.Context, userID int64, masterPassword, service string) (login, password string, err error) {
	meta, err := s.cryptoStorage.GetMetaByUserID(ctx, userID)
	if err != nil {
		return "", "", err
	}
	key := s.getExistingArgon2Key(ctx, masterPassword, meta.Salt)

	data, err := s.cryptoStorage.GetVaultDataByService(ctx, service)
	if err != nil {
		return "", "", err
	}

	login, err = crypto.Decrypt(data.Login, key, data.LoginNonce)
	if err != nil {
		return "", "", err
	}

	password, err = crypto.Decrypt(data.Password, key, data.PasswordNonce)
	if err != nil {
		return "", "", err
	}

	return login, password, nil
}

// GenerateNewPassword – ...
func (s *Service) GenerateNewPassword(ctx context.Context) {

}
