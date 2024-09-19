package usecase

import "backend/domain"

type CourseUseCase interface {
	GetCourseInfo(courseID int) (*domain.Course, error)
	StoreCourseInfo(Course *domain.Course) error
	GetAllCourseInfo() ([]*domain.Course, error)
}
