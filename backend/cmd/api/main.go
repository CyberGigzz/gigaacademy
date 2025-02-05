package main

import (
	"log"
	"net/http"

	"github.com/CyberGigzz/gigaacademy/internal/config"
	"github.com/CyberGigzz/gigaacademy/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	db, err := db.OpenDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Printf("Server listening on :%d", cfg.Port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
