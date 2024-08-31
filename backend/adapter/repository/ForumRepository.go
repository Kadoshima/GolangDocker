package repository

import "backend/domain"

type ForumRepository interface {
	Create(forums *domain.Forums) error
}
