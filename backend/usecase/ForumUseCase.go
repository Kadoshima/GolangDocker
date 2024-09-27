package usecase

import "backend/domain"

type ForumUseCase interface {
	CreateForum(title, description string, createdBy, status, visibility int, category string, attachments string) (*domain.Forums, error)
	GetForum() ([]*domain.Forums, error)
	JoinForum(int) error
}
