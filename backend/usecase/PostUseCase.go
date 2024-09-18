package usecase

import "backend/domain"

type PostUseCase interface {
	NewPost(forumID, userID int, content, tags string, parentID *int) (*domain.Post, error)
}
