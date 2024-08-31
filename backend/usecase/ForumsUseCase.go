package usecase

import "backend/domain"

type ForumsUseCase interface {
	NowForum(title string, description string, user *domain.User, visibility int, category string) *domain.Forums
}
