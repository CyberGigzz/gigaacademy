package services

import "github.com/CyberGigzz/gigaacademy/internal/repositories"

type Services struct {
	CourseService *CourseService
}

func NewServices(repos *repositories.Models) *Services {
	return &Services{
		CourseService: NewCourseService(&repos.Course),
	}
}
