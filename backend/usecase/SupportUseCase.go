package usecase

import "backend/domain"

type SupportUseCase interface {
	NewSupportRequest(supportRequest *domain.SupportRequest) (*domain.SupportRequest, error)
	SupportIsComplete(supportID int) (*domain.SupportRequest, error)
	CloseSupportRequest(supportID int) (*domain.SupportRequest, error)
	GetSupportRequest(supportID int) (*domain.SupportRequest, error)
	GetDepartmentSupportRequests(userID int) ([]*domain.SupportRequest, error)
}
