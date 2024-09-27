package usecase

import "backend/domain"

type PostUseCase interface {
	NewPost(forumID, userID int, content string, parentID int) (*domain.Post, error)
	GetPosts(forumID int) ([]domain.Post, error)
}
