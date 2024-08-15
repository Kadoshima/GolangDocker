package usecase

import (
	"backend/domain"
)

// インターフェースを実装
type UserUseCase interface {
	CreateUser(user *domain.User) error
	// UserInfoGet(userId int) error
	// UserInfoUpdate(user *domain.User) error
}
