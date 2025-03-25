package services

import (
	"fmt"

	"github.com/CyberGigzz/gigaacademy/internal/models"
	"github.com/CyberGigzz/gigaacademy/internal/repositories"
)

type CourseService struct {
	repo *repositories.CourseRepository
}

func NewCourseService(repo *repositories.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (cs *CourseService) GetAllCourses() ([]*models.Course, error) {
	return cs.repo.GetAllCourses()
}

func (cs *CourseService) GetCourse(courseID int) string {
	return fmt.Sprintf("Course with ID %d retrieved.", courseID)
}
