package repository

import (
	"database/sql"
)

type AuthRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (ur *AuthRepositoryImpl) GetPasswordByUserID(userID int) (string, error) {

	var password string
	err := ur.db.QueryRow(
		"SELECT password FROM `users` WHERE `id`=?", userID,
	).Scan(&password)

	if err != nil {
		return "", err
	}
	return password, nil

}

func (ur *AuthRepositoryImpl) GetUserIDByStudentID(studentID string) (int, error) {
	var userID int
	err := ur.db.QueryRow(
		"SELECT id FROM `users` WHERE `student_id`=?", studentID,
	).Scan(&userID)

	if err != nil {
		return -1, err
	}
	return userID, nil
}
