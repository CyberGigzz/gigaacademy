package main

import (
	"net/http"

	"github.com/CyberGigzz/gigaacademy/internal/handlers"
	"github.com/CyberGigzz/gigaacademy/internal/services"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	services := services.NewServices(&app.models)
	handler := handlers.NewHandler(services)

	// Courses
	r.Route("/v1/courses", func(r chi.Router) {
		r.Get("/", handler.CourseHandler.GetAllCoursesHandler)
		r.Get("/{id}", handler.CourseHandler.GetCourseByIDHandler)
		// r.Post("/", handler.CourseHandler.CreateCourseHandler)
	})

	return r
}
