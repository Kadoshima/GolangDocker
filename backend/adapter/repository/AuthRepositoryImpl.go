package repository

import (
	"database/sql"
)

type AuthRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (ur *AuthRepositoryImpl) Select(userID int) error {

	return nil
}
