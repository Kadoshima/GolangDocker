package repository

import (
	"backend/domain"
	"database/sql"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) Save(user *domain.User) error {
	_, err := ur.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Nickname, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
