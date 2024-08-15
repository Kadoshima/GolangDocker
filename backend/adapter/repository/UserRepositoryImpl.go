package repository

import (
	"backend/domain"
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) Save(user *domain.User) error {
	_, err := ur.db.Exec(
		"INSERT INTO `users`(`student_id`, `nickname`, `email`, `password`, `department_id`, `course_id`)"+
			" VALUES (?, ?, ?, ?, ?, ?)",
		user.StudentID, user.Nickname, user.Email, user.Password, user.DepartmentID, user.CourseID,
	)
	if err != nil {
		return err
	}
	return nil
}
