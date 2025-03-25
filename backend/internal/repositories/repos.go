package repositories

import "database/sql"

type Models struct {
	Course CourseRepository
}

func NewModels(db *sql.DB) Models {
	return Models{
		Course: CourseRepository{db: db},
	}
}
