package domain

import "time"

type Course struct {
	Id           int
	DepartmentId int
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
