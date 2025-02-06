package main

import (
	"github.com/CyberGigzz/gigaacademy/internal/config"
	"github.com/CyberGigzz/gigaacademy/internal/db"
	"github.com/CyberGigzz/gigaacademy/internal/repositories"
)

type application struct {
	config config.Config
	models repositories.Models
}

func main() {
	cfg := config.Load()
	db, err := db.OpenDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	app := &application{
		config: cfg,
		models: repositories.NewModels(db),
	}

	if err := app.serve(); err != nil {
		panic(err)
	}

}
