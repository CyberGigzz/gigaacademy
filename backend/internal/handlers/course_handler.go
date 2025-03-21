package handlers

import (
	"net/http"

	"github.com/CyberGigzz/gigaacademy/internal/services"
)

type CourseHandler struct {
	service *services.CourseService
}

func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{service: courseService}
}

func (ch *CourseHandler) GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
	result := ch.service.GetAllCourses() // Corrected line
	w.Write([]byte(result))
}
