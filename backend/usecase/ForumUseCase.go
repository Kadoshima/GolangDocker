package usecase

import "backend/domain"

type ForumUseCase interface {
	CreateForum(title, description string, createdBy, status, visibility int, category string) (*domain.Forums, error)
	JoinForum(int) error
}
