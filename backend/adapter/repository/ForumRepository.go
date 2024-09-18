package repository

import "backend/domain"

type ForumRepository interface {
	SelectAllForums() ([]domain.Forums, error)
	SelectForum(ForumID int) (*domain.Forums, error)
	CreateForum(forum *domain.Forums) (*domain.Forums, error)
	UpdateForum(forum *domain.Forums) (*domain.Forums, error)
	DeleteForum(forum *domain.Forums) error
}
