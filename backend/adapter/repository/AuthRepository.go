package repository

type AuthRepository interface {
	Select(userID int) error
}
