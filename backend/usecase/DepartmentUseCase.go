package usecase

import "backend/domain"

type DepartmentUseCase interface {
	GetAllDepartments() ([]domain.Department, error)
	GetDepartmentInfo(departmentID int) (*domain.Department, error)
	StoreDepartmentInfo(department *domain.Department) error
}
