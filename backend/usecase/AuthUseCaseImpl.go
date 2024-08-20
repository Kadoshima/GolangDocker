package usecase

import "backend/adapter/repository"

type AuthUseCaseImpl struct {
	UserRepository repository.UserRepository
	AuthRepository repository.AuthRepository
}

func NewAuthUseCase(userRepository repository.UserRepository, authRepository repository.AuthRepository) *AuthUseCaseImpl {
	return &AuthUseCaseImpl{UserRepository: userRepository, AuthRepository: authRepository}
}

func (au *AuthUseCaseImpl) Login(userID int, password string) (string, error) {
	println("login")
	return "ok", nil
}
