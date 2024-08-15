package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
)

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{UserRepository: userRepository}
}

func (uu *UserUseCaseImpl) CreateUser(user *domain.User) error {

	// リポジトリを使ってユーザーを保存
	if err := uu.UserRepository.Save(user); err != nil {
		return err
	}

	return nil
}