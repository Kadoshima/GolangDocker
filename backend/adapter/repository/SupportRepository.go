package repository

import (
	"backend/domain"
)

type SupportRepository interface {
	CreateSupport(support domain.SupportRequest) (domain.SupportRequest, error)
	SelectSupport() ([]domain.SupportRequest, error)
	UpdateSupport(support domain.SupportRequest) (domain.SupportRequest, error)
	DeleteSupport(support domain.SupportRequest) error
	CompleteSupport(support domain.SupportRequest) (domain.SupportRequest, error)
	SelectSupportByID(supportID int) (*domain.SupportRequest, error)
	CloseSupport(support domain.SupportRequest) (domain.SupportRequest, error)
	SelectSupportByForumIDAndStatus(forumID int, status domain.SupportRequestStatus) (*domain.SupportRequest, error)
}
