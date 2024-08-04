package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
)

type UserUsecaseImpl struct {
	UserRepositoryInterface *repository.UserRepositoryInterface
}

func (uu *UserUsecaseImpl) CreateUser(user *domain.User) error {

	// リポジトリを使ってユーザーを保存
	if err := uu.UserRepositoryInterface.Save(user); err != nil {
		return err
	}

	return nil
}
