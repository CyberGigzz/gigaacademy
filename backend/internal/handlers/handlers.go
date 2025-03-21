package handlers

import (
	"github.com/CyberGigzz/gigaacademy/internal/services"
)

type Handlers struct {
	CourseHandler *CourseHandler
}

func NewHandler(services *services.Services) *Handlers {
	return &Handlers{
		CourseHandler: NewCourseHandler(services.CourseService),
	}
}
