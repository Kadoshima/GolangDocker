package repository

import "backend/domain"

type DepartmentRepository interface {
	SelectAllDepartments() ([]*domain.Department, error)
	SelectDepartment(departmentId int) (*domain.Department, error)
	CreateDepartment(department *domain.Department) (*domain.Department, error)
	UpdateDepartment(department *domain.Department) (*domain.Department, error)
	DeleteDepartment(departmentId int) error
}
