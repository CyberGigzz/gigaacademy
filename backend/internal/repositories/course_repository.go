package repositories

import (
	"database/sql"

	"github.com/CyberGigzz/gigaacademy/internal/models"
)

type CourseRepository struct {
	db *sql.DB
}

func (r *CourseRepository) GetAllCourses() ([]*models.Course, error) {

	rows, err := r.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*models.Course{}
	for rows.Next() {
		course := &models.Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Code, &course.Credits, &course.StartDate, &course.EndDate, &course.MaxCapacity, &course.Status, &course.CreatedAt, &course.UpdatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
