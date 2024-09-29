package repository

import (
	"backend/domain"
)

// インターフェースを実装
type UserRepository interface {
	Create(user *domain.User) error
	Select(userID int) (*domain.User, error)
	Update(user *domain.User, sql string, sqlArgument []interface{}) error
	SelectByUserID(userID int) (*domain.User, error)
}
