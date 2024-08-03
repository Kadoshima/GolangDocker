package domain

import (
	"time"
)

type User struct {
	ID           int
	StudentID    string
	Nickname     string
	Email        string
	Password     string
	DepartmentID int
	CourseID     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
