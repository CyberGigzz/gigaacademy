package main

import (
	"net/http"

	"github.com/CyberGigzz/gigaacademy/internal/handlers"
	"github.com/CyberGigzz/gigaacademy/internal/services"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	services := services.NewServices()
	handler := handlers.NewHandler(services)

	// Courses
	r.Get("/courses", handler.CourseHandler.GetAllCoursesHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}
