package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/CyberGigzz/gigaacademy/internal/models"
	"github.com/CyberGigzz/gigaacademy/internal/services"
	"github.com/go-chi/chi/v5"
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

func (ch *CourseHandler) GetCourseByIDHandler(w http.ResponseWriter, r *http.Request) {
	courseIDStr := chi.URLParam(r, "id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	course, err := ch.service.GetCourseByID(courseID)
	if err != nil {
		http.Error(w, "Failed to fetch course: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func (ch *CourseHandler) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	if err := ch.service.CreateCourse(&course); err != nil {
		http.Error(w, "Failed to create course: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}
