package service

import (
	"context"
	"github.com/kalyuzhin/password-manager/internal/model"
)

// ThinClientService – ...
type ThinClientService struct {
	cryptoStorage Storage
}

// SaveNewPassword – ...
func (s *ThinClientService) SaveNewPassword(ctx context.Context, userID int64, data model.VaultDataDTO) error {
	s.cryptoStorage.GetUserAuthKey(ctx, userID)

	return nil
}
