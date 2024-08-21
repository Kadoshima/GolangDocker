package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"fmt"
	"reflect"
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

func (uu *UserUseCaseImpl) UserInfoGet(userID int) (*domain.User, error) {

	user, err := uu.UserRepository.Select(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUseCaseImpl) UserInfoUpdate(user *domain.User) error {

	val := reflect.ValueOf(user).Elem()
	typ := val.Type()
	var sql string
	var sqlArgument []interface{}

	// 各フィールドをループして、初期値で無ければ更新
	// Field(0)がIDなので1からfor文
	for i := 1; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		// sql,sql引数更新
		if !field.IsZero() {
			sqlArgument = append(sqlArgument, field.Interface())
			if sql == "" {
				sql = fmt.Sprintf("`%s` = ?", domain.DatabaseFields[fieldName])
			} else {
				sql = fmt.Sprintf("%s, %s = ?", sql, fieldName)
			}
		}
	}
	// sqlが空の場合、エラーを返す(更新がない時)
	if sql == "" {
		return fmt.Errorf("no fields to update")
	}

	// 更新処理をリポジトリに委譲
	if err := uu.UserRepository.Update(user, sql, sqlArgument); err != nil {
		return err
	}

	return nil
}
