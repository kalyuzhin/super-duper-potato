package service

import (
	"context"
	"math/rand"

	"github.com/kalyuzhin/password-manager/internal/crypto"
	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

// Service – ...
type Service struct {
	cryptoStorage sqlite.DB
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

	passwordEnc, passwordNonce := crypto.Encrypt(password, key)
	loginEnc, loginNonce := crypto.Encrypt(login, key)

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
	key, salt := crypto.GenerateArgon2Key(masterPassword, rand.Int63())
	meta := model.MetaData{
		KDFType:      "Argon2",
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
func (s *Service) GenerateNewSecurePassword(ctx context.Context, length uint8) string {
	password := crypto.GenerateRandomSecurePassword(length)

	return password
}

// GetPassword – ...
func (s *Service) GetPassword(ctx context.Context) {

}
