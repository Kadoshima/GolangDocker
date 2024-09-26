package repository

type AuthRepository interface {
	GetPasswordByUserID(userID int) (string, error)
	GetUserIDByStudentID(studentID string) (int, error)
}
