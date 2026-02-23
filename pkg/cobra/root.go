package cobra

import (
	"context"

	"github.com/spf13/cobra"
)

// App – ...
type App interface {
	SaveNewPassword(ctx context.Context, userID int64, masterPassword, service, login, password string) error
	GenerateNewSecurePassword(_ context.Context, length uint8) (string, error)
	GetVaultData(ctx context.Context, userID int64, masterPassword, service string) (login, password string, err error)
	DeleteVaultData(ctx context.Context, userID int64, masterPassword, service string) error
}

// NewRootCmd – ...
func NewRootCmd(app App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password-manager",
		Short: "Password Manager",
		Long: `
Password Manager provides local and remote usage`,
	}

	cmd.AddCommand(
		NewGenerateCmd(app),
		NewGetCmd(app),
		NewSaveCmd(app),
		NewRmCmd(app),
	)

	return cmd
}
