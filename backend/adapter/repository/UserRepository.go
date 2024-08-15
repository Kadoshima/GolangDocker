package repository

import (
	"backend/domain"
)

// インターフェースを実装
type UserRepository interface {
	Save(user *domain.User) error
	Select(userID int) (string, error)
}
