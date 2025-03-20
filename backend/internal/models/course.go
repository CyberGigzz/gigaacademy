package models

import "time"

type CourseStatus string

const (
	ActiveStatus   CourseStatus = "active"
	InactiveStatus CourseStatus = "inactive"
	ArchivedStatus CourseStatus = "archived"
)

type Course struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Code        string       `json:"code"`
	Credits     int          `json:"credits"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	MaxCapacity int          `json:"max_capacity"`
	Status      CourseStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func (c *Course) IsActive() bool {
	return c.Status == ActiveStatus
}
