package repository

import "backend/domain"

type PostRepository interface {
	SelectPost(forumID int) ([]domain.Post, error)
	CreatePost(post *domain.Post) (*domain.Post, error)
	UpdatePost(post *domain.Post) (*domain.Post, error)
	DeletePost(postID int) error
}
