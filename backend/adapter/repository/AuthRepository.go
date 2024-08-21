package repository

type AuthRepository interface {
	GetPasswordByUserID(userID int) (string, error)
}
