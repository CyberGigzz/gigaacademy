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

func (r *CourseRepository) GetCourseByID(id int) (*models.Course, error) {
	course := &models.Course{}
	err := r.db.QueryRow("SELECT id, name, description, code, credits, start_date, end_date, max_capacity, status, created_at, updated_at FROM courses WHERE id = $1", id).Scan(
		&course.ID, &course.Name, &course.Description, &course.Code, &course.Credits, &course.StartDate, &course.EndDate, &course.MaxCapacity, &course.Status, &course.CreatedAt, &course.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (r *CourseRepository) CreateCourse(course *models.Course) error {
	_, err := r.db.Exec("INSERT INTO courses (name, description, code, credits, start_date, end_date, max_capacity, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		course.Name, course.Description, course.Code, course.Credits, course.StartDate, course.EndDate, course.MaxCapacity, course.Status,
	)
	if err != nil {
		return err
	}
	return nil
}
