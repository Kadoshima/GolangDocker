package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
)

type ForumUseCaseImpl struct {
	ForumRepository repository.ForumRepository
}

func NewForumUseCase(ForumRepository repository.ForumRepository) *ForumUseCaseImpl {
	return &ForumUseCaseImpl{
		ForumRepository: ForumRepository,
	}
}

func (f *ForumUseCaseImpl) NowForum(title string, description string, user *domain.User, visibility int, category string) *domain.Forums {

	// UseCaseの中ではドメイン層のロジックを利用しながら、データの永続化やリトリーバルを含む一連のプロセスを調整します。

	return domain.NewForum(title, description, user, visibility, category)
}
