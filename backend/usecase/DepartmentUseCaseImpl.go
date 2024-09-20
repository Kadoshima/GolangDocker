package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/infrastructure/auth"
)

type DepartmentUseCaseImpl struct {
	DepartmentRepository repository.DepartmentRepository
	JWTService           *auth.JWTService
}

func NewDepartmentUseCase(repository repository.DepartmentRepository, JWTService *auth.JWTService) *DepartmentUseCaseImpl {
	DepartmentUseCaseImpl := &DepartmentUseCaseImpl{
		DepartmentRepository: repository,
		JWTService:           JWTService,
	}
	return DepartmentUseCaseImpl
}

func (di DepartmentUseCaseImpl) GetAllDepartments() ([]domain.Department, error) {
	// 全てのdepartment情報を返す
	departments, err := di.DepartmentRepository.SelectAllDepartments()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (di DepartmentUseCaseImpl) GetDepartmentInfo(departmentID int) (*domain.Department, error) {
	// 特定のdepartment情報を返す
	department, err := di.DepartmentRepository.SelectDepartment(departmentID)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (di DepartmentUseCaseImpl) StoreDepartmentInfo(department *domain.Department) error {
	// Departmentの追加は一旦いらない
	return nil
}
