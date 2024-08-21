package repository

import (
	"backend/domain"
	"database/sql"
	"fmt"
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

func (ur *UserRepositoryImpl) Select(userID int) (*domain.User, error) {

	// domain.Userへのポインタ
	user := &domain.User{}

	err := ur.db.QueryRow(
		"SELECT id, student_id, nickname, email, department_id, course_id FROM `users` WHERE `id`=?", userID,
	).Scan(&user.ID, &user.StudentID, &user.Nickname, &user.Email, &user.DepartmentID, &user.CourseID)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepositoryImpl) Update(user *domain.User, sql string, sqlArgument []interface{}) error {
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `id` = ?", sql)
	sqlArgument = append(sqlArgument, user.ID)
	_, err := ur.db.Exec(query, sqlArgument...)

	if err != nil {
		return err
	}
	return nil
}
