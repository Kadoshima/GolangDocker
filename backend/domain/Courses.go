package domain

import "time"

type Course struct {
	ID           int
	DepartmentID int
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
