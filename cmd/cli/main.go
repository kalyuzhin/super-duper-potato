package main

import (
	"context"
	"github.com/kalyuzhin/password-manager/internal/repository/sqlite"
	"log"
)

func main() {
	_ = context.Background()
	_, err := sqlite.NewDB("passwords.db")
	if err != nil {
		log.Fatal(err)
	}

}
