package main

import (
	"fmt"
	"net/http"
	"time"
	// "time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 3000),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Starting server on :3000")
	return srv.ListenAndServe()
}
