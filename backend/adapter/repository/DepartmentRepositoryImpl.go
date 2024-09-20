package repository

import (
	"backend/domain"
	"database/sql"
	"time"
)

type DepartmentRepositoryImpl struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepositoryImpl {
	return &DepartmentRepositoryImpl{db: db}
}

func (dr *DepartmentRepositoryImpl) SelectAllDepartments() ([]domain.Department, error) {
	rows, err := dr.db.Query("SELECT id,name FROM departments")
	if err != nil {
		return nil, err
	}

	var departments []domain.Department
	for rows.Next() {
		var department domain.Department
		if err := rows.Scan(&department.ID, &department.Name); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	return departments, nil
}

func (dr *DepartmentRepositoryImpl) SelectDepartment(departmentID int) (*domain.Department, error) {
	department := &domain.Department{}
	err := dr.db.QueryRow(
		"SELECT id, name FROM departments WHERE id = ?",
		departmentID,
	).Scan(&department.ID, &department.Name)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (dr *DepartmentRepositoryImpl) CreateDepartment(department *domain.Department) (*domain.Department, error) {
	now := time.Now()
	department.CreatedAt = now
	department.UpdatedAt = now

	result, err := dr.db.Exec(
		"INSERT INTO departments (name, created_at, updated_at) VALUES (?, ?, ?)",
		department.Name, department.CreatedAt, department.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	departmentID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	department.ID = int(departmentID)
	return department, nil
}

func (dr *DepartmentRepositoryImpl) UpdateDepartment(department *domain.Department) (*domain.Department, error) {
	department.UpdatedAt = time.Now()
	_, err := dr.db.Exec(
		"UPDATE departments SET name = ?, updated_at = ? WHERE id = ?",
		department.Name, department.UpdatedAt, department.ID,
	)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (dr *DepartmentRepositoryImpl) DeleteDepartment(departmentID int) error {
	_, err := dr.db.Exec(
		"DELETE FROM departments WHERE id = ?",
		departmentID,
	)
	return err
}
