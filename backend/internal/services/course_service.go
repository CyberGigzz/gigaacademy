package services

import (
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

func (cs *CourseService) GetCourseByID(id int) (*models.Course, error) {
	return cs.repo.GetCourseByID(id)
}

func (cs *CourseService) CreateCourse(course *models.Course) error {
	return cs.repo.CreateCourse(course)
}
