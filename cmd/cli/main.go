package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
	"github.com/kalyuzhin/password-manager/internal/service"
	"github.com/kalyuzhin/password-manager/pkg/cobra"
)

func main() {
	_ = context.Background()
	db, err := sqlite.NewDB(defaultDBPath())
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewService(db)

	cmd := cobra.NewRootCmd(s)
	if err = cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func defaultDBPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "passwords.db"
	}

	dir := filepath.Join(home, ".password-manager")

	_ = os.Mkdir(dir, 0700)

	return filepath.Join(dir, "passwords.db")
}
