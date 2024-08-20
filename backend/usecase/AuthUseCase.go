package usecase

type AuthUseCase interface {
	Login(userID int, password string) (string, error)
	// Logout(token string) error
}
