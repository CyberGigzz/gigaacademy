package services

import "fmt"

type CourseService struct{}

func (cs *CourseService) GetCourse(courseID int) string {
	return fmt.Sprintf("Course with ID %d retrieved.", courseID)
}

func (cs *CourseService) GetAllCourses() string {
	return "All courses retrieved."
}
