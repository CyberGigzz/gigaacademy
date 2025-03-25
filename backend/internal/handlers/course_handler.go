package handlers

import (
	"encoding/json"
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
	courses, err := ch.service.GetAllCourses()
	if err != nil {
		http.Error(w, "Failed to fetch courses: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
