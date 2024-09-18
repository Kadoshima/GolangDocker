package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/infrastructure/auth"
	"time"
)

type ForumUseCaseImpl struct {
	ForumRepository repository.ForumRepository
	JWTService      *auth.JWTService
}

func NewForumUseCase(forumRepository repository.ForumRepository, jwtService *auth.JWTService) *ForumUseCaseImpl {
	return &ForumUseCaseImpl{
		ForumRepository: forumRepository,
		JWTService:      jwtService,
	}
}

func (fu *ForumUseCaseImpl) CreateForum(title, description string, createdBy, status, visibility int, category string) (*domain.Forums, error) {
	forum := &domain.Forums{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
		Status:      status,
		Visibility:  visibility,
		Category:    category,
		NumPosts:    0,
		Moderators:  []int{createdBy},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdForum, err := fu.ForumRepository.CreateForum(forum)
	if err != nil {
		return nil, err
	}

	return createdForum, nil
}

func (fu *ForumUseCaseImpl) JoinForum(forumID int) (err error) {
	return nil
}
