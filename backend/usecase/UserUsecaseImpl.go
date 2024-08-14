package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

func (uu *UserUsecaseImpl) CreateUser(user *domain.User) error {

	// リポジトリを使ってユーザーを保存
	if err := uu.UserRepository.Save(user); err != nil {
		return err
	}

	return nil
}
