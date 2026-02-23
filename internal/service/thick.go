package service

import (
	"context"
	"runtime"

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
	userExists, err := s.cryptoStorage.CheckUserExists(ctx, userID)
	if err != nil {
		return err
	}

	var masterKey []byte
	if userExists {
		meta, err := s.cryptoStorage.GetMetaByUserID(ctx, userID)
		if err != nil {
			return err
		}
		masterKey = s.getExistingArgon2Key(ctx, masterPassword, meta.Salt)
	} else {
		masterKey, err = s.getNewArgon2Key(ctx, userID, masterPassword)
		if err != nil {
			return err
		}
	}

	encKey, authKey, err := crypto.DeriveKeys(masterKey)
	if err != nil {
		return err
	}

	defer func() {
		crypto.ClearMemory(encKey)
		crypto.ClearMemory(authKey)
	}()

	crypto.ClearMemory(masterKey)
	authHash := crypto.GetHash(authKey)

	if userExists {
		storedAuthKey, err := s.cryptoStorage.GetUserAuthKey(ctx, userID)
		if err != nil {
			return err
		}

		if !crypto.CompareHash(authHash, storedAuthKey) {
			return errorspkg.New("auth failed")
		}
	} else {
		err = s.cryptoStorage.InsertUser(ctx, userID, authHash)
		if err != nil {
			return err
		}
	}

	passwordEnc, passwordNonce, err := crypto.Encrypt(password, encKey)
	if err != nil {
		return err
	}

	loginEnc, loginNonce, err := crypto.Encrypt(login, encKey)
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
		KDFKeyLength: 64,
		KDFMemory:    128 * 1024,
		KDFThreads:   uint8(runtime.NumCPU()),
		KDFTime:      4,
		Name:         "my",
		Salt:         salt,
	}

	err = s.cryptoStorage.InsertMeta(ctx, userID, meta)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (s *Service) getExistingArgon2Key(_ context.Context, masterPassword string, salt []byte) []byte {
	return crypto.GetArgonKey(masterPassword, salt)
}

// GenerateNewSecurePassword – ...
func (s *Service) GenerateNewSecurePassword(_ context.Context, length uint8) (string, error) {
	return crypto.GenerateRandomSecurePassword(length)
}

// GetVaultData – ...
func (s *Service) GetVaultData(ctx context.Context, userID int64, masterPassword, service string) (login, password string, err error) {
	meta, err := s.cryptoStorage.GetMetaByUserID(ctx, userID)
	if err != nil {
		return "", "", err
	}
	masterKey := s.getExistingArgon2Key(ctx, masterPassword, meta.Salt)

	encKey, authKey, err := crypto.DeriveKeys(masterKey)
	if err != nil {
		return "", "", err
	}

	defer func() {
		crypto.ClearMemory(encKey)
		crypto.ClearMemory(authKey)
	}()

	crypto.ClearMemory(masterKey)
	authHash := crypto.GetHash(authKey)

	storedAuthKey, err := s.cryptoStorage.GetUserAuthKey(ctx, userID)
	if err != nil {
		return "", "", err
	}

	if !crypto.CompareHash(authHash, storedAuthKey) {
		return "", "", errorspkg.New("auth failed")
	}

	data, err := s.cryptoStorage.GetVaultDataByService(ctx, service)
	if err != nil {
		return "", "", err
	}

	login, err = crypto.Decrypt(data.Login, encKey, data.LoginNonce)
	if err != nil {
		return "", "", err
	}

	password, err = crypto.Decrypt(data.Password, encKey, data.PasswordNonce)
	if err != nil {
		return "", "", err
	}

	return login, password, nil
}

// DeleteVaultData – ...
func (s *Service) DeleteVaultData(ctx context.Context, userID int64, masterPassword, service string) error {
	storedAuthHash, err := s.cryptoStorage.GetUserAuthKey(ctx, userID)
	if err != nil {
		return err
	}

	meta, err := s.cryptoStorage.GetMetaByUserID(ctx, userID)
	if err != nil {
		return err
	}

	masterKey := s.getExistingArgon2Key(ctx, masterPassword, meta.Salt)

	_, authKey, err := crypto.DeriveKeys(masterKey)
	if err != nil {
		return err
	}

	crypto.ClearMemory(masterKey)

	authHash := crypto.GetHash(authKey)
	if !crypto.CompareHash(authHash, storedAuthHash) {
		return errorspkg.New("auth failed")
	}

	err = s.cryptoStorage.DeleteVaultDataUser(ctx, userID, service)
	if err != nil {
		return err
	}

	return nil
}
