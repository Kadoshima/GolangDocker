package repository

import (
	"backend/domain"
	"database/sql"
)

type ForumRepositoryImpl struct {
	db *sql.DB
}

func NewForumRepository(db *sql.DB) *ForumRepositoryImpl {
	return &ForumRepositoryImpl{db: db}
}

func (f *ForumRepositoryImpl) Create(forums *domain.Forums) error {
	return nil
}
