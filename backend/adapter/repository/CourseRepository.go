package repository

import "backend/domain"

type CourseRepository interface {
	SelectAllCourse() ([]*domain.Course, error)
	SelectCourse(courseID int) (*domain.Course, error)
	CreateCourse(course *domain.Course) (*domain.Course, error)
	DeleteCourse(courseID int) error
	UpdateCourse(course *domain.Course) (*domain.Course, error)
}
