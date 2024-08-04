package repository

import (
	"backend/domain"
)

// インターフェースを実装
type UserRepositoryInterface interface {
	Save(user *domain.User) error
}
