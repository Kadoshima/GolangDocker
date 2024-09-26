package usecase

type AuthUseCase interface {
	Login(userID int, password string) (string, error)
	GetUserIDByStudentID(studentID string) (int, error)
}
