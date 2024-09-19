package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/infrastructure/auth"
	"time"
)

type PostUseCaseImpl struct {
	postRepository repository.PostRepository
	JWTService     *auth.JWTService
}

func NewPostUseCase(postRepository repository.PostRepository, JWTService *auth.JWTService) *PostUseCaseImpl {
	return &PostUseCaseImpl{postRepository, JWTService}
}

func (pu *PostUseCaseImpl) NewPost(forumID, userID int, content, tags string, parentID int) (*domain.Post, error) {
	post := &domain.Post{
		ForumId:   forumID,
		UserId:    userID,
		Content:   content,
		Tags:      tags,
		Status:    1,
		ParentId:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdPost, err := pu.postRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}
