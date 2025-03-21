package repositories

import (
	"database/sql"

	"github.com/CyberGigzz/gigaacademy/internal/models"
)

type CourseRepository struct {
	DB *sql.DB
}

func (r *CourseRepository) getAllCourses() ([]*models.Course, error) {

	rows, err := r.DB.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*models.Course{}
	for rows.Next() {
		course := &models.Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Code, &course.Credits, &course.StartDate, &course.EndDate, &course.MaxCapacity, &course.Status, &course.CreatedAt, &course.UpdatedAt)
		courses = append(courses, course)
		if err != nil {
			return nil, err
		}
	}

	return courses, nil
}
