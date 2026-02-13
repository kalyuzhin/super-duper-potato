package service

import (
	"context"

	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
)

// Service – ...
type Service struct {
	cryptoStorage sqlite.DB
}

// SaveNewPassword – ...
func (s *Service) SaveNewPassword(ctx context.Context, userID int64, masterPassword, metaName, service, login, password string) {

}

// GetPassword – ...
func (s *Service) GetPassword(ctx context.Context) {

}
