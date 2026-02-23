package main

import (
	"context"
	"log"

	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
	"github.com/kalyuzhin/password-manager/internal/service"
	"github.com/kalyuzhin/password-manager/pkg/cobra"
)

func main() {
	_ = context.Background()
	db, err := sqlite.NewDB("passwords.db")
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewService(db)

	cobra.Execute(s)
}
