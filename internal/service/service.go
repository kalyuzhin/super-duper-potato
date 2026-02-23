package service

import "context"

type PasswordService interface {
	Save(ctx context.Context) error
}
